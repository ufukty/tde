package strw

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_StringFold(t *testing.T) {
	testCases := []struct {
		str    string
		length int
		want   []string
	}{
		{"0", 2, []string{"0"}},
		{"0123", 1, []string{"0", "1", "2", "3"}},
		{"0123", 2, []string{"01", "23"}},
		{"0123456", 2, []string{"01", "23", "45", "6"}},
		{"0123456", 4, []string{"0123", "456"}},
		{"0123456", 10, []string{"0123456"}},
	}

	for _, testCase := range testCases {
		if got := Fold(testCase.str, testCase.length); slices.Compare(testCase.want, got) != 0 {
			t.Error("validation", testCase.want, got)
		}
	}
}
