package wordreverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)




func Test_WordReverse(t *testing.T) {
	testParameters := map[string]string{
		"Hello world":         "dlrow olleH",
		"The quick brown fox": "xof nworb kciuq ehT",
	}
	for input, want := range testParameters {
		output := WordReverse(input)
		assert.Equal(t, want, output)
	}
}
