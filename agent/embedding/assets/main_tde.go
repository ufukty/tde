//go:build tde
// +build tde

package main

import (
	wordreverse "GoGP/examples/word-reverse/word_reverse"
	"GoGP/testing/evolution"
)

func main() {
	e := evolution.NewE(candidates)
	wordreverse.TDE_WordReverse(e)
	e.Export()
}
