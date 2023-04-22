package cp_printer

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
	"reflect"

	"golang.org/x/exp/slices"
)

var tokenMap = map[string]string{
	"ILLEGAL":     "ILLEGAL",
	"EOF":         "EOF",
	"COMMENT":     "COMMENT",
	"IDENT":       "IDENT",
	"INT":         "INT",
	"FLOAT":       "FLOAT",
	"IMAG":        "IMAG",
	"CHAR":        "CHAR",
	"STRING":      "STRING",
	"+":           "ADD",
	"-":           "SUB",
	"*":           "MUL",
	"/":           "QUO",
	"%":           "REM",
	"&":           "AND",
	"|":           "OR",
	"^":           "XOR",
	"<<":          "SHL",
	">>":          "SHR",
	"&^":          "AND_NOT",
	"+=":          "ADD_ASSIGN",
	"-=":          "SUB_ASSIGN",
	"*=":          "MUL_ASSIGN",
	"/=":          "QUO_ASSIGN",
	"%=":          "REM_ASSIGN",
	"&=":          "AND_ASSIGN",
	"|=":          "OR_ASSIGN",
	"^=":          "XOR_ASSIGN",
	"<<=":         "SHL_ASSIGN",
	">>=":         "SHR_ASSIGN",
	"&^=":         "AND_NOT_ASSIGN",
	"&&":          "LAND",
	"||":          "LOR",
	"<-":          "ARROW",
	"++":          "INC",
	"--":          "DEC",
	"==":          "EQL",
	"<":           "LSS",
	">":           "GTR",
	"=":           "ASSIGN",
	"!":           "NOT",
	"!=":          "NEQ",
	"<=":          "LEQ",
	">=":          "GEQ",
	":=":          "DEFINE",
	"...":         "ELLIPSIS",
	"(":           "LPAREN",
	"[":           "LBRACK",
	"{":           "LBRACE",
	",":           "COMMA",
	".":           "PERIOD",
	")":           "RPAREN",
	"]":           "RBRACK",
	"}":           "RBRACE",
	";":           "SEMICOLON",
	":":           "COLON",
	"break":       "BREAK",
	"case":        "CASE",
	"chan":        "CHAN",
	"const":       "CONST",
	"continue":    "CONTINUE",
	"default":     "DEFAULT",
	"defer":       "DEFER",
	"else":        "ELSE",
	"fallthrough": "FALLTHROUGH",
	"for":         "FOR",
	"func":        "FUNC",
	"go":          "GO",
	"goto":        "GOTO",
	"if":          "IF",
	"import":      "IMPORT",
	"interface":   "INTERFACE",
	"map":         "MAP",
	"package":     "PACKAGE",
	"range":       "RANGE",
	"return":      "RETURN",
	"select":      "SELECT",
	"struct":      "STRUCT",
	"switch":      "SWITCH",
	"type":        "TYPE",
	"var":         "VAR",
	"~":           "TILDE",
}

var skipFields = []string{
	"Package", "Scope", "Objects", "Imports", "Unresolved",
	"Doc", "Obj",
	"TokPos", "EndPos", "StartPos", "NamePos", "Star",
	"Opening", "Closing", "Return", "Ellipsis",
	"Lbrace", "Rbrace", "Lparen", "Rparen",
	"ValuePos", "FileStart", "FileEnd",
}

var indentationPattern = []byte("    ")

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}

type printer struct {
	last             byte
	indent           int
	writeBuffer      []byte
	indentBeforeType bool
	target           io.Writer
}

func (p *printer) bufferIndentation() {
	for j := 0; j < p.indent; j++ {
		p.writeBuffer = append(p.writeBuffer, indentationPattern...)
	}
	p.indentBeforeType = false
}

func (p *printer) write(data []byte) {
	var bufferedSoFar = 0
	for i, b := range data {
		if b == '\n' {
			// add the passed line into buffer
			p.writeBuffer = append(p.writeBuffer, data[bufferedSoFar:i+1]...)
			bufferedSoFar += (i + 1) - bufferedSoFar

			// don't print the indentation, but leave a mark to print before next write
			p.indentBeforeType = true
		} else {
			if p.indentBeforeType {
				p.bufferIndentation()
			}
		}
	}
	if len(data) > bufferedSoFar {
		p.writeBuffer = append(p.writeBuffer, data[bufferedSoFar:]...)
	}
	p.target.Write(p.writeBuffer)
	p.writeBuffer = []byte{}
}

func (p *printer) lineEnd() {
	p.write([]byte{'\n'})
}

func (p *printer) printf(format string, value ...any) {
	p.write([]byte(fmt.Sprintf(format, value...)))
}

func (p *printer) recurse(value reflect.Value) {
	if isNil(value) {
		p.printf("nil")
		return
	}

	kind := value.Kind()

	// open curly braces
	switch kind {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		p.printf("%s{", value.Type())
		p.indent++
	}

	switch kind {

	case reflect.Interface:
		p.recurse(value.Elem())

	case reflect.Map:
		if value.Len() > 0 {
			p.lineEnd()
			for _, key := range value.MapKeys() {
				p.recurse(key)
				p.printf(": ")
				p.recurse(value.MapIndex(key))
				p.printf(",")
				p.lineEnd()
			}
		}

	case reflect.Pointer:
		p.printf("&")
		p.recurse(value.Elem())

	case reflect.Array:
		if value.Len() > 0 {
			p.lineEnd()
			for i, n := 0, value.Len(); i < n; i++ {
				p.recurse(value.Index(i))
				p.printf(",")
				p.lineEnd()
			}
		}

	case reflect.Slice:
		if s, ok := value.Interface().([]byte); ok {
			p.printf("%#q", s)
			return
		}
		if value.Len() > 0 {
			p.lineEnd()
			for i, n := 0, value.Len(); i < n; i++ {
				p.recurse(value.Index(i))
				p.printf(",")
				p.lineEnd()
			}
		}

	case reflect.Struct:
		typ := value.Type()
		first := true
		for i, n := 0, typ.NumField(); i < n; i++ {
			if name := typ.Field(i).Name; token.IsExported(name) {
				if slices.Index(skipFields, name) == -1 {
					value := value.Field(i)
					if first {
						p.lineEnd()
						first = false
					}
					p.printf("%s: ", name)
					p.recurse(value)
					p.printf(",")
					p.lineEnd()
				}
			}
		}

	default:
		v := value.Interface()
		switch v := v.(type) {
		case string:
			p.printf("%q", v)
		case token.Token:
			p.printf("%s", "token."+tokenMap[v.String()])

		default:
			p.printf("%v", v)
		}
	}

	// close curly braces
	switch kind {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		p.indent--
		p.printf("}")
	}
}

func (p *printer) Print(x any) {
	p.recurse(reflect.ValueOf(x))
}

func Println(node ast.Node) {
	Fprintln(os.Stdout, node)
}

func Fprintln(writer io.Writer, node ast.Node) {
	p := printer{
		last:             '\n',
		indent:           0,
		writeBuffer:      []byte{},
		indentBeforeType: false,
		target:           writer,
	}
	p.Print(node)
	p.printf("\n")
}

func Print(node ast.Node) {
	Fprint(os.Stdout, node)
}

func Fprint(writer io.Writer, node ast.Node) {
	p := printer{
		last:             '\n',
		indent:           0,
		writeBuffer:      []byte{},
		indentBeforeType: false,
		target:           writer,
	}
	p.Print(node)
}
