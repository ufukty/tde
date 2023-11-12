package mutation

import (
	"tde/internal/utilities"

	"go/ast"
	"go/token"
)

// classified by type
var (
	Literals = []token.Token{
		token.IDENT,  // main
		token.INT,    // 12345
		token.FLOAT,  // 123.45
		token.IMAG,   // 123.45i
		token.CHAR,   // 'a'
		token.STRING, // "abc"
	}
	Arithmetic = []token.Token{
		token.ADD, // +
		token.SUB, // -
		token.MUL, // *
		token.QUO, // /
		token.REM, // %
	}
	ArithmeticAssignment = []token.Token{
		token.ADD_ASSIGN, // +=
		token.SUB_ASSIGN, // -=
		token.MUL_ASSIGN, // *=
		token.QUO_ASSIGN, // /=
		token.REM_ASSIGN, // %=
	}
	Bitwise = []token.Token{
		token.AND,     // &
		token.OR,      // |
		token.XOR,     // ^
		token.SHL,     // <<
		token.SHR,     // >>
		token.AND_NOT, // &^
	}
	BitwiseAssignment = []token.Token{
		token.AND_ASSIGN,     // &=
		token.OR_ASSIGN,      // |=
		token.XOR_ASSIGN,     // ^=
		token.SHL_ASSIGN,     // <<=
		token.SHR_ASSIGN,     // >>=
		token.AND_NOT_ASSIGN, // &^=
	}
	Logical = []token.Token{
		token.NOT,   // !
		token.LAND,  // &&
		token.LOR,   // ||
		token.ARROW, // <-
		token.INC,   // ++
		token.DEC,   // --
	}
	Comparison = []token.Token{
		token.EQL, // ==
		token.LSS, // <
		token.GTR, // >
		token.NEQ, // !=
		token.LEQ, // <=
		token.GEQ, // >=
	}
	Structure = []token.Token{
		token.ASSIGN,    // =
		token.DEFINE,    //  ==
		token.LPAREN,    // (
		token.RPAREN,    // )
		token.LBRACK,    // [
		token.RBRACK,    // ]
		token.LBRACE,    // {
		token.RBRACE,    // }
		token.ELLIPSIS,  // ...
		token.COMMA,     //
		token.PERIOD,    // .
		token.SEMICOLON, // ;
		token.COLON,     //  =
	}
	Misc = []token.Token{
		token.ILLEGAL,
		token.EOF,
		token.COMMENT,
	}
	Keywords = []token.Token{
		token.BREAK,
		token.CASE,
		token.CHAN,
		token.CONST,
		token.CONTINUE,
		token.DEFAULT,
		token.DEFER,
		token.ELSE,
		token.FALLTHROUGH,
		token.FOR,
		token.FUNC,
		token.GO,
		token.GOTO,
		token.IF,
		token.IMPORT,
		token.INTERFACE,
		token.MAP,
		token.PACKAGE,
		token.RANGE,
		token.RETURN,
		token.SELECT,
		token.STRUCT,
		token.SWITCH,
		token.TYPE,
		token.VAR,
	}
	Additional = []token.Token{
		token.TILDE,
	}
)

// classified by nodes accept as parameter value
var (
	AcceptedByBasicLit = []token.Token{
		token.INT,
		token.FLOAT,
		token.IMAG,
		token.CHAR,
		token.STRING,
	}
	AcceptedByBinaryExpr = []token.Token{
		token.NOT,  // !
		token.LAND, // &&
		token.LOR,  // ||
		token.EQL,  // ==
		token.LSS,  // <
		token.GTR,  // >
		token.NEQ,  // !=
		token.LEQ,  // <=
		token.GEQ,  // >=
	}
	AcceptedByAssignStmt = []token.Token{
		token.ASSIGN, // =
		token.DEFINE, // :=
	}
	AcceptedByBranchStmt = []token.Token{
		token.BREAK,
		token.CONTINUE,
		token.GOTO,
		token.FALLTHROUGH,
	}
	AcceptedByIncDecStmt = []token.Token{
		token.INC,
		token.DEC,
	}
	AcceptedByRangeStmt = []token.Token{
		token.ASSIGN,
		token.DEFINE,
	}
	AcceptedByUnaryExpr = []token.Token{
		token.AND,
	}
	AcceptedByGenDecl = []token.Token{
		token.IMPORT,
		token.CONST,
		token.TYPE,
		token.VAR,
	}
)

func listTokenContainingNodes(n ast.Node) (tokenContainingNodes []ast.Node) {
	ast.Inspect(n, func(n ast.Node) bool {
		switch n := n.(type) {
		case
			*ast.BasicLit,
			*ast.UnaryExpr,
			*ast.BinaryExpr,
			*ast.IncDecStmt,
			*ast.AssignStmt,
			*ast.BranchStmt,
			*ast.RangeStmt,
			*ast.GenDecl:
			tokenContainingNodes = append(tokenContainingNodes, n)
		}
		return true
	})
	return
}

func chooseNewTokenAndAssign(n ast.Node) (newToken token.Token) {
	switch n := n.(type) {
	case *ast.BasicLit:
		newToken = *utilities.PickExcept(AcceptedByBasicLit, []token.Token{n.Kind})
		n.Kind = newToken
	case *ast.UnaryExpr:
		newToken = *utilities.PickExcept(AcceptedByUnaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.BinaryExpr:
		newToken = *utilities.PickExcept(AcceptedByBinaryExpr, []token.Token{n.Op})
		n.Op = newToken
	case *ast.IncDecStmt:
		newToken = *utilities.PickExcept(AcceptedByIncDecStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.AssignStmt:
		newToken = *utilities.PickExcept(AcceptedByAssignStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.BranchStmt:
		newToken = *utilities.PickExcept(AcceptedByBranchStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.RangeStmt:
		newToken = *utilities.PickExcept(AcceptedByRangeStmt, []token.Token{n.Tok})
		n.Tok = newToken
	case *ast.GenDecl: // what are the chances for an import statement can work as a type declaration?
		newToken = *utilities.PickExcept(AcceptedByGenDecl, []token.Token{n.Tok})
		n.Tok = newToken
	}
	return
}

func Perform(n ast.Node) (changedNode ast.Node, newToken token.Token, ok bool) {
	// list available
	tokenContainingNodes := listTokenContainingNodes(n)
	if len(tokenContainingNodes) == 0 {
		return changedNode, newToken, false
	}
	changedNode = *utilities.Pick(tokenContainingNodes)
	newToken = chooseNewTokenAndAssign(changedNode)
	return changedNode, newToken, true
}
