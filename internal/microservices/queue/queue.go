package queue

import (
	"sync"
	"time"
)

type QueueConsumer[T any] func(T)

type Replay[T any] struct {
	asyncQueue    []T
	consumer      QueueConsumer[T]
	sendAgainLock sync.Mutex
	interval      time.Duration
}

func NewReplayQueue[T any](consumer QueueConsumer[T]) *Replay[T] {
	q := &Replay[T]{
		consumer: consumer,
		interval: time.Second,
	}
	go q.tick()
	return q
}

func (q *Replay[T]) Queue(item T) {
	q.sendAgainLock.Lock()
	q.asyncQueue = append(q.asyncQueue, item)
	q.sendAgainLock.Unlock()
}

func (q *Replay[T]) dequeue() {
	if !q.sendAgainLock.TryLock() {
		return
	}
	var replayingQueue []T
	copy(replayingQueue, q.asyncQueue)
	q.asyncQueue = []T{}
	q.sendAgainLock.Unlock()

	for _, item := range replayingQueue {
		q.consumer(item)
	}
}

func (q *Replay[T]) tick() {
	for range time.Tick(q.interval) {
		q.dequeue()
	}
}
