//go:build tde
// +build tde

package main

import (
	wordreverse "tde/examples/word-reverse/word_reverse"
	"tde/pkg/evolution"
)

func main() {
	e := evolution.NewE(candidates)
	wordreverse.TDE_WordReverse(e)
	e.Export()
}
