package node_constructor_v2

import (
	"fmt"
	"go/ast"
)

var (
	prefixVariableName  = "tdeVar"
	prefixFunctionName  = "tdeFunc"
	prefixTypeName      = "tdeType"
	prefixInterfaceName = "tdeInterface"
	prefixStruct        = "tdeStruct"
	counter             = -1
)

func NewVariable() ast.Ident {
	counter++
	return *ast.NewIdent(fmt.Sprintf("%s%d", prefixVariableName, counter))
}

func NewFunction() ast.Ident {
	counter++
	return *ast.NewIdent(fmt.Sprintf("%s%d", prefixFunctionName, counter))
}

func NewTypeName() ast.Ident {
	counter++
	return *ast.NewIdent(fmt.Sprintf("%s%d", prefixTypeName, counter))
}

func NewInterfaceName() ast.Ident {
	counter++
	return *ast.NewIdent(fmt.Sprintf("%s%d", prefixInterfaceName, counter))
}

func NewStructName() ast.Ident {
	counter++
	return *ast.NewIdent(fmt.Sprintf("%s%d", prefixStruct, counter))
}
