package traverse

import "slices"

type ref interface {
	Set(value any) bool
	Get() any
}

// TODO: Get
type sliceItemBefore[T any] struct {
	slice             slice[T]
	insertBeforeIndex int
}

func (ref *sliceItemBefore[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		newSlice := slices.Insert(*ref.slice.Get(), ref.insertBeforeIndex, value)
		ok := ref.slice.Set(&newSlice)
		return ok
	}
	return false
}

func newSliceItemBefore[T any](sr slice[T], insertBeforeIndex int) *sliceItemBefore[T] {
	return &sliceItemBefore[T]{
		slice:             sr,
		insertBeforeIndex: insertBeforeIndex,
	}
}

type sliceEnding[T any] struct {
	slice slice[T]
}

func (ref *sliceEnding[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		newSlice := append(*ref.slice.Get(), value)
		ok := ref.slice.Set(&newSlice)
		return ok
	}
	return false
}

func newSliceEndingRef[T any](sr slice[T]) *sliceEnding[T] {
	return &sliceEnding[T]{
		slice: sr,
	}
}

type sliceitem[T any] struct {
	sliceAddr *[]T
	index     int
}

func (ref *sliceitem[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		(*ref.sliceAddr)[ref.index] = value
		return true
	}
	return false
}

func (ref *sliceitem[T]) Get() any {
	return (*ref.sliceAddr)[ref.index]
}

func newSliceItemRef[T any](slicePtr *[]T, index int) *sliceitem[T] {
	return &sliceitem[T]{
		sliceAddr: slicePtr,
		index:     index,
	}
}

type slice[T any] struct {
	addr *[]T
}

func (ref *slice[T]) Set(valuePtr any) bool {
	if valuePtr, ok := valuePtr.(*[]T); ok {
		*ref.addr = *valuePtr
		return true
	}
	return false
}

func (ref *slice[T]) Get() *[]T {
	return ref.addr
}

func newSliceRef[T any](addr *[]T) *slice[T] {
	return &slice[T]{
		addr: addr,
	}
}

type field[T any] struct {
	addr *T
}

func (ref *field[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		*ref.addr = value
		return true
	}
	return false
}

func (ref *field[T]) Get() any {
	return *ref.addr
}

func newFieldRef[T any](addr *T) *field[T] {
	return &field[T]{
		addr: addr,
	}
}
