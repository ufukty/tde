package traverse

import "golang.org/x/exp/slices"

type Ref interface {
	Set(value any) bool
}

type SliceItemInsertBeforeRef[T any] struct {
	sliceAddr         *[]T
	insertBeforeIndex int
}

func (r *SliceItemInsertBeforeRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		(*r.sliceAddr) = slices.Insert((*r.sliceAddr), r.insertBeforeIndex, value)
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

func (r *SliceEndingRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		*(r.sliceAddr) = append(*(r.sliceAddr), value)
		return true
	}
	return false
}

func NewSliceEndingRef[T any](slicePtr *[]T) *SliceEndingRef[T] {
	return &SliceEndingRef[T]{
		sliceAddr: slicePtr,
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
