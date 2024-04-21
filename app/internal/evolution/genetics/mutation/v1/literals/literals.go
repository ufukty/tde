package literals

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"tde/internal/evolution/genetics/mutation/v1/models"
	"tde/internal/utilities/pick"
	"tde/internal/utilities/randoms"
)

func listApplicableNodes(n ast.Node) (applicableNodes []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch n := n.(type) {
		case *ast.BasicLit:
			switch n.Kind {
			case
				token.STRING,
				token.CHAR,
				token.INT,
				token.FLOAT:
			}
			applicableNodes = append(applicableNodes, n)
		case *ast.Ident:
			if n.Name == "true" || n.Name == "false" {
				applicableNodes = append(applicableNodes, n)
			}
		}
		return true
	})
	return
}

var allowedCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:;\n ")

func mutateLiteralString(s string) string {
	r := []rune(s)
	rnd := randoms.UniformIntN(len(r))
	r[rnd], _ = pick.Pick(allowedCharacters)
	return string(r)
}

func mutateLiteralChar() string {
	p, _ := pick.Pick(allowedCharacters)
	return string(p)
}

func mutateLiteralInteger(str string) string {
	integer, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return str
	}
	if pick.Coin() {
		integer++
	} else {
		integer--
	}
	return fmt.Sprintf("%d", integer)
}

func mutateLiteralFloat(str string) string {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return str
	}
	if pick.Coin() {
		float *= 1.1
	} else {
		float *= 0.9
	}
	return fmt.Sprintf("%f", float)
}

func mutateIdentBoolean(str string) string {
	if str == "true" {
		return "false"
	} else if str == "false" {
		return "true"
	} else {
		return str
	}
}

func literalValueAlter(choosenNode ast.Node) {
	switch choosenNode := choosenNode.(type) {
	case *ast.BasicLit:
		switch choosenNode.Kind {
		case token.CHAR:
			choosenNode.Value = mutateLiteralChar()
		case token.STRING:
			choosenNode.Value = mutateLiteralString(choosenNode.Value)
		case token.INT:
			choosenNode.Value = mutateLiteralInteger(choosenNode.Value)
		case token.FLOAT:
			choosenNode.Value = mutateLiteralFloat(choosenNode.Value)
		}
	case *ast.Ident:
		choosenNode.Name = mutateIdentBoolean(choosenNode.Name)
	}
}

func LiteralValueAlter(ctx *models.MutationParameters) error {
	applicableNodes := listApplicableNodes(ctx.FuncDecl.Body)
	choosenNode, err := pick.Pick(applicableNodes)
	if err != nil {
		return fmt.Errorf("picking one ouf of many applicable nodes: %w", err)
	}
	literalValueAlter(choosenNode)
	return nil
}
