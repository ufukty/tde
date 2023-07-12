package clean

import "go/ast"

func CommentSlice(src []*ast.Comment) []*ast.Comment {
	if src == nil {
		return nil
	}
	dst := []*ast.Comment{}
	for _, srcItem := range src {
		dst = append(dst, Comment(srcItem))
	}
	return dst
}

func CommentGroupSlice(src []*ast.CommentGroup) []*ast.CommentGroup {
	if src == nil {
		return nil
	}
	dst := []*ast.CommentGroup{}
	for _, srcItem := range src {
		dst = append(dst, CommentGroup(srcItem))
	}
	return dst
}

func DeclSlice(src []ast.Decl) []ast.Decl {
	if src == nil {
		return nil
	}
	dst := []ast.Decl{}
	for _, srcItem := range src {
		dst = append(dst, Decl(srcItem))
	}
	return dst
}

func ExprSlice(src []ast.Expr) []ast.Expr {
	if src == nil {
		return nil
	}
	dst := []ast.Expr{}
	for _, srcItem := range src {
		dst = append(dst, Expr(srcItem))
	}
	return dst
}

func FieldSlice(src []*ast.Field) []*ast.Field {
	if src == nil {
		return nil
	}
	dst := []*ast.Field{}
	for _, srcItem := range src {
		dst = append(dst, Field(srcItem))
	}
	return dst
}

func IdentSlice(src []*ast.Ident) []*ast.Ident {
	if src == nil {
		return nil
	}
	dst := []*ast.Ident{}
	for _, srcItem := range src {
		dst = append(dst, Ident(srcItem))
	}
	return dst
}

func ImportSpecSlice(src []*ast.ImportSpec) []*ast.ImportSpec {
	if src == nil {
		return nil
	}
	dst := []*ast.ImportSpec{}
	for _, srcItem := range src {
		dst = append(dst, ImportSpec(srcItem))
	}
	return dst
}

func SpecSlice(src []ast.Spec) []ast.Spec {
	if src == nil {
		return nil
	}
	dst := []ast.Spec{}
	for _, srcItem := range src {
		dst = append(dst, Spec(srcItem))
	}
	return dst
}

func StmtSlice(src []ast.Stmt) []ast.Stmt {
	if src == nil {
		return nil
	}
	dst := []ast.Stmt{}
	for _, srcItem := range src {
		dst = append(dst, Stmt(srcItem))
	}
	return dst
}
