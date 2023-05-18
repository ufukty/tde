package utilities

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_GetPrimeFactors(t *testing.T) {
	t.Run("explicit test", func(t *testing.T) {
		type TestCase struct {
			input int
			want  []int
		}
		testCases := []TestCase{
			{2, []int{2}},
			{3, []int{3}},
			{4, []int{2, 2}},
			{5, []int{5}},
			{6, []int{2, 3}},
			{7, []int{7}},
			{8, []int{2, 2, 2}},
			{9, []int{3, 3}},
			{10, []int{2, 5}},
			{11, []int{11}},
			{12, []int{2, 2, 3}},
			{13, []int{13}},
			{14, []int{2, 7}},
			{15, []int{3, 5}},
			{16, []int{2, 2, 2, 2}},
			{17, []int{17}},
			{18, []int{2, 3, 3}},
			{19, []int{19}},
			{20, []int{2, 2, 5}},
			{21, []int{3, 7}},
			{22, []int{2, 11}},
			{23, []int{23}},
			{24, []int{2, 2, 2, 3}},
			{25, []int{5, 5}},
			{26, []int{2, 13}},
			{27, []int{3, 3, 3}},
			{28, []int{2, 2, 7}},
			{29, []int{29}},
			{30, []int{2, 3, 5}},
			{31, []int{31}},
			{32, []int{2, 2, 2, 2, 2}},
			{33, []int{3, 11}},
			{34, []int{2, 17}},
			{35, []int{5, 7}},
			{36, []int{2, 2, 3, 3}},
			{37, []int{37}},
			{38, []int{2, 19}},
			{39, []int{3, 13}},
			{40, []int{2, 2, 2, 5}},
			{41, []int{41}},
			{42, []int{2, 3, 7}},
			{43, []int{43}},
			{44, []int{2, 2, 11}},
			{45, []int{3, 3, 5}},
			{46, []int{2, 23}},
			{47, []int{47}},
			{48, []int{2, 2, 2, 2, 3}},
			{49, []int{7, 7}},
		}
		for _, testCase := range testCases {
			if got := GetPrimeFactors(testCase.input); slices.Compare(testCase.want, got) != 0 {
				t.Error("validation. input:", testCase.input, "want:", testCase.want, "got:", got)
			}

		}
	})

	t.Run("fast test", func(t *testing.T) {
		for i := 50; i < 100000; i++ {
			factors := GetPrimeFactors(i)
			multiplication := 1
			for _, factor := range factors {
				multiplication *= factor
			}
			if multiplication != i {
				t.Error("validation. i=", i)
			}
		}
	})

}
