package clean

import (
	"go/ast"
	"go/token"
)

func Comment(src *ast.Comment) *ast.Comment {
	if src == nil {
		return nil
	}
	return &ast.Comment{
		Slash: token.Pos(src.Slash),
		Text:  string(src.Text),
	}
}

func CommentGroup(src *ast.CommentGroup) *ast.CommentGroup {
	if src == nil {
		return nil
	}
	return &ast.CommentGroup{
		List: CommentSlice(src.List),
	}
}

func Field(src *ast.Field) *ast.Field {
	if src == nil {
		return nil
	}
	return &ast.Field{
		Doc:     CommentGroup(src.Doc),
		Names:   IdentSlice(src.Names),
		Type:    Expr(src.Type),
		Tag:     BasicLit(src.Tag),
		Comment: CommentGroup(src.Comment),
	}
}

func FieldList(src *ast.FieldList) *ast.FieldList {
	if src == nil {
		return nil
	}
	return &ast.FieldList{
		Opening: token.Pos(src.Opening),
		List:    FieldSlice(src.List),
		Closing: token.Pos(src.Closing),
	}
}

func BadExpr(src *ast.BadExpr) *ast.BadExpr {
	if src == nil {
		return nil
	}
	return &ast.BadExpr{
		From: token.Pos(src.From),
		To:   token.Pos(src.To),
	}
}

func Ident(src *ast.Ident) *ast.Ident {
	if src == nil {
		return nil
	}
	return &ast.Ident{
		NamePos: token.Pos(src.NamePos),
		Name:    string(src.Name),
		Obj:     nil,
	}
}

func Ellipsis(src *ast.Ellipsis) *ast.Ellipsis {
	if src == nil {
		return nil
	}
	return &ast.Ellipsis{
		Ellipsis: token.Pos(src.Ellipsis),
		Elt:      Expr(src.Elt),
	}
}

func BasicLit(src *ast.BasicLit) *ast.BasicLit {
	if src == nil {
		return nil
	}
	return &ast.BasicLit{
		ValuePos: token.Pos(src.ValuePos),
		Kind:     token.Token(src.Kind),
		Value:    string(src.Value),
	}
}

func FuncLit(src *ast.FuncLit) *ast.FuncLit {
	if src == nil {
		return nil
	}
	return &ast.FuncLit{
		Type: FuncType(src.Type),
		Body: BlockStmt(src.Body),
	}
}

func CompositeLit(src *ast.CompositeLit) *ast.CompositeLit {
	if src == nil {
		return nil
	}
	return &ast.CompositeLit{
		Type:       Expr(src.Type),
		Lbrace:     token.Pos(src.Lbrace),
		Elts:       ExprSlice(src.Elts),
		Rbrace:     token.Pos(src.Rbrace),
		Incomplete: bool(src.Incomplete),
	}
}

func ParenExpr(src *ast.ParenExpr) *ast.ParenExpr {
	if src == nil {
		return nil
	}
	return &ast.ParenExpr{
		Lparen: token.Pos(src.Lparen),
		X:      Expr(src.X),
		Rparen: token.Pos(src.Rparen),
	}
}

func SelectorExpr(src *ast.SelectorExpr) *ast.SelectorExpr {
	if src == nil {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(src.X),
		Sel: Ident(src.Sel),
	}
}

func IndexExpr(src *ast.IndexExpr) *ast.IndexExpr {
	if src == nil {
		return nil
	}
	return &ast.IndexExpr{
		X:      Expr(src.X),
		Lbrack: token.Pos(src.Lbrack),
		Index:  Expr(src.Index),
		Rbrack: token.Pos(src.Rbrack),
	}
}

func IndexListExpr(src *ast.IndexListExpr) *ast.IndexListExpr {
	if src == nil {
		return nil
	}
	return &ast.IndexListExpr{
		X:       Expr(src.X),
		Lbrack:  token.Pos(src.Lbrack),
		Indices: ExprSlice(src.Indices),
		Rbrack:  token.Pos(src.Rbrack),
	}
}

func SliceExpr(src *ast.SliceExpr) *ast.SliceExpr {
	if src == nil {
		return nil
	}
	return &ast.SliceExpr{
		X:      Expr(src.X),
		Lbrack: token.Pos(src.Lbrack),
		Low:    Expr(src.Low),
		High:   Expr(src.High),
		Max:    Expr(src.Max),
		Slice3: bool(src.Slice3),
		Rbrack: token.Pos(src.Rbrack),
	}
}

