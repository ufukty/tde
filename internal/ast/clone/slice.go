package clone

import "go/ast"

func commentSlice(src []*ast.Comment) []*ast.Comment {
	if src == nil {
		return nil
	}
	dst := []*ast.Comment{}
	for _, srcItem := range src {
		dst = append(dst, Comment(srcItem))
	}
	return dst
}

func commentGroupSlice(src []*ast.CommentGroup) []*ast.CommentGroup {
	if src == nil {
		return nil
	}
	dst := []*ast.CommentGroup{}
	for _, srcItem := range src {
		dst = append(dst, CommentGroup(srcItem))
	}
	return dst
}

func declSlice(src []ast.Decl) []ast.Decl {
	if src == nil {
		return nil
	}
	dst := []ast.Decl{}
	for _, srcItem := range src {
		dst = append(dst, Decl(srcItem))
	}
	return dst
}

func exprSlice(src []ast.Expr) []ast.Expr {
	if src == nil {
		return nil
	}
	dst := []ast.Expr{}
	for _, srcItem := range src {
		dst = append(dst, Expr(srcItem))
	}
	return dst
}

func fieldSlice(src []*ast.Field) []*ast.Field {
	if src == nil {
		return nil
	}
	dst := []*ast.Field{}
	for _, srcItem := range src {
		dst = append(dst, Field(srcItem))
	}
	return dst
}

func identSlice(src []*ast.Ident) []*ast.Ident {
	if src == nil {
		return nil
	}
	dst := []*ast.Ident{}
	for _, srcItem := range src {
		dst = append(dst, Ident(srcItem))
	}
	return dst
}

func importSpecSlice(src []*ast.ImportSpec) []*ast.ImportSpec {
	if src == nil {
		return nil
	}
	dst := []*ast.ImportSpec{}
	for _, srcItem := range src {
		dst = append(dst, ImportSpec(srcItem))
	}
	return dst
}

func specSlice(src []ast.Spec) []ast.Spec {
	if src == nil {
		return nil
	}
	dst := []ast.Spec{}
	for _, srcItem := range src {
		dst = append(dst, Spec(srcItem))
	}
	return dst
}

func stmtSlice(src []ast.Stmt) []ast.Stmt {
	if src == nil {
		return nil
	}
	dst := []ast.Stmt{}
	for _, srcItem := range src {
		dst = append(dst, Stmt(srcItem))
	}
	return dst
}
