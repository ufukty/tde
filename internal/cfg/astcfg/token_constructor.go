package astcfg

import "go/token"

type TokenConstructor struct {
	Literals             []token.Token
	Arithmetic           []token.Token
	ArithmeticAssignment []token.Token
	Bitwise              []token.Token
	Logical              []token.Token
	LogicalAssignment    []token.Token
	Comparison           []token.Token
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
		Logical: []token.Token{
			token.LAND,  // &&
			token.LOR,   // ||
			token.ARROW, // <-
			token.INC,   // ++
			token.DEC,   // --
		},
		LogicalAssignment: []token.Token{
			token.AND_ASSIGN,     // &=
			token.OR_ASSIGN,      // |=
			token.XOR_ASSIGN,     // ^=
			token.SHL_ASSIGN,     // <<=
			token.SHR_ASSIGN,     // >>=
			token.AND_NOT_ASSIGN, // &^=
		},
		Comparison: []token.Token{
			token.EQL, // ==
			token.LSS, // <
			token.GTR, // >
			token.NEQ, // !=
			token.LEQ, // <=
			token.GEQ, // >=
		},
	}
}

var tokenConstructor = NewNodeConstructor()

const (
	ILLEGAL token.Token = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)

	operator_beg
	// Operators and delimiters

	ASSIGN // =
	NOT    // !

	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	operator_end

	keyword_beg
	// Keywords
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR
	keyword_end

	additional_beg
	// additional tokens, handled in an ad-hoc manner
	TILDE
	additional_end
)
