package mutation

import (
	"fmt"
	"go/ast"
	"go/token"
	"tde/internal/utilities/pick"
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

func chooseNewTokenAndAssign(n ast.Node) (newToken token.Token, err error) {
	switch n := n.(type) {
	case *ast.BasicLit:
		newToken, err = pick.Except(AcceptedByBasicLit, []token.Token{n.Kind})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Kind = newToken
	case *ast.UnaryExpr:
		newToken, err = pick.Except(AcceptedByUnaryExpr, []token.Token{n.Op})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Op = newToken
	case *ast.BinaryExpr:
		newToken, err = pick.Except(AcceptedByBinaryExpr, []token.Token{n.Op})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Op = newToken
	case *ast.IncDecStmt:
		newToken, err = pick.Except(AcceptedByIncDecStmt, []token.Token{n.Tok})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Tok = newToken
	case *ast.AssignStmt:
		newToken, err = pick.Except(AcceptedByAssignStmt, []token.Token{n.Tok})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Tok = newToken
	case *ast.BranchStmt:
		newToken, err = pick.Except(AcceptedByBranchStmt, []token.Token{n.Tok})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Tok = newToken
	case *ast.RangeStmt:
		newToken, err = pick.Except(AcceptedByRangeStmt, []token.Token{n.Tok})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Tok = newToken
	case *ast.GenDecl: // what are the chances for an import statement can work as a type declaration?
		newToken, err = pick.Except(AcceptedByGenDecl, []token.Token{n.Tok})
		if err != nil {
			return -1, fmt.Errorf("picking the token: %w", err)
		}
		n.Tok = newToken
	}
	return
}

func Perform(n ast.Node) (changedNode ast.Node, newToken token.Token, err error) {
	// list available
	tokenContainingNodes := listTokenContainingNodes(n)
	if len(tokenContainingNodes) == 0 {
		return changedNode, newToken, fmt.Errorf("no applicable nodes")
	}
	changedNode, err = pick.Pick(tokenContainingNodes)
	if err != nil {
		return nil, -1, fmt.Errorf("picking the mutation node: %w", err)
	}
	newToken, err = chooseNewTokenAndAssign(changedNode)
	if err != nil {
		return nil, -1, fmt.Errorf("choosing the new token and assigning it: %w", err)
	}
	return changedNode, newToken, nil
}