func TypeAssertExpr(src *ast.TypeAssertExpr) *ast.TypeAssertExpr {
	if src == nil {
		return nil
	}
	return &ast.TypeAssertExpr{
		X:      Expr(src.X),
		Lparen: token.Pos(src.Lparen),
		Type:   Expr(src.Type),
		Rparen: token.Pos(src.Rparen),
	}
}

func CallExpr(src *ast.CallExpr) *ast.CallExpr {
	if src == nil {
		return nil
	}
	return &ast.CallExpr{
		Fun:      Expr(src.Fun),
		Lparen:   token.Pos(src.Lparen),
		Args:     ExprSlice(src.Args),
		Ellipsis: token.Pos(src.Ellipsis),
		Rparen:   token.Pos(src.Rparen),
	}
}

func StarExpr(src *ast.StarExpr) *ast.StarExpr {
	if src == nil {
		return nil
	}
	return &ast.StarExpr{
		Star: token.Pos(src.Star),
		X:    Expr(src.X),
	}
}

func UnaryExpr(src *ast.UnaryExpr) *ast.UnaryExpr {
	if src == nil {
		return nil
	}
	return &ast.UnaryExpr{
		OpPos: token.Pos(src.OpPos),
		Op:    token.Token(src.Op),
		X:     Expr(src.X),
	}
}

func BinaryExpr(src *ast.BinaryExpr) *ast.BinaryExpr {
	if src == nil {
		return nil
	}
	return &ast.BinaryExpr{
		X:     Expr(src.X),
		OpPos: token.Pos(src.OpPos),
		Op:    token.Token(src.Op),
		Y:     Expr(src.Y),
	}
}

func KeyValueExpr(src *ast.KeyValueExpr) *ast.KeyValueExpr {
	if src == nil {
		return nil
	}
	return &ast.KeyValueExpr{
		Key:   Expr(src.Key),
		Colon: token.Pos(src.Colon),
		Value: Expr(src.Value),
	}
}

func ArrayType(src *ast.ArrayType) *ast.ArrayType {
	if src == nil {
		return nil
	}
	return &ast.ArrayType{
		Lbrack: token.Pos(src.Lbrack),
		Len:    Expr(src.Len),
		Elt:    Expr(src.Elt),
	}
}

func StructType(src *ast.StructType) *ast.StructType {
	if src == nil {
		return nil
	}
	return &ast.StructType{
		Struct:     token.Pos(src.Struct),
		Fields:     FieldList(src.Fields),
		Incomplete: bool(src.Incomplete),
	}
}

func FuncType(src *ast.FuncType) *ast.FuncType {
	if src == nil {
		return nil
	}
	return &ast.FuncType{
		Func:       token.Pos(src.Func),
		TypeParams: FieldList(src.TypeParams),
		Params:     FieldList(src.Params),
		Results:    FieldList(src.Results),
	}
}

func InterfaceType(src *ast.InterfaceType) *ast.InterfaceType {
	if src == nil {
		return nil
	}
	return &ast.InterfaceType{
		Interface:  token.Pos(src.Interface),
		Methods:    FieldList(src.Methods),
		Incomplete: bool(src.Incomplete),
	}
}

func MapType(src *ast.MapType) *ast.MapType {
	if src == nil {
		return nil
	}
	return &ast.MapType{
		Map:   token.Pos(src.Map),
		Key:   Expr(src.Key),
		Value: Expr(src.Value),
	}
}

func ChanType(src *ast.ChanType) *ast.ChanType {
	if src == nil {
		return nil
	}
	return &ast.ChanType{
		Begin: token.Pos(src.Begin),
		Arrow: token.Pos(src.Arrow),
		Dir:   ast.ChanDir(src.Dir),
		Value: Expr(src.Value),
	}
}

func BadStmt(src *ast.BadStmt) *ast.BadStmt {
	if src == nil {
		return nil
	}
	return &ast.BadStmt{
		From: token.Pos(src.From),
		To:   token.Pos(src.To),
	}
}

func DeclStmt(src *ast.DeclStmt) *ast.DeclStmt {
	if src == nil {
		return nil
	}
	return &ast.DeclStmt{
		Decl: Decl(src.Decl),
	}
}

