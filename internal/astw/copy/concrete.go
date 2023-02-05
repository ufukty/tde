package copy

import (
	"go/ast"

	"golang.org/x/exp/maps"
)

func Comment(src *ast.Comment) *ast.Comment {
	dst := &ast.Comment{}

	if src.Text != "" {
		dst.Text = src.Text
	}

	return dst
}

func CommentGroup(src *ast.CommentGroup) *ast.CommentGroup {
	dst := &ast.CommentGroup{}

	if src.List != nil {
		dst.List = CommentSlice(src.List)
	}

	return dst
}

func Field(src *ast.Field) *ast.Field {
	dst := &ast.Field{}

	if src.Names != nil {
		dst.Names = IdentSlice(src.Names)
	}
	if src.Type != nil {
		dst.Type = Expr(src.Type)
	}
	if src.Tag != nil {
		dst.Tag = BasicLit(src.Tag)
	}

	return dst
}

func FieldList(src *ast.FieldList) *ast.FieldList {
	dst := &ast.FieldList{}

	if src.List != nil {
		dst.List = FieldSlice(src.List)
	}

	return dst
}

func BadExpr(src *ast.BadExpr) *ast.BadExpr {
	dst := &ast.BadExpr{}

	return dst
}

func Ident(src *ast.Ident) *ast.Ident {
	dst := &ast.Ident{}

	dst.Name = src.Name

	if src.Obj != nil {
		dst.Obj = Object(src.Obj)
	}

	return dst
}

func Ellipsis(src *ast.Ellipsis) *ast.Ellipsis {
	dst := &ast.Ellipsis{}

	if src.Elt != nil {
		dst.Elt = Expr(src.Elt)
	}

	return dst
}

func BasicLit(src *ast.BasicLit) *ast.BasicLit {
	dst := &ast.BasicLit{}

	dst.Kind = src.Kind

	return dst
}

func FuncLit(src *ast.FuncLit) *ast.FuncLit {
	dst := &ast.FuncLit{}

	if src.Type != nil {
		dst.Type = FuncType(src.Type)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func CompositeLit(src *ast.CompositeLit) *ast.CompositeLit {
	dst := &ast.CompositeLit{}

	dst.Incomplete = src.Incomplete

	if src.Type != nil {
		dst.Type = Expr(src.Type)
	}
	if src.Elts != nil {
		dst.Elts = ExprSlice(src.Elts)
	}

	return dst
}

func ParenExpr(src *ast.ParenExpr) *ast.ParenExpr {
	dst := &ast.ParenExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}

	return dst
}

func SelectorExpr(src *ast.SelectorExpr) *ast.SelectorExpr {
	dst := &ast.SelectorExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Sel != nil {
		dst.Sel = Ident(src.Sel)
	}

	return dst
}

func IndexExpr(src *ast.IndexExpr) *ast.IndexExpr {
	dst := &ast.IndexExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Index != nil {
		dst.Index = Expr(src.Index)
	}

	return dst
}

func IndexListExpr(src *ast.IndexListExpr) *ast.IndexListExpr {
	dst := &ast.IndexListExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Indices != nil {
		dst.Indices = ExprSlice(src.Indices)
	}

	return dst
}

func SliceExpr(src *ast.SliceExpr) *ast.SliceExpr {
	dst := &ast.SliceExpr{}

	dst.Slice3 = src.Slice3

	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Low != nil {
		dst.Low = Expr(src.Low)
	}
	if src.High != nil {
		dst.High = Expr(src.High)
	}
	if src.Max != nil {
		dst.Max = Expr(src.Max)
	}

	return dst
}

func TypeAssertExpr(src *ast.TypeAssertExpr) *ast.TypeAssertExpr {
	dst := &ast.TypeAssertExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Type != nil {
		dst.Type = Expr(src.Type)
	}

	return dst
}

func CallExpr(src *ast.CallExpr) *ast.CallExpr {
	dst := &ast.CallExpr{}

	if src.Fun != nil {
		dst.Fun = Expr(src.Fun)
	}
	if src.Args != nil {
		dst.Args = ExprSlice(src.Args)
	}

	return dst
}

func StarExpr(src *ast.StarExpr) *ast.StarExpr {
	dst := &ast.StarExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}

	return dst
}

func UnaryExpr(src *ast.UnaryExpr) *ast.UnaryExpr {
	dst := &ast.UnaryExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}

	return dst
}

