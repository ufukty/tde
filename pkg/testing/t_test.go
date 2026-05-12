package testing

import (
	"testing"
)

func exampleTarget(in string) string {
	return ""
}

func exampleUnitTest(t *T) {
	cases := map[string]string{
		"Fusce a orci leo":             "oel icro a ecsuF",
		"Nulla a mollis est":           "tse sillom a alluN",
		"Etiam a semper nisl":          "lsin repmes a maitE",
		"Cras eget nulla diam":         "maid allun tege sarC",
		"Etiam in diam ligula":         "alugil maid ni maitE",
		"Aliquam erat volutpat":        "taptulov tare mauqilA",
		"Integer vel tinsidunt erat":   "tare tnudicnit lev regetnI",
		"Phasellus ut commodo neque":   "euqen odommoc tu sullesahP",
		"Curabitur id elementum augue": "eugua mutnemele di rutibaruC",
	}

	for input, want := range cases {
		got := exampleTarget(input)
		if !t.Assert(got, want) {
			t.Fatalf("WordReverse(%q) = %q (want: %q)", input, got, want)
		}
	}
}

// enough if it doesn't panic
func Test_T(t *testing.T) {
	tt := NewT("00000000-0000-0000-0000-000000000000")
	exampleUnitTest(tt)
	if len(tt.AssertionResults) != 9 {
		t.Fatalf("assert: expected %d, got %d", 9, len(tt.AssertionResults))
	}
}