func EmptyStmt(src *ast.EmptyStmt) *ast.EmptyStmt {
	if src == nil {
		return nil
	}
	return &ast.EmptyStmt{
		Semicolon: token.Pos(src.Semicolon),
		Implicit:  bool(src.Implicit),
	}
}

func LabeledStmt(src *ast.LabeledStmt) *ast.LabeledStmt {
	if src == nil {
		return nil
	}
	return &ast.LabeledStmt{
		Label: Ident(src.Label),
		Colon: token.Pos(src.Colon),
		Stmt:  Stmt(src.Stmt),
	}
}

func ExprStmt(src *ast.ExprStmt) *ast.ExprStmt {
	if src == nil {
		return nil
	}
	return &ast.ExprStmt{
		X: Expr(src.X),
	}
}

func SendStmt(src *ast.SendStmt) *ast.SendStmt {
	if src == nil {
		return nil
	}
	return &ast.SendStmt{
		Chan:  Expr(src.Chan),
		Arrow: token.Pos(src.Arrow),
		Value: Expr(src.Value),
	}
}

func IncDecStmt(src *ast.IncDecStmt) *ast.IncDecStmt {
	if src == nil {
		return nil
	}
	return &ast.IncDecStmt{
		X:      Expr(src.X),
		TokPos: token.Pos(src.TokPos),
		Tok:    token.Token(src.Tok),
	}
}

func AssignStmt(src *ast.AssignStmt) *ast.AssignStmt {
	if src == nil {
		return nil
	}
	return &ast.AssignStmt{
		Lhs:    ExprSlice(src.Lhs),
		TokPos: token.Pos(src.TokPos),
		Tok:    token.Token(src.Tok),
		Rhs:    ExprSlice(src.Rhs),
	}
}

func GoStmt(src *ast.GoStmt) *ast.GoStmt {
	if src == nil {
		return nil
	}
	return &ast.GoStmt{
		Go:   token.Pos(src.Go),
		Call: CallExpr(src.Call),
	}
}

func DeferStmt(src *ast.DeferStmt) *ast.DeferStmt {
	if src == nil {
		return nil
	}
	return &ast.DeferStmt{
		Defer: token.Pos(src.Defer),
		Call:  CallExpr(src.Call),
	}
}

func ReturnStmt(src *ast.ReturnStmt) *ast.ReturnStmt {
	if src == nil {
		return nil
	}
	return &ast.ReturnStmt{
		Return:  token.Pos(src.Return),
		Results: ExprSlice(src.Results),
	}
}

func BranchStmt(src *ast.BranchStmt) *ast.BranchStmt {
	if src == nil {
		return nil
	}
	return &ast.BranchStmt{
		TokPos: token.Pos(src.TokPos),
		Tok:    token.Token(src.Tok),
		Label:  Ident(src.Label),
	}
}

func BlockStmt(src *ast.BlockStmt) *ast.BlockStmt {
	if src == nil {
		return nil
	}
	return &ast.BlockStmt{
		Lbrace: token.Pos(src.Lbrace),
		List:   StmtSlice(src.List),
		Rbrace: token.Pos(src.Rbrace),
	}
}

func IfStmt(src *ast.IfStmt) *ast.IfStmt {
	if src == nil {
		return nil
	}
	return &ast.IfStmt{
		If:   token.Pos(src.If),
		Init: Stmt(src.Init),
		Cond: Expr(src.Cond),
		Body: BlockStmt(src.Body),
		Else: Stmt(src.Else),
	}
}

func CaseClause(src *ast.CaseClause) *ast.CaseClause {
	if src == nil {
		return nil
	}
	return &ast.CaseClause{
		Case:  token.Pos(src.Case),
		List:  ExprSlice(src.List),
		Colon: token.Pos(src.Colon),
		Body:  StmtSlice(src.Body),
	}
}

func SwitchStmt(src *ast.SwitchStmt) *ast.SwitchStmt {
	if src == nil {
		return nil
	}
	return &ast.SwitchStmt{
		Switch: token.Pos(src.Switch),
		Init:   Stmt(src.Init),
		Tag:    Expr(src.Tag),
		Body:   BlockStmt(src.Body),
	}
}

func TypeSwitchStmt(src *ast.TypeSwitchStmt) *ast.TypeSwitchStmt {
	if src == nil {
		return nil
	}
	return &ast.TypeSwitchStmt{
		Switch: token.Pos(src.Switch),
		Init:   Stmt(src.Init),
		Assign: Stmt(src.Assign),
		Body:   BlockStmt(src.Body),
	}
}

