package codefragment

import (
	"fmt"
	"testing"
)

func compareByteArrays(left, right []byte) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func Test_ProduceRandomFragment(t *testing.T) {
	var prevFragment = []byte{}

	for i := 0; i < 100; i++ {
		var fragment = ProduceRandomFragment()
		fmt.Println(string(fragment))

		if compareByteArrays(fragment, prevFragment) {
			t.Error("Produced same code fragment twice.")
		} else {
			prevFragment = fragment
		}
	}

}