func BinaryExpr(src *ast.BinaryExpr) *ast.BinaryExpr {
	dst := &ast.BinaryExpr{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Y != nil {
		dst.Y = Expr(src.Y)
	}

	return dst
}

func KeyValueExpr(src *ast.KeyValueExpr) *ast.KeyValueExpr {
	dst := &ast.KeyValueExpr{}

	if src.Key != nil {
		dst.Key = Expr(src.Key)
	}
	if src.Value != nil {
		dst.Value = Expr(src.Value)
	}

	return dst
}

func ArrayType(src *ast.ArrayType) *ast.ArrayType {
	dst := &ast.ArrayType{}

	if src.Len != nil {
		dst.Len = Expr(src.Len)
	}
	if src.Elt != nil {
		dst.Elt = Expr(src.Elt)
	}

	return dst
}

func StructType(src *ast.StructType) *ast.StructType {
	dst := &ast.StructType{}

	if src.Fields != nil {
		dst.Fields = FieldList(src.Fields)
	}

	return dst
}

func FuncType(src *ast.FuncType) *ast.FuncType {
	dst := &ast.FuncType{}

	if src.TypeParams != nil {
		dst.TypeParams = FieldList(src.TypeParams)
	}
	if src.Params != nil {
		dst.Params = FieldList(src.Params)
	}
	if src.Results != nil {
		dst.Results = FieldList(src.Results)
	}

	return dst
}

func InterfaceType(src *ast.InterfaceType) *ast.InterfaceType {
	dst := &ast.InterfaceType{}

	if src.Methods != nil {
		dst.Methods = FieldList(src.Methods)
	}

	return dst
}

func MapType(src *ast.MapType) *ast.MapType {
	dst := &ast.MapType{}

	if src.Key != nil {
		dst.Key = Expr(src.Key)
	}
	if src.Value != nil {
		dst.Value = Expr(src.Value)
	}

	return dst
}

func ChanType(src *ast.ChanType) *ast.ChanType {
	dst := &ast.ChanType{}

	if src.Value != nil {
		dst.Value = Expr(src.Value)
	}

	return dst
}

func BadStmt(src *ast.BadStmt) *ast.BadStmt {
	dst := &ast.BadStmt{}

	return dst
}

func DeclStmt(src *ast.DeclStmt) *ast.DeclStmt {
	dst := &ast.DeclStmt{}

	if src.Decl != nil {
		dst.Decl = Decl(src.Decl)
	}

	return dst
}

func EmptyStmt(src *ast.EmptyStmt) *ast.EmptyStmt {
	dst := &ast.EmptyStmt{}

	return dst
}

func LabeledStmt(src *ast.LabeledStmt) *ast.LabeledStmt {
	dst := &ast.LabeledStmt{}

	if src.Label != nil {
		dst.Label = Ident(src.Label)
	}
	if src.Stmt != nil {
		dst.Stmt = Stmt(src.Stmt)
	}

	return dst
}

func ExprStmt(src *ast.ExprStmt) *ast.ExprStmt {
	dst := &ast.ExprStmt{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}

	return dst
}

func SendStmt(src *ast.SendStmt) *ast.SendStmt {
	dst := &ast.SendStmt{}

	if src.Chan != nil {
		dst.Chan = Expr(src.Chan)
	}
	if src.Value != nil {
		dst.Value = Expr(src.Value)
	}

	return dst
}

func IncDecStmt(src *ast.IncDecStmt) *ast.IncDecStmt {
	dst := &ast.IncDecStmt{}

	if src.X != nil {
		dst.X = Expr(src.X)
	}

	return dst
}

func AssignStmt(src *ast.AssignStmt) *ast.AssignStmt {
	dst := &ast.AssignStmt{}

	if src.Lhs != nil {
		dst.Lhs = ExprSlice(src.Lhs)
	}
	if src.Rhs != nil {
		dst.Rhs = ExprSlice(src.Rhs)
	}

	return dst
}

func GoStmt(src *ast.GoStmt) *ast.GoStmt {
	dst := &ast.GoStmt{}

	if src.Call != nil {
		dst.Call = CallExpr(src.Call)
	}

	return dst
}

func DeferStmt(src *ast.DeferStmt) *ast.DeferStmt {
	dst := &ast.DeferStmt{}

	if src.Call != nil {
		dst.Call = CallExpr(src.Call)
	}

	return dst
}

func ReturnStmt(src *ast.ReturnStmt) *ast.ReturnStmt {
	dst := &ast.ReturnStmt{}

	if src.Results != nil {
		dst.Results = ExprSlice(src.Results)
	}

	return dst
}

func BranchStmt(src *ast.BranchStmt) *ast.BranchStmt {
	dst := &ast.BranchStmt{}

	if src.Label != nil {
		dst.Label = Ident(src.Label)
	}

	return dst
}

func BlockStmt(src *ast.BlockStmt) *ast.BlockStmt {
	dst := &ast.BlockStmt{}

	if src.List != nil {
		dst.List = StmtSlice(src.List)
	}

	return dst
}

func IfStmt(src *ast.IfStmt) *ast.IfStmt {
	dst := &ast.IfStmt{}

	if src.Init != nil {
		dst.Init = Stmt(src.Init)
	}
	if src.Cond != nil {
		dst.Cond = Expr(src.Cond)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}
	if src.Else != nil {
		dst.Else = Stmt(src.Else)
	}

	return dst
}

func CaseClause(src *ast.CaseClause) *ast.CaseClause {
	dst := &ast.CaseClause{}

	if src.List != nil {
		dst.List = ExprSlice(src.List)
	}
	if src.Body != nil {
		dst.Body = StmtSlice(src.Body)
	}

	return dst
}

func SwitchStmt(src *ast.SwitchStmt) *ast.SwitchStmt {
	dst := &ast.SwitchStmt{}

	if src.Init != nil {
		dst.Init = Stmt(src.Init)
	}
	if src.Tag != nil {
		dst.Tag = Expr(src.Tag)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func TypeSwitchStmt(src *ast.TypeSwitchStmt) *ast.TypeSwitchStmt {
	dst := &ast.TypeSwitchStmt{}

	if src.Init != nil {
		dst.Init = Stmt(src.Init)
	}
	if src.Assign != nil {
		dst.Assign = Stmt(src.Assign)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func CommClause(src *ast.CommClause) *ast.CommClause {
	dst := &ast.CommClause{}

	if src.Comm != nil {
		dst.Comm = Stmt(src.Comm)
	}
	if src.Body != nil {
		dst.Body = StmtSlice(src.Body)
	}

	return dst
}

func SelectStmt(src *ast.SelectStmt) *ast.SelectStmt {
	dst := &ast.SelectStmt{}

	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func ForStmt(src *ast.ForStmt) *ast.ForStmt {
	dst := &ast.ForStmt{}

	if src.Init != nil {
		dst.Init = Stmt(src.Init)
	}
	if src.Cond != nil {
		dst.Cond = Expr(src.Cond)
	}
	if src.Post != nil {
		dst.Post = Stmt(src.Post)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func RangeStmt(src *ast.RangeStmt) *ast.RangeStmt {
	dst := &ast.RangeStmt{}

	if src.Key != nil {
		dst.Key = Expr(src.Key)
	}
	if src.Value != nil {
		dst.Value = Expr(src.Value)
	}
	if src.X != nil {
		dst.X = Expr(src.X)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func ImportSpec(src *ast.ImportSpec) *ast.ImportSpec {
	dst := &ast.ImportSpec{}

	if src.Comment != nil {
		dst.Comment = CommentGroup(src.Comment)
	}

	if src.Name != nil {
		dst.Name = Ident(src.Name)
	}
	if src.Path != nil {
		dst.Path = BasicLit(src.Path)
	}

	return dst
}

func ValueSpec(src *ast.ValueSpec) *ast.ValueSpec {
	dst := &ast.ValueSpec{}

	if src.Comment != nil {
		dst.Comment = CommentGroup(src.Comment)
	}

	if src.Names != nil {
		dst.Names = IdentSlice(src.Names)
	}
	if src.Type != nil {
		dst.Type = Expr(src.Type)
	}
	if src.Values != nil {
		dst.Values = ExprSlice(src.Values)
	}

	return dst
}

func TypeSpec(src *ast.TypeSpec) *ast.TypeSpec {
	dst := &ast.TypeSpec{}

	if src.Comment != nil {
		dst.Comment = CommentGroup(src.Comment)
	}

	if src.Name != nil {
		dst.Name = Ident(src.Name)
	}
	if src.TypeParams != nil {
		dst.TypeParams = FieldList(src.TypeParams)
	}
	if src.Type != nil {
		dst.Type = Expr(src.Type)
	}

	return dst
}

func BadDecl(src *ast.BadDecl) *ast.BadDecl {
	dst := &ast.BadDecl{}

	return dst
}

func GenDecl(src *ast.GenDecl) *ast.GenDecl {
	dst := &ast.GenDecl{}

	if src.Specs != nil {
		dst.Specs = SpecSlice(src.Specs)
	}

	return dst
}

func FuncDecl(src *ast.FuncDecl) *ast.FuncDecl {
	dst := &ast.FuncDecl{}

	if src.Recv != nil {
		dst.Recv = FieldList(src.Recv)
	}
	if src.Name != nil {
		dst.Name = Ident(src.Name)
	}
	if src.Type != nil {
		dst.Type = FuncType(src.Type)
	}
	if src.Body != nil {
		dst.Body = BlockStmt(src.Body)
	}

	return dst
}

func File(src *ast.File) *ast.File {
	dst := &ast.File{}

	if src.Name != nil {
		dst.Name = Ident(src.Name)
	}
	if src.Decls != nil {
		dst.Decls = DeclSlice(src.Decls)
	}
	if src.Imports != nil {
		dst.Imports = ImportSpecSlice(src.Imports)
	}
	if src.Unresolved != nil {
		dst.Unresolved = IdentSlice(src.Unresolved)
	}
	// {n.Comments, types.CommentGroupSlice},

	return dst
}

func Package(src *ast.Package) *ast.Package {
	dst := &ast.Package{}

	if src.Name != "" {
		dst.Name = src.Name
	}
	if src.Scope != nil {
		dst.Scope = Scope(src.Scope)
	}
	if src.Imports != nil {
		dst.Imports = maps.Clone(src.Imports)
	}
	if src.Files != nil {
		dst.Files = maps.Clone(src.Files)
	}

	return dst
}
