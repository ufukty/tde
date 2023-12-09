package nodes

import (
	"fmt"
	"go/ast"
	"tde/internal/utilities/pick"
)

// only valid values are types such int, float, string, bool
func (c *Creator) IdentType(l int) (*ast.Ident, error) {
	if l == 0 {
		return nil, nil
	}
	p, err := pick.Pick([]string{
		"bool",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"string",
	})
	if err != nil {
		return nil, fmt.Errorf("picking a type ident: %w", err)
	}
	return ast.NewIdent(p), err
}

// FIXME: is there any usecase thar is not achievable with a slice but only with a ...T array
func (c *Creator) ArrayType(l int) (*ast.ArrayType, error) {
	if l == 0 {
		return nil, nil
	}
	Elt, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ArrayType.Elt: %w", err)
	}

	return &ast.ArrayType{
		Len: nil,
		Elt: Elt,
	}, nil
}

func (c *Creator) ChanType(l int) (*ast.ChanType, error) {
	if l == 0 {
		return nil, nil
	}
	Value, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ChanType.Value: %w", err)
	}
	Dir, err := pick.Pick([]ast.ChanDir{ast.SEND, ast.RECV})
	if err != nil {
		return nil, fmt.Errorf("picking a direction for ChangType.Dir: %w", err)
	}

	return &ast.ChanType{
		Dir:   Dir,
		Value: Value,
	}, nil
}

// FIXME: add generics support
func (c *Creator) FuncType(l int) (*ast.FuncType, error) {
	if l == 0 {
		return nil, nil
	}
	// TypeParams, err := c.FieldList(l - 1)
	// if err != nil {
	// 	return nil, fmt.Errorf("generating FuncType.TypeParams: %w", err)
	// }
	Params, err := c.FieldList(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating FuncType.Params: %w", err)
	}
	Results, err := c.FieldList(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating FuncType.Results: %w", err)
	}

	return &ast.FuncType{
		// TypeParams: TypeParams,
		Params:  Params,
		Results: Results,
	}, nil
}

func (c *Creator) InterfaceType(l int) (*ast.InterfaceType, error) {
	if l == 0 {
		return nil, nil
	}
	Methods, err := c.FieldList(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating InterfaceType.Methods: %w", err)
	}

	return &ast.InterfaceType{
		Incomplete: false,
		Methods:    Methods,
	}, nil
}

func (c *Creator) MapType(l int) (*ast.MapType, error) {
	if l == 0 {
		return nil, nil
	}
	Key, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating MapType.Key: %w", err)
	}
	Value, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating MapType.Value: %w", err)
	}

	return &ast.MapType{
		Key:   Key,
		Value: Value,
	}, nil
}

func (c *Creator) StructType(l int) (*ast.StructType, error) {
	if l == 0 {
		return nil, nil
	}
	Fields, err := c.FieldList(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating StructType.Fields: %w", err)
	}

	return &ast.StructType{
		Incomplete: false,
		Fields:     Fields,
	}, nil
}

func (c *Creator) ParenExprForType(l int) (*ast.ParenExpr, error) {
	if l == 0 {
		return nil, nil
	}
	X, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating ParenExprForType.X: %w", err)
	}

	return &ast.ParenExpr{
		X: X,
	}, nil
}

// should cover needs for:
//   - package.Type such as: types.NodeType
//     -
func (c *Creator) SelectorExprForType(l int) (*ast.SelectorExpr, error) {
	if l == 0 {
		return nil, nil
	}
	X, err := c.Expr(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SelectorExprForType.X: %w", err)
	}
	Sel, err := c.Ident(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating SelectorExprForType.Sel: %w", err)
	}

	return &ast.SelectorExpr{
		X:   X,
		Sel: Sel,
	}, nil
}

func (c *Creator) StarExprForType(l int) (*ast.StarExpr, error) {
	if l == 0 {
		return nil, nil
	}
	X, err := c.Type(l - 1)
	if err != nil {
		return nil, fmt.Errorf("generating StarExprForType.X: %w", err)
	}

	return &ast.StarExpr{
		X: X,
	}, nil
}
