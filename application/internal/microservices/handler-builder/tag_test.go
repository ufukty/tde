package handler_builder

import (
	"fmt"
	"testing"

	"golang.org/x/exp/maps"
)

func Test_A(t *testing.T) {
	type TestCase struct {
		Input  string
		Output map[string]string
	}
	var testCases = []TestCase{
		{`Content-Type:"multipart/form-data"`, map[string]string{
			"Content-Type": "multipart/form-data",
		}},
		{`form:"file" header:"filename"`, map[string]string{
			"form":   "file",
			"header": "filename",
		}},
	}

	for _, testCase := range testCases {
		fmt.Println("Test case:", testCase.Input)
		var pairs = ParseTag(testCase.Input)
		for key, value := range pairs {
			fmt.Printf("%s => %s\n", key, value)
		}
		if !maps.Equal(testCase.Output, pairs) {
			t.Errorf("%s", testCase.Input)
		}
		fmt.Println()
	}
}
