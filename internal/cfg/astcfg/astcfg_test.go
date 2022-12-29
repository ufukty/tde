package astcfg

import (
	"tde/internal/code"
	"tde/pkg/tde"
	"testing"
)

func Test_Develop(t *testing.T) {
	// TODO: Give valid go function body
	// TODO: Compare output with input; Expected: less than 10% difference; more than 0% difference; still valid syntax

	c := code.NewCode()
	c.LoadFromString(`package main
	
	func Foo() {

	}
	`)
	f, _ := c.FindFunction("Foo")
	original := c.PartialString(f)

	cfg := ASTCFG{}
	cfg.Develop(f)
	developed := c.PartialString(f)

	if dist := tde.StringDistance(original, developed); !(0 < dist && dist < 0.1) {
		t.Error("More difference than expected:", dist)
	}
}
