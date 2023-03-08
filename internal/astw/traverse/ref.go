package traverse

import (
	"golang.org/x/exp/slices"
)

type Ref interface {
	Set(value any) bool
}

type SliceItemInsertBeforeRef[T any] struct {
	sliceRef          SliceRef[T]
	insertBeforeIndex int
}

func (ref *SliceItemInsertBeforeRef[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		// fmt.Println("SliceItemInsertBeforeRef>", *ref.sliceRef.Get(), value)
		newSlice := slices.Insert(*ref.sliceRef.Get(), ref.insertBeforeIndex, value)
		ok := ref.sliceRef.Set(&newSlice)
		// fmt.Println("SliceItemInsertBeforeRef>>", *ref.sliceRef.Get())
		return ok
	}
	return false
}

func NewSliceItemInsertBeforeRef[T any](sliceRef SliceRef[T], insertBeforeIndex int) *SliceItemInsertBeforeRef[T] {
	return &SliceItemInsertBeforeRef[T]{
		sliceRef:          sliceRef,
		insertBeforeIndex: insertBeforeIndex,
	}
}

type SliceEndingRef[T any] struct {
	sliceRef SliceRef[T]
}

func (ref *SliceEndingRef[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		// fmt.Println("SliceEndingRef>", ref.sliceRef, value)
		newSlice := append(*ref.sliceRef.Get(), value)
		ok := ref.sliceRef.Set(&newSlice)
		// fmt.Println("SliceEndingRef>>", ref.sliceRef)
		return ok
	}
	return false
}

func NewSliceEndingRef[T any](sliceRef SliceRef[T]) *SliceEndingRef[T] {
	return &SliceEndingRef[T]{
		sliceRef: sliceRef,
	}
}

type SliceItemRef[T any] struct {
	sliceAddr *[]T
	index     int
}

func (ref *SliceItemRef[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		// fmt.Println("SliceItemRef>", *ref.sliceAddr, value)
		(*ref.sliceAddr)[ref.index] = value
		// fmt.Println("SliceItemRef>", *ref.sliceAddr)
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

type SliceRef[T any] struct {
	addr *[]T
}

func (ref *SliceRef[T]) Set(valuePtr any) bool {
	if valuePtr, ok := valuePtr.(*[]T); ok {
		// fmt.Printf("SliceRef+ %p, %p\n", *ref.addr, ref.addr)
		// fmt.Println("SliceRef>", *ref.addr, ref.addr, reflect.TypeOf(*ref.addr), reflect.ValueOf(*ref.addr), reflect.ValueOf(ref.addr), valuePtr)
		*ref.addr = *valuePtr
		// fmt.Printf("SliceRef++ %p, %p\n", *ref.addr, ref.addr)
		// fmt.Println("SliceRef>>", *ref.addr, ref.addr, reflect.TypeOf(*ref.addr), reflect.ValueOf(*ref.addr), reflect.ValueOf(ref.addr))
		return true
	}
	return false
}

func (ref *SliceRef[T]) Get() *[]T {
	return ref.addr
}

func NewSliceRef[T any](addr *[]T) *SliceRef[T] {
	return &SliceRef[T]{
		addr: addr,
	}
}

type DirectRef[T any] struct {
	addr *T
}

func (ref *DirectRef[T]) Set(valuePtr any) bool {
	if value, ok := valuePtr.(T); ok {
		// fmt.Println("DirectRef>", *ref.addr, value)
		*ref.addr = value
		// fmt.Println("DirectRef>>", *ref.addr)
		return true
	}
	return false
}

func NewDirectRef[T any](addr *T) *DirectRef[T] {
	return &DirectRef[T]{
		addr: addr,
	}
}
