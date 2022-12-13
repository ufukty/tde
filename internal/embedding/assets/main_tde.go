//go:build tde
// +build tde

package main

import (
	wordreverse "tde/examples/word-reverse/word_reverse"
	"tde/pkg/tde"
)

func main() {
	e := tde.NewE(candidates)
	wordreverse.TDE_WordReverse(e)
	e.Export()
}