func CommClause(src *ast.CommClause) *ast.CommClause {
	if src == nil {
		return nil
	}
	return &ast.CommClause{
		Case:  token.Pos(src.Case),
		Comm:  Stmt(src.Comm),
		Colon: token.Pos(src.Colon),
		Body:  StmtSlice(src.Body),
	}
}

func SelectStmt(src *ast.SelectStmt) *ast.SelectStmt {
	if src == nil {
		return nil
	}
	return &ast.SelectStmt{
		Select: token.Pos(src.Select),
		Body:   BlockStmt(src.Body),
	}
}

func ForStmt(src *ast.ForStmt) *ast.ForStmt {
	if src == nil {
		return nil
	}
	return &ast.ForStmt{
		For:  token.Pos(src.For),
		Init: Stmt(src.Init),
		Cond: Expr(src.Cond),
		Post: Stmt(src.Post),
		Body: BlockStmt(src.Body),
	}
}

func RangeStmt(src *ast.RangeStmt) *ast.RangeStmt {
	if src == nil {
		return nil
	}
	return &ast.RangeStmt{
		For:    token.Pos(src.For),
		Key:    Expr(src.Key),
		Value:  Expr(src.Value),
		TokPos: token.Pos(src.TokPos),
		Tok:    token.Token(src.Tok),
		X:      Expr(src.X),
		Body:   BlockStmt(src.Body),
	}
}

func ImportSpec(src *ast.ImportSpec) *ast.ImportSpec {
	if src == nil {
		return nil
	}
	return &ast.ImportSpec{
		Doc:     CommentGroup(src.Doc),
		Name:    Ident(src.Name),
		Path:    BasicLit(src.Path),
		Comment: CommentGroup(src.Comment),
		EndPos:  token.Pos(src.EndPos),
	}
}

func ValueSpec(src *ast.ValueSpec) *ast.ValueSpec {
	if src == nil {
		return nil
	}
	return &ast.ValueSpec{
		Doc:     CommentGroup(src.Doc),
		Names:   IdentSlice(src.Names),
		Type:    Expr(src.Type),
		Values:  ExprSlice(src.Values),
		Comment: CommentGroup(src.Comment),
	}
}

func TypeSpec(src *ast.TypeSpec) *ast.TypeSpec {
	if src == nil {
		return nil
	}
	return &ast.TypeSpec{
		Doc:        CommentGroup(src.Doc),
		Name:       Ident(src.Name),
		TypeParams: FieldList(src.TypeParams),
		Assign:     token.Pos(src.Assign),
		Type:       Expr(src.Type),
		Comment:    CommentGroup(src.Comment),
	}
}

func BadDecl(src *ast.BadDecl) *ast.BadDecl {
	if src == nil {
		return nil
	}
	return &ast.BadDecl{
		From: token.Pos(src.From),
		To:   token.Pos(src.To),
	}
}

func GenDecl(src *ast.GenDecl) *ast.GenDecl {
	if src == nil {
		return nil
	}
	return &ast.GenDecl{
		Doc:    CommentGroup(src.Doc),
		TokPos: token.Pos(src.TokPos),
		Tok:    token.Token(src.Tok),
		Lparen: token.Pos(src.Lparen),
		Specs:  SpecSlice(src.Specs),
		Rparen: token.Pos(src.Rparen),
	}
}

func FuncDecl(src *ast.FuncDecl) *ast.FuncDecl {
	if src == nil {
		return nil
	}
	return &ast.FuncDecl{
		Doc:  CommentGroup(src.Doc),
		Recv: FieldList(src.Recv),
		Name: Ident(src.Name),
		Type: FuncType(src.Type),
		Body: BlockStmt(src.Body),
	}
}

func File(src *ast.File) *ast.File {
	if src == nil {
		return nil
	}
	return &ast.File{
		Doc:        CommentGroup(src.Doc),
		Package:    token.Pos(src.Package),
		Name:       Ident(src.Name),
		Decls:      DeclSlice(src.Decls),
		Scope:      nil,
		Imports:    nil,
		Unresolved: nil,
		Comments:   CommentGroupSlice(src.Comments),
	}
}

func Package(src *ast.Package) *ast.Package {
	if src == nil {
		return nil
	}
	return &ast.Package{
		Name:    string(src.Name),
		Scope:   nil,
		Imports: nil,
		Files:   FileMap(src.Files),
	}
}
