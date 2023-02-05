package copy

import (
	"go/ast"
)

func Comment(src *ast.Comment) *ast.Comment {
	dst := *src
	return &dst
}

func CommentGroup(src *ast.CommentGroup) *ast.CommentGroup {
	dst := *src
	return &dst
}

func Field(src *ast.Field) *ast.Field {
	dst := *src
	dst.Names = IdentSlice(src.Names)
	return &dst
}

func FieldList(src *ast.FieldList) *ast.FieldList {
	dst := *src
	dst.List = FieldSlice(src.List)
	return &dst
}

func BadExpr(src *ast.BadExpr) *ast.BadExpr {
	dst := *src
	return &dst
}

func Ident(src *ast.Ident) *ast.Ident {
	dst := *src
	dst.Name = src.Name
	return &dst
}

func Ellipsis(src *ast.Ellipsis) *ast.Ellipsis {
	dst := *src
	return &dst
}

func BasicLit(src *ast.BasicLit) *ast.BasicLit {
	dst := *src
	dst.Kind = src.Kind
	return &dst
}

func FuncLit(src *ast.FuncLit) *ast.FuncLit {
	dst := *src
	return &dst
}

func CompositeLit(src *ast.CompositeLit) *ast.CompositeLit {
	dst := *src
	dst.Elts = ExprSlice(src.Elts)
	return &dst
}

func ParenExpr(src *ast.ParenExpr) *ast.ParenExpr {
	dst := *src
	return &dst
}

func SelectorExpr(src *ast.SelectorExpr) *ast.SelectorExpr {
	dst := *src
	return &dst
}

func IndexExpr(src *ast.IndexExpr) *ast.IndexExpr {
	dst := *src
	return &dst
}

func IndexListExpr(src *ast.IndexListExpr) *ast.IndexListExpr {
	dst := *src
	dst.Indices = ExprSlice(src.Indices)
	return &dst
}

func SliceExpr(src *ast.SliceExpr) *ast.SliceExpr {
	dst := *src
	dst.Slice3 = src.Slice3
	return &dst
}

func TypeAssertExpr(src *ast.TypeAssertExpr) *ast.TypeAssertExpr {
	dst := *src
	return &dst
}

func CallExpr(src *ast.CallExpr) *ast.CallExpr {
	dst := *src
	dst.Args = ExprSlice(src.Args)
	return &dst
}

func StarExpr(src *ast.StarExpr) *ast.StarExpr {
	dst := *src
	return &dst
}

func UnaryExpr(src *ast.UnaryExpr) *ast.UnaryExpr {
	dst := *src
	return &dst
}

func BinaryExpr(src *ast.BinaryExpr) *ast.BinaryExpr {
	dst := *src
	return &dst
}

func KeyValueExpr(src *ast.KeyValueExpr) *ast.KeyValueExpr {
	dst := *src
	return &dst
}

func ArrayType(src *ast.ArrayType) *ast.ArrayType {
	dst := *src
	return &dst
}

func StructType(src *ast.StructType) *ast.StructType {
	dst := *src
	return &dst
}

func FuncType(src *ast.FuncType) *ast.FuncType {
	dst := *src
	return &dst
}

func InterfaceType(src *ast.InterfaceType) *ast.InterfaceType {
	dst := *src
	return &dst
}

func MapType(src *ast.MapType) *ast.MapType {
	dst := *src
	return &dst
}

func ChanType(src *ast.ChanType) *ast.ChanType {
	dst := *src
	return &dst
}

func BadStmt(src *ast.BadStmt) *ast.BadStmt {
	dst := *src
	return &dst
}

func DeclStmt(src *ast.DeclStmt) *ast.DeclStmt {
	dst := *src
	return &dst
}

func EmptyStmt(src *ast.EmptyStmt) *ast.EmptyStmt {
	dst := *src
	return &dst
}

func LabeledStmt(src *ast.LabeledStmt) *ast.LabeledStmt {
	dst := *src
	return &dst
}

func ExprStmt(src *ast.ExprStmt) *ast.ExprStmt {
	dst := *src
	return &dst
}

