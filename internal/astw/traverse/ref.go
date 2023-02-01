package traverse

type Ref interface {
	Set(value any) bool
}

type SliceRef[T any] struct {
	addr *[]T
}

func (ref *SliceRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		*(ref.addr) = append(*(ref.addr), value)
		return true
	}
	return false
}

func NewSliceRef[T any](slicePtr *[]T) *SliceRef[T] {
	return &SliceRef[T]{
		addr: slicePtr,
	}
}

type StructRef[T any] struct {
	addr *T
}

func (ref *StructRef[T]) Set(value any) bool {
	if value, ok := value.(T); ok {
		*ref.addr = value
		return true
	}
	return false
}

func NewStructRef[T any](structOrFieldPtr *T) *StructRef[T] {
	return &StructRef[T]{
		addr: structOrFieldPtr,
	}
}
