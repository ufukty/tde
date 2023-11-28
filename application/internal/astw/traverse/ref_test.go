package traverse

import (
	"fmt"
	"testing"
)

func Test_DirectRef(t *testing.T) {
	s := struct {
		StructField string
	}{}
	ref := newFieldRef(&s.StructField)
	valueToSet := "Hello World"

	if !ref.Set(valueToSet) {
		t.Error("Call returned false")
	}
	if s.StructField != valueToSet {
		t.Error("Validation")
	}
}

func Test_SliceItemInsertBeforeRef(t *testing.T) {
	slice := []int{10, 11, 13, 14, 15}
	sliceRef := *newSliceRef(&slice)
	sliceItemRef := newSliceItemBefore(sliceRef, 2)
	if !sliceItemRef.Set(12) {
		t.Error("On process")
	}
	// fmt.Println(sliceItemRef.Get())
	fmt.Println(sliceRef.Get())
	if slice[2] != 12 {
		t.Error("Validation")
	}
}

func Test_SliceRef(t *testing.T) {
	slice := []int{10, 11, 12, 13, 14}
	sliceRef := newSliceEndingRef(*newSliceRef(&slice))
	sliceRef.Set(15)
	if slice[5] != 15 {
		t.Error("Validation")
	}
}

func Test_SliceItemRef(t *testing.T) {
	slice := []int{10, 11, 22, 13, 14}
	sliceRef := newSliceItemRef(&slice, 2)
	sliceRef.Set(12)
	if slice[2] != 12 {
		t.Error("Validation")
	}
}
