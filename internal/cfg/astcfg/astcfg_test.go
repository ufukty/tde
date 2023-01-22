package astcfg

import (
	"tde/internal/ast_wrapper"
	"testing"

	"github.com/pkg/errors"
)

const TEST_FILE = `package main
	
func Addition(a, b int) int {
	return a + b
}
`

func Test_NoSyntaxError(t *testing.T) {
	_, astNode, _ := ast_wrapper.ParseString(TEST_FILE)
	funcDecl, _ := ast_wrapper.FindFuncDecl(astNode, "Addition")
	MakeAddition(funcDecl)
}

func Test_Develop(t *testing.T) {
	// TODO: Give valid go function body
	// TODO: Compare output with input; Expected: less than 10% difference; more than 0% difference; still valid syntax

	_, astFile, err := ast_wrapper.ParseString(TEST_FILE)
	if err != nil {
		t.Error(errors.Wrapf(err, "failed ParseString"))
	}
	funcDecl, err := ast_wrapper.FindFuncDecl(astFile, "Addition")
	if err != nil {
		t.Error(errors.Wrapf(err, "failed FindFuncDecl"))
	}
	funcDeclModified := funcDecl

	if funcDecl == funcDeclModified {
		t.Error("failed on comparison. no change found.")
	}

	if _, err := ast_wrapper.String(funcDeclModified); err != nil {
		t.Error(errors.Wrap(err, "Has no valid syntax anymore."))
	}
}
