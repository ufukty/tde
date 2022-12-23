package character

import (
	"fmt"
	"tde/internal/utilities"
	"testing"
)

func Test_ProduceRandomFragment(t *testing.T) {
	var prevFragment = []byte{}

	for i := 0; i < 100; i++ {
		var fragment = ProduceRandomFragment()
		fmt.Println(i, " ", string(fragment))

		if utilities.CompareSlices(fragment, prevFragment) {
			t.Error("Produced same code fragment twice.")
		} else {
			prevFragment = fragment
		}
	}

}