func SendStmt(src *ast.SendStmt) *ast.SendStmt {
	dst := *src
	return &dst
}

func IncDecStmt(src *ast.IncDecStmt) *ast.IncDecStmt {
	dst := *src
	return &dst
}

func AssignStmt(src *ast.AssignStmt) *ast.AssignStmt {
	dst := *src
	dst.Lhs = ExprSlice(src.Lhs)
	dst.Rhs = ExprSlice(src.Rhs)
	return &dst
}

func GoStmt(src *ast.GoStmt) *ast.GoStmt {
	dst := *src
	return &dst
}

func DeferStmt(src *ast.DeferStmt) *ast.DeferStmt {
	dst := *src
	return &dst
}

func ReturnStmt(src *ast.ReturnStmt) *ast.ReturnStmt {
	dst := *src
	dst.Results = ExprSlice(src.Results)
	return &dst
}

func BranchStmt(src *ast.BranchStmt) *ast.BranchStmt {
	dst := *src
	return &dst
}

func BlockStmt(src *ast.BlockStmt) *ast.BlockStmt {
	dst := *src
	dst.List = StmtSlice(src.List)
	return &dst
}

func IfStmt(src *ast.IfStmt) *ast.IfStmt {
	dst := *src
	return &dst
}

func CaseClause(src *ast.CaseClause) *ast.CaseClause {
	dst := *src
	dst.List = ExprSlice(src.List)
	dst.Body = StmtSlice(src.Body)
	return &dst
}

func SwitchStmt(src *ast.SwitchStmt) *ast.SwitchStmt {
	dst := *src
	return &dst
}

func TypeSwitchStmt(src *ast.TypeSwitchStmt) *ast.TypeSwitchStmt {
	dst := *src
	return &dst
}

func CommClause(src *ast.CommClause) *ast.CommClause {
	dst := *src
	dst.Body = StmtSlice(src.Body)
	return &dst
}

func SelectStmt(src *ast.SelectStmt) *ast.SelectStmt {
	dst := *src
	return &dst
}

func ForStmt(src *ast.ForStmt) *ast.ForStmt {
	dst := *src
	return &dst
}

func RangeStmt(src *ast.RangeStmt) *ast.RangeStmt {
	dst := *src
	return &dst
}

func ImportSpec(src *ast.ImportSpec) *ast.ImportSpec {
	dst := *src
	return &dst
}

func ValueSpec(src *ast.ValueSpec) *ast.ValueSpec {
	dst := *src
	dst.Names = IdentSlice(src.Names)
	dst.Values = ExprSlice(src.Values)
	return &dst
}

func TypeSpec(src *ast.TypeSpec) *ast.TypeSpec {
	dst := *src
	return &dst
}

func BadDecl(src *ast.BadDecl) *ast.BadDecl {
	dst := *src
	return &dst
}

func GenDecl(src *ast.GenDecl) *ast.GenDecl {
	dst := *src
	dst.Specs = SpecSlice(src.Specs)
	return &dst
}

func FuncDecl(src *ast.FuncDecl) *ast.FuncDecl {
	dst := *src
	return &dst
}

func File(src *ast.File) *ast.File {
	dst := *src
	dst.Decls = DeclSlice(src.Decls)
	dst.Imports = ImportSpecSlice(src.Imports)
	dst.Unresolved = IdentSlice(src.Unresolved)
	dst.Comments = CommentGroupSlice(src.Comments)
	return &dst
}

func Package(src *ast.Package) *ast.Package {
	dst := *src
	dst.Imports = ObjectMap(src.Imports)
	dst.Files = FileMap(src.Files)
	return &dst
}

func ObjKind(src *ast.ObjKind) *ast.ObjKind {
	dst := &src
	return *dst
}

func Object(src *ast.Object) *ast.Object {
	dst := &src
	return *dst
}

func Scope(src *ast.Scope) *ast.Scope {
	if src == nil {
		return nil
	}
	dest := &ast.Scope{}
	dest.Objects = ObjectMap(src.Objects)
	if src.Outer != nil {
		dest.Outer = Scope(src.Outer)
	}
	return dest
}
