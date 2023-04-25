package queue

import (
	"sync"
	"time"
)

type QueueConsumerResponse bool

const (
	DONE             = QueueConsumerResponse(true)
	SEND_AGAIN_LATER = QueueConsumerResponse(false)
)

type QueueConsumer[T any] func(T) QueueConsumerResponse

type Replay[T any] struct {
	asyncQueue    []T
	consumer      QueueConsumer[T]
	sendAgainLock sync.Mutex
}

func NewReplayQueue[T any](consumer QueueConsumer[T]) *Replay[T] {
	q := &Replay[T]{
		consumer: consumer,
	}
	go q.interval()
	return q
}

func (q *Replay[T]) Send(item T) {
	if q.consumer(item) == SEND_AGAIN_LATER {
		q.sendAgainLock.Lock()
		q.asyncQueue = append(q.asyncQueue, item)
		q.sendAgainLock.Unlock()
	}
}

func (q *Replay[T]) sendAgain() {
	if q.sendAgainLock.TryLock() {
		return
	}
	var (
		length  = len(q.asyncQueue)
		passed  = 0
		removed = 0
		item    T
	)
	for passed+removed < length {
		item, q.asyncQueue = q.asyncQueue[passed], q.asyncQueue[passed+1:]
		if q.consumer(item) == DONE {
			removed++
		} else {
			q.asyncQueue = append(q.asyncQueue, item)
			passed++
		}
	}
	q.sendAgainLock.Unlock()
}

func (q *Replay[T]) interval() {
	for range time.Tick(time.Second) {
		q.sendAgain()
	}
}
