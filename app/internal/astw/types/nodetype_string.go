// Code generated by "stringer -type=NodeType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ArrayType-0]
	_ = x[AssignStmt-1]
	_ = x[BadDecl-2]
	_ = x[BadExpr-3]
	_ = x[BadStmt-4]
	_ = x[BasicLit-5]
	_ = x[BinaryExpr-6]
	_ = x[BlockStmt-7]
	_ = x[BranchStmt-8]
	_ = x[CallExpr-9]
	_ = x[CaseClause-10]
	_ = x[ChanType-11]
	_ = x[CommClause-12]
	_ = x[Comment-13]
	_ = x[CommentGroup-14]
	_ = x[CompositeLit-15]
	_ = x[DeclStmt-16]
	_ = x[DeferStmt-17]
	_ = x[Ellipsis-18]
	_ = x[EmptyStmt-19]
	_ = x[ExprStmt-20]
	_ = x[Field-21]
	_ = x[FieldList-22]
	_ = x[File-23]
	_ = x[ForStmt-24]
	_ = x[FuncDecl-25]
	_ = x[FuncLit-26]
	_ = x[FuncType-27]
	_ = x[GenDecl-28]
	_ = x[GoStmt-29]
	_ = x[Ident-30]
	_ = x[IfStmt-31]
	_ = x[ImportSpec-32]
	_ = x[IncDecStmt-33]
	_ = x[IndexExpr-34]
	_ = x[IndexListExpr-35]
	_ = x[InterfaceType-36]
	_ = x[KeyValueExpr-37]
	_ = x[LabeledStmt-38]
	_ = x[MapType-39]
	_ = x[Package-40]
	_ = x[ParenExpr-41]
	_ = x[RangeStmt-42]
	_ = x[ReturnStmt-43]
	_ = x[SelectorExpr-44]
	_ = x[SelectStmt-45]
	_ = x[SendStmt-46]
	_ = x[SliceExpr-47]
	_ = x[StarExpr-48]
	_ = x[StructType-49]
	_ = x[SwitchStmt-50]
	_ = x[TypeAssertExpr-51]
	_ = x[TypeSpec-52]
	_ = x[TypeSwitchStmt-53]
	_ = x[UnaryExpr-54]
	_ = x[ValueSpec-55]
	_ = x[Expr-56]
	_ = x[Stmt-57]
	_ = x[Decl-58]
	_ = x[Spec-59]
	_ = x[TypeExpr-60]
	_ = x[CommentGroupSlice-61]
	_ = x[CommentSlice-62]
	_ = x[DeclSlice-63]
	_ = x[ExprSlice-64]
	_ = x[FieldSlice-65]
	_ = x[IdentSlice-66]
	_ = x[ImportSpecSlice-67]
	_ = x[SpecSlice-68]
	_ = x[StmtSlice-69]
}

const _NodeType_name = "ArrayTypeAssignStmtBadDeclBadExprBadStmtBasicLitBinaryExprBlockStmtBranchStmtCallExprCaseClauseChanTypeCommClauseCommentCommentGroupCompositeLitDeclStmtDeferStmtEllipsisEmptyStmtExprStmtFieldFieldListFileForStmtFuncDeclFuncLitFuncTypeGenDeclGoStmtIdentIfStmtImportSpecIncDecStmtIndexExprIndexListExprInterfaceTypeKeyValueExprLabeledStmtMapTypePackageParenExprRangeStmtReturnStmtSelectorExprSelectStmtSendStmtSliceExprStarExprStructTypeSwitchStmtTypeAssertExprTypeSpecTypeSwitchStmtUnaryExprValueSpecExprStmtDeclSpecTypeExprCommentGroupSliceCommentSliceDeclSliceExprSliceFieldSliceIdentSliceImportSpecSliceSpecSliceStmtSlice"

var _NodeType_index = [...]uint16{0, 9, 19, 26, 33, 40, 48, 58, 67, 77, 85, 95, 103, 113, 120, 132, 144, 152, 161, 169, 178, 186, 191, 200, 204, 211, 219, 226, 234, 241, 247, 252, 258, 268, 278, 287, 300, 313, 325, 336, 343, 350, 359, 368, 378, 390, 400, 408, 417, 425, 435, 445, 459, 467, 481, 490, 499, 503, 507, 511, 515, 523, 540, 552, 561, 570, 580, 590, 605, 614, 623}

func (i NodeType) String() string {
	if i < 0 || i >= NodeType(len(_NodeType_index)-1) {
		return "NodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeType_name[_NodeType_index[i]:_NodeType_index[i+1]]
}
