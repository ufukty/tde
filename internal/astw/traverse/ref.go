package traverse

import (
	"golang.org/x/exp/slices"
)

type Ref interface {
	Set(value any) bool
}

type SliceItemInsertBeforeRef[T any] struct {
	sliceAddr         *[]T
	insertBeforeIndex int
}

func (ref *SliceItemInsertBeforeRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		*ref.sliceAddr = slices.Insert((*ref.sliceAddr), ref.insertBeforeIndex, value)
		return true
	}
	return false
}

func NewSliceItemInsertBeforeRef[T any](slicePtr *[]T, insertBeforeIndex int) *SliceItemInsertBeforeRef[T] {
	return &SliceItemInsertBeforeRef[T]{
		sliceAddr:         slicePtr,
		insertBeforeIndex: insertBeforeIndex,
	}
}

type SliceEndingRef[T any] struct {
	sliceAddr *[]T
}

func (ref *SliceEndingRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		*ref.sliceAddr = append(*ref.sliceAddr, value)
		return true
	}
	return false
}

func NewSliceEndingRef[T any](slicePtr *[]T) *SliceEndingRef[T] {
	return &SliceEndingRef[T]{
		sliceAddr: slicePtr,
	}
}

type SliceItemRef[T any] struct {
	sliceAddr *[]T
	index     int
}

func (ref *SliceItemRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		(*ref.sliceAddr)[ref.index] = value
		return true
	}
	return false
}

func NewSliceItemRef[T any](slicePtr *[]T, index int) *SliceItemRef[T] {
	return &SliceItemRef[T]{
		sliceAddr: slicePtr,
		index:     index,
	}
}

type DirectRef[T any] struct {
	addr *T
}

func (ref *DirectRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		*ref.addr = value
		return true
	}
	return false
}

func NewDirectRef[T any](addr *T) *DirectRef[T] {
	return &DirectRef[T]{
		addr: addr,
	}
}
