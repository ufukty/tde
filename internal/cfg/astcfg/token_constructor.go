package astcfg

import "go/token"

type TokenConstructor struct {
	// classified by type
	Literals             []token.Token
	Arithmetic           []token.Token
	ArithmeticAssignment []token.Token
	Bitwise              []token.Token
	BitwiseAssignment    []token.Token
	Logical              []token.Token
	LogicalAssignment    []token.Token
	Comparison           []token.Token
	Structure            []token.Token
	Misc                 []token.Token
	Keywords             []token.Token
	Additional           []token.Token

	// classified by nodes accept as parameter value
	BasicLit              []token.Token
	AcceptedByBinaryExpr  []token.Token
	AccepetedByAssignStmt []token.Token
	AcceptedByBranchStmt  []token.Token
	AccepetedByIncDecStmt []token.Token
	AcceptedByRangeStmt   []token.Token
	AcceptedByUnaryExpr   []token.Token
}

func NewTokenConstructor() *TokenConstructor {
	return &TokenConstructor{
		Literals: []token.Token{
			token.IDENT,  // main
			token.INT,    // 12345
			token.FLOAT,  // 123.45
			token.IMAG,   // 123.45i
			token.CHAR,   // 'a'
			token.STRING, // "abc"
		},
		Arithmetic: []token.Token{
			token.ADD, // +
			token.SUB, // -
			token.MUL, // *
			token.QUO, // /
			token.REM, // %
		},
		ArithmeticAssignment: []token.Token{
			token.ADD_ASSIGN, // +=
			token.SUB_ASSIGN, // -=
			token.MUL_ASSIGN, // *=
			token.QUO_ASSIGN, // /=
			token.REM_ASSIGN, // %=
		},
		Bitwise: []token.Token{
			token.AND,     // &
			token.OR,      // |
			token.XOR,     // ^
			token.SHL,     // <<
			token.SHR,     // >>
			token.AND_NOT, // &^
		},
		BitwiseAssignment: []token.Token{
			token.AND_ASSIGN,     // &=
			token.OR_ASSIGN,      // |=
			token.XOR_ASSIGN,     // ^=
			token.SHL_ASSIGN,     // <<=
			token.SHR_ASSIGN,     // >>=
			token.AND_NOT_ASSIGN, // &^=
		},
		Logical: []token.Token{
			token.NOT,   // !
			token.LAND,  // &&
			token.LOR,   // ||
			token.ARROW, // <-
			token.INC,   // ++
			token.DEC,   // --
		},
		Comparison: []token.Token{
			token.EQL, // ==
			token.LSS, // <
			token.GTR, // >
			token.NEQ, // !=
			token.LEQ, // <=
			token.GEQ, // >=
		},
		Structure: []token.Token{
			token.ASSIGN,    // =
			token.DEFINE,    // :=
			token.LPAREN,    // (
			token.RPAREN,    // )
			token.LBRACK,    // [
			token.RBRACK,    // ]
			token.LBRACE,    // {
			token.RBRACE,    // }
			token.ELLIPSIS,  // ...
			token.COMMA,     // ,
			token.PERIOD,    // .
			token.SEMICOLON, // ;
			token.COLON,     // :
		},
		Misc: []token.Token{
			token.ILLEGAL,
			token.EOF,
			token.COMMENT,
		},
		Keywords: []token.Token{
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
		},
		Additional: []token.Token{
			token.TILDE,
		},
		BasicLit: []token.Token{
			token.INT,
			token.FLOAT,
			// token.IMAG,
			// token.CHAR,
			token.STRING,
		},
		AcceptedByBinaryExpr: []token.Token{
			token.NOT,  // !
			token.LAND, // &&
			token.LOR,  // ||
			token.EQL,  // ==
			token.LSS,  // <
			token.GTR,  // >
			token.NEQ,  // !=
			token.LEQ,  // <=
			token.GEQ,  // >=
		},
		AccepetedByAssignStmt: []token.Token{
			token.ASSIGN, // =
			token.DEFINE, // :=
		},
		AcceptedByBranchStmt: []token.Token{
			token.BREAK,
			token.CONTINUE,
			token.GOTO,
			token.FALLTHROUGH,
		},
		AccepetedByIncDecStmt: []token.Token{
			token.INC,
			token.DEC,
		},
		AcceptedByRangeStmt: []token.Token{
			token.ASSIGN,
			token.DEFINE,
		},
		AcceptedByUnaryExpr: []token.Token{
			token.AND,
		},
	}
}

var tokenConstructor = NewTokenConstructor()
