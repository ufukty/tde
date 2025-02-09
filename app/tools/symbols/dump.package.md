# 

**TypeNames/Named**

| String | Name | Type | IsAlias | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|
| `type main.Formatter interface{Format(f main.State, verb rune)}` |`Formatter` |`main.Formatter` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.GoStringer interface{GoString() string}` |`GoStringer` |`main.GoStringer` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.ScanState interface{Read(buf []byte) (n int, err error); ReadRune() (r rune, size int, err error); SkipSpace(); Token(skipSpace bool, f func(rune) bool) (token []byte, err error); UnreadRune() error; Width() (wid int, ok bool)}` |`ScanState` |`main.ScanState` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.Scanner interface{Scan(state main.ScanState, verb rune) error}` |`Scanner` |`main.Scanner` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.State interface{Flag(c int) bool; Precision() (prec int, ok bool); Width() (wid int, ok bool); Write(b []byte) (n int, err error)}` |`State` |`main.State` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.Stringer interface{String() string}` |`Stringer` |`main.Stringer` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.buffer []byte` |`buffer` |`main.buffer` |`%!s(bool=false)` |`%!s(int=4)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.fmt struct{buf *main.buffer; main.fmtFlags; wid int; prec int; intbuf [68]byte}` |`fmt` |`main.fmt` |`%!s(bool=false)` |`%!s(int=19)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.fmtFlags struct{widPresent bool; precPresent bool; minus bool; plus bool; sharp bool; space bool; zero bool; plusV bool; sharpV bool}` |`fmtFlags` |`main.fmtFlags` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.pp struct{buf main.buffer; arg any; value reflect.Value; fmt main.fmt; reordered bool; goodArgNum bool; panicking bool; erroring bool; wrapErrs bool; wrappedErrs []int}` |`pp` |`main.pp` |`%!s(bool=false)` |`%!s(int=26)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.readRune struct{reader io.Reader; buf [4]byte; pending int; pendBuf [4]byte; peekRune rune}` |`readRune` |`main.readRune` |`%!s(bool=false)` |`%!s(int=3)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.scanError struct{err error}` |`scanError` |`main.scanError` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.ss struct{rs io.RuneScanner; buf main.buffer; count int; atEOF bool; main.ssave}` |`ss` |`main.ss` |`%!s(bool=false)` |`%!s(int=37)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.ssave struct{validSave bool; nlIsEnd bool; nlIsSpace bool; argLimit int; limit int; maxWid int}` |`ssave` |`main.ssave` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.stringReader string` |`stringReader` |`main.stringReader` |`%!s(bool=false)` |`%!s(int=1)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.wrapError struct{msg string; err error}` |`wrapError` |`main.wrapError` |`%!s(bool=false)` |`%!s(int=2)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `type main.wrapErrors struct{msg string; errs []error}` |`wrapErrors` |`main.wrapErrors` |`%!s(bool=false)` |`%!s(int=2)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Consts/Basic**

| String | Name | Type | Val | Name | Kind | Info |
|---|---|---|---|---|---|---|
| `const main.badIndexString untyped string` |`badIndexString` |`untyped string` |`"(BADINDEX)"` |`untyped string` |`UntypedString` |`` |
| `const main.badPrecString untyped string` |`badPrecString` |`untyped string` |`"%!(BADPREC)"` |`untyped string` |`UntypedString` |`` |
| `const main.badWidthString untyped string` |`badWidthString` |`untyped string` |`"%!(BADWIDTH)"` |`untyped string` |`UntypedString` |`` |
| `const main.binaryDigits untyped string` |`binaryDigits` |`untyped string` |`"01"` |`untyped string` |`UntypedString` |`` |
| `const main.commaSpaceString untyped string` |`commaSpaceString` |`untyped string` |`", "` |`untyped string` |`UntypedString` |`` |
| `const main.decimalDigits untyped string` |`decimalDigits` |`untyped string` |`"0123456789"` |`untyped string` |`UntypedString` |`` |
| `const main.eof untyped int` |`eof` |`untyped int` |`-1` |`untyped int` |`UntypedInt` |`` |
| `const main.exponent untyped string` |`exponent` |`untyped string` |`"eEpP"` |`untyped string` |`UntypedString` |`` |
| `const main.extraString untyped string` |`extraString` |`untyped string` |`"%!(EXTRA "` |`untyped string` |`UntypedString` |`` |
| `const main.floatVerbs untyped string` |`floatVerbs` |`untyped string` |`"beEfFgGv"` |`untyped string` |`UntypedString` |`` |
| `const main.hexadecimalDigits untyped string` |`hexadecimalDigits` |`untyped string` |`"0123456789aAbBcCdDeEfF"` |`untyped string` |`UntypedString` |`` |
| `const main.hugeWid untyped int` |`hugeWid` |`untyped int` |`1073741824` |`untyped int` |`UntypedInt` |`` |
| `const main.intBits untyped int` |`intBits` |`untyped int` |`64` |`untyped int` |`UntypedInt` |`` |
| `const main.invReflectString untyped string` |`invReflectString` |`untyped string` |`"<invalid reflect.Value>"` |`untyped string` |`UntypedString` |`` |
| `const main.ldigits untyped string` |`ldigits` |`untyped string` |`"0123456789abcdefx"` |`untyped string` |`UntypedString` |`` |
| `const main.mapString untyped string` |`mapString` |`untyped string` |`"map["` |`untyped string` |`UntypedString` |`` |
| `const main.missingString untyped string` |`missingString` |`untyped string` |`"(MISSING)"` |`untyped string` |`UntypedString` |`` |
| `const main.nilAngleString untyped string` |`nilAngleString` |`untyped string` |`"<nil>"` |`untyped string` |`UntypedString` |`` |
| `const main.nilParenString untyped string` |`nilParenString` |`untyped string` |`"(nil)"` |`untyped string` |`UntypedString` |`` |
| `const main.nilString untyped string` |`nilString` |`untyped string` |`"nil"` |`untyped string` |`UntypedString` |`` |
| `const main.noVerbString untyped string` |`noVerbString` |`untyped string` |`"%!(NOVERB)"` |`untyped string` |`UntypedString` |`` |
| `const main.octalDigits untyped string` |`octalDigits` |`untyped string` |`"01234567"` |`untyped string` |`UntypedString` |`` |
| `const main.panicString untyped string` |`panicString` |`untyped string` |`"(PANIC="` |`untyped string` |`UntypedString` |`` |
| `const main.percentBangString untyped string` |`percentBangString` |`untyped string` |`"%!"` |`untyped string` |`UntypedString` |`` |
| `const main.period untyped string` |`period` |`untyped string` |`"."` |`untyped string` |`UntypedString` |`` |
| `const main.sign untyped string` |`sign` |`untyped string` |`"+-"` |`untyped string` |`UntypedString` |`` |
| `const main.signed untyped bool` |`signed` |`untyped bool` |`true` |`untyped bool` |`UntypedBool` |`` |
| `const main.udigits untyped string` |`udigits` |`untyped string` |`"0123456789ABCDEFX"` |`untyped string` |`UntypedString` |`` |
| `const main.uintptrBits untyped int` |`uintptrBits` |`untyped int` |`64` |`untyped int` |`UntypedInt` |`` |
| `const main.unsigned untyped bool` |`unsigned` |`untyped bool` |`false` |`untyped bool` |`UntypedBool` |`` |
**Funcs/Signature**

| String | Name | Type | FullName | Origin | Pkg | RecvTypeParams | Recv | Params | Results |
|---|---|---|---|---|---|---|---|---|---|
| `func main.Append(b []byte, a ...any) []byte` |`Append` |`func(b []byte, a ...any) []byte` |`main.Append` |`func main.Append(b []byte, a ...any) []byte` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(b []byte, a []any)` |`([]byte)` |
| `func main.Appendf(b []byte, format string, a ...any) []byte` |`Appendf` |`func(b []byte, format string, a ...any) []byte` |`main.Appendf` |`func main.Appendf(b []byte, format string, a ...any) []byte` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(b []byte, format string, a []any)` |`([]byte)` |
| `func main.Appendln(b []byte, a ...any) []byte` |`Appendln` |`func(b []byte, a ...any) []byte` |`main.Appendln` |`func main.Appendln(b []byte, a ...any) []byte` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(b []byte, a []any)` |`([]byte)` |
| `func main.Errorf(format string, a ...any) error` |`Errorf` |`func(format string, a ...any) error` |`main.Errorf` |`func main.Errorf(format string, a ...any) error` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(format string, a []any)` |`(error)` |
| `func main.FormatString(state main.State, verb rune) string` |`FormatString` |`func(state main.State, verb rune) string` |`main.FormatString` |`func main.FormatString(state main.State, verb rune) string` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(state main.State, verb rune)` |`(string)` |
| `func main.Fprint(w io.Writer, a ...any) (n int, err error)` |`Fprint` |`func(w io.Writer, a ...any) (n int, err error)` |`main.Fprint` |`func main.Fprint(w io.Writer, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(w io.Writer, a []any)` |`(n int, err error)` |
| `func main.Fprintf(w io.Writer, format string, a ...any) (n int, err error)` |`Fprintf` |`func(w io.Writer, format string, a ...any) (n int, err error)` |`main.Fprintf` |`func main.Fprintf(w io.Writer, format string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(w io.Writer, format string, a []any)` |`(n int, err error)` |
| `func main.Fprintln(w io.Writer, a ...any) (n int, err error)` |`Fprintln` |`func(w io.Writer, a ...any) (n int, err error)` |`main.Fprintln` |`func main.Fprintln(w io.Writer, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(w io.Writer, a []any)` |`(n int, err error)` |
| `func main.Fscan(r io.Reader, a ...any) (n int, err error)` |`Fscan` |`func(r io.Reader, a ...any) (n int, err error)` |`main.Fscan` |`func main.Fscan(r io.Reader, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r io.Reader, a []any)` |`(n int, err error)` |
| `func main.Fscanf(r io.Reader, format string, a ...any) (n int, err error)` |`Fscanf` |`func(r io.Reader, format string, a ...any) (n int, err error)` |`main.Fscanf` |`func main.Fscanf(r io.Reader, format string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r io.Reader, format string, a []any)` |`(n int, err error)` |
| `func main.Fscanln(r io.Reader, a ...any) (n int, err error)` |`Fscanln` |`func(r io.Reader, a ...any) (n int, err error)` |`main.Fscanln` |`func main.Fscanln(r io.Reader, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r io.Reader, a []any)` |`(n int, err error)` |
| `func main.Print(a ...any) (n int, err error)` |`Print` |`func(a ...any) (n int, err error)` |`main.Print` |`func main.Print(a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any)` |`(n int, err error)` |
| `func main.Printf(format string, a ...any) (n int, err error)` |`Printf` |`func(format string, a ...any) (n int, err error)` |`main.Printf` |`func main.Printf(format string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(format string, a []any)` |`(n int, err error)` |
| `func main.Println(a ...any) (n int, err error)` |`Println` |`func(a ...any) (n int, err error)` |`main.Println` |`func main.Println(a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any)` |`(n int, err error)` |
| `func main.Scan(a ...any) (n int, err error)` |`Scan` |`func(a ...any) (n int, err error)` |`main.Scan` |`func main.Scan(a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any)` |`(n int, err error)` |
| `func main.Scanf(format string, a ...any) (n int, err error)` |`Scanf` |`func(format string, a ...any) (n int, err error)` |`main.Scanf` |`func main.Scanf(format string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(format string, a []any)` |`(n int, err error)` |
| `func main.Scanln(a ...any) (n int, err error)` |`Scanln` |`func(a ...any) (n int, err error)` |`main.Scanln` |`func main.Scanln(a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any)` |`(n int, err error)` |
| `func main.Sprint(a ...any) string` |`Sprint` |`func(a ...any) string` |`main.Sprint` |`func main.Sprint(a ...any) string` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any)` |`(string)` |
| `func main.Sprintf(format string, a ...any) string` |`Sprintf` |`func(format string, a ...any) string` |`main.Sprintf` |`func main.Sprintf(format string, a ...any) string` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(format string, a []any)` |`(string)` |
| `func main.Sprintln(a ...any) string` |`Sprintln` |`func(a ...any) string` |`main.Sprintln` |`func main.Sprintln(a ...any) string` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any)` |`(string)` |
| `func main.Sscan(str string, a ...any) (n int, err error)` |`Sscan` |`func(str string, a ...any) (n int, err error)` |`main.Sscan` |`func main.Sscan(str string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(str string, a []any)` |`(n int, err error)` |
| `func main.Sscanf(str string, format string, a ...any) (n int, err error)` |`Sscanf` |`func(str string, format string, a ...any) (n int, err error)` |`main.Sscanf` |`func main.Sscanf(str string, format string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(str string, format string, a []any)` |`(n int, err error)` |
| `func main.Sscanln(str string, a ...any) (n int, err error)` |`Sscanln` |`func(str string, a ...any) (n int, err error)` |`main.Sscanln` |`func main.Sscanln(str string, a ...any) (n int, err error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(str string, a []any)` |`(n int, err error)` |
| `func main.errorHandler(errp *error)` |`errorHandler` |`func(errp *error)` |`main.errorHandler` |`func main.errorHandler(errp *error)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(errp *error)` |`()` |
| `func main.getField(v reflect.Value, i int) reflect.Value` |`getField` |`func(v reflect.Value, i int) reflect.Value` |`main.getField` |`func main.getField(v reflect.Value, i int) reflect.Value` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(v reflect.Value, i int)` |`(reflect.Value)` |
| `func main.hasX(s string) bool` |`hasX` |`func(s string) bool` |`main.hasX` |`func main.hasX(s string) bool` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(s string)` |`(bool)` |
| `func main.hexDigit(d rune) (int, bool)` |`hexDigit` |`func(d rune) (int, bool)` |`main.hexDigit` |`func main.hexDigit(d rune) (int, bool)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(d rune)` |`(int, bool)` |
| `func main.indexRune(s string, r rune) int` |`indexRune` |`func(s string, r rune) int` |`main.indexRune` |`func main.indexRune(s string, r rune) int` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(s string, r rune)` |`(int)` |
| `func main.intFromArg(a []any, argNum int) (num int, isInt bool, newArgNum int)` |`intFromArg` |`func(a []any, argNum int) (num int, isInt bool, newArgNum int)` |`main.intFromArg` |`func main.intFromArg(a []any, argNum int) (num int, isInt bool, newArgNum int)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(a []any, argNum int)` |`(num int, isInt bool, newArgNum int)` |
| `func main.isSpace(r rune) bool` |`isSpace` |`func(r rune) bool` |`main.isSpace` |`func main.isSpace(r rune) bool` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r rune)` |`(bool)` |
| `func main.newPrinter() *main.pp` |`newPrinter` |`func() *main.pp` |`main.newPrinter` |`func main.newPrinter() *main.pp` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`()` |`(*main.pp)` |
| `func main.newScanState(r io.Reader, nlIsSpace bool, nlIsEnd bool) (s *main.ss, old main.ssave)` |`newScanState` |`func(r io.Reader, nlIsSpace bool, nlIsEnd bool) (s *main.ss, old main.ssave)` |`main.newScanState` |`func main.newScanState(r io.Reader, nlIsSpace bool, nlIsEnd bool) (s *main.ss, old main.ssave)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r io.Reader, nlIsSpace bool, nlIsEnd bool)` |`(s *main.ss, old main.ssave)` |
| `func main.notSpace(r rune) bool` |`notSpace` |`func(r rune) bool` |`main.notSpace` |`func main.notSpace(r rune) bool` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r rune)` |`(bool)` |
| `func main.parseArgNumber(format string) (index int, wid int, ok bool)` |`parseArgNumber` |`func(format string) (index int, wid int, ok bool)` |`main.parseArgNumber` |`func main.parseArgNumber(format string) (index int, wid int, ok bool)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(format string)` |`(index int, wid int, ok bool)` |
| `func main.parsenum(s string, start int, end int) (num int, isnum bool, newi int)` |`parsenum` |`func(s string, start int, end int) (num int, isnum bool, newi int)` |`main.parsenum` |`func main.parsenum(s string, start int, end int) (num int, isnum bool, newi int)` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(s string, start int, end int)` |`(num int, isnum bool, newi int)` |
| `func main.tooLarge(x int) bool` |`tooLarge` |`func(x int) bool` |`main.tooLarge` |`func main.tooLarge(x int) bool` |`package fmt ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(x int)` |`(bool)` |
**Vars/Signature**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | RecvTypeParams | Recv | Params | Results |
|---|---|---|---|---|---|---|---|---|---|---|
| `var main.IsSpace func(r rune) bool` |`IsSpace` |`func(r rune) bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.IsSpace func(r rune) bool` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(r rune)` |`(bool)` |
| `var main.Parsenum func(s string, start int, end int) (num int, isnum bool, newi int)` |`Parsenum` |`func(s string, start int, end int) (num int, isnum bool, newi int)` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.Parsenum func(s string, start int, end int) (num int, isnum bool, newi int)` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(s string, start int, end int)` |`(num int, isnum bool, newi int)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var main.space [][2]uint16` |`space` |`[][2]uint16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.space [][2]uint16` |`[2]uint16` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var main.errBool error` |`errBool` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.errBool error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var main.errComplex error` |`errComplex` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.errComplex error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var main.ppFree sync.Pool` |`ppFree` |`sync.Pool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.ppFree sync.Pool` |`%!s(int=5)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var main.ssFree sync.Pool` |`ssFree` |`sync.Pool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var main.ssFree sync.Pool` |`%!s(int=5)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

## `Package`/fmt


## `Package`/fmt

**PkgNames/Basic**

| String | Name | Type | Name | Kind | Info |
|---|---|---|---|---|---|
| `package errors` |`errors` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package slices` |`slices` |`invalid type` |`invalid type` |`Invalid` |`` |

### `Package`/fmt/Errorf/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var w *main.wrapError` |`w` |`*main.wrapError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w *main.wrapError` |`main.wrapError` |

##### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var errs []error` |`errs` |`[]error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var errs []error` |`error` |

###### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var argNum int` |`argNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var argNum int` |`int` |`Int` |`IsInteger` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`


######## `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var e error` |`e` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

######### `Package`/fmt/Errorf/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/Error/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var e *main.wrapError` |`e` |`*main.wrapError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e *main.wrapError` |`main.wrapError` |

### `Package`/fmt/Unwrap/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var e *main.wrapError` |`e` |`*main.wrapError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e *main.wrapError` |`main.wrapError` |

### `Package`/fmt/Error/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var e *main.wrapErrors` |`e` |`*main.wrapErrors` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e *main.wrapErrors` |`main.wrapErrors` |

### `Package`/fmt/Unwrap/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var e *main.wrapErrors` |`e` |`*main.wrapErrors` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e *main.wrapErrors` |`main.wrapErrors` |

## `Package`/fmt


## `Package`/fmt

**PkgNames/Basic**

| String | Name | Type | Name | Kind | Info |
|---|---|---|---|---|---|
| `package strconv` |`strconv` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package utf8 ("unicode/utf8")` |`utf8` |`invalid type` |`invalid type` |`Invalid` |`` |

### `Package`/fmt/clearflags/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

### `Package`/fmt/init/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf *main.buffer` |`buf` |`*main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf *main.buffer` |`main.buffer` |
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

### `Package`/fmt/writePadding/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var buf main.buffer` |`buf` |`main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf main.buffer` |`%!s(int=4)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var padding main.buffer` |`padding` |`main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var padding main.buffer` |`%!s(int=4)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var newLen int` |`newLen` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newLen int` |`int` |`Int` |`IsInteger` |
| `var oldLen int` |`oldLen` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldLen int` |`int` |`Int` |`IsInteger` |
| `var padByte byte` |`padByte` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var padByte byte` |`byte` |`Uint8` |`` |

#### `Package`/fmt/writePadding/`BlockStmt`/`IfStmt`


##### `Package`/fmt/writePadding/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/writePadding/`BlockStmt`/`IfStmt`


##### `Package`/fmt/writePadding/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/writePadding/`BlockStmt`/`IfStmt`


##### `Package`/fmt/writePadding/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/writePadding/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/writePadding/`BlockStmt`/`RangeStmt`/`BlockStmt`


### `Package`/fmt/pad/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var width int` |`width` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var width int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/pad/`BlockStmt`/`IfStmt`


##### `Package`/fmt/pad/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/pad/`BlockStmt`/`IfStmt`


##### `Package`/fmt/pad/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/pad/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/padString/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
| `var width int` |`width` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var width int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

#### `Package`/fmt/padString/`BlockStmt`/`IfStmt`


##### `Package`/fmt/padString/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/padString/`BlockStmt`/`IfStmt`


##### `Package`/fmt/padString/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/padString/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtBoolean/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var v bool` |`v` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v bool` |`bool` |`Bool` |`IsBoolean` |

#### `Package`/fmt/fmtBoolean/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtBoolean/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtBoolean/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtUnicode/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |
| `var oldZero bool` |`oldZero` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldZero bool` |`bool` |`Bool` |`IsBoolean` |
| `var prec int` |`prec` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var prec int` |`int` |`Int` |`IsInteger` |
| `var u uint64` |`u` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var u uint64` |`uint64` |`Uint64` |`` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |

#### `Package`/fmt/fmtUnicode/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtUnicode/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var width int` |`width` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var width int` |`int` |`Int` |`IsInteger` |

###### `Package`/fmt/fmtUnicode/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtUnicode/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtUnicode/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtUnicode/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtUnicode/`BlockStmt`/`ForStmt`


##### `Package`/fmt/fmtUnicode/`BlockStmt`/`ForStmt`/`BlockStmt`


#### `Package`/fmt/fmtUnicode/`BlockStmt`/`ForStmt`


##### `Package`/fmt/fmtUnicode/`BlockStmt`/`ForStmt`/`BlockStmt`


### `Package`/fmt/fmtInteger/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var base int` |`base` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var base int` |`int` |`Int` |`IsInteger` |
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |
| `var isSigned bool` |`isSigned` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var isSigned bool` |`bool` |`Bool` |`IsBoolean` |
| `var negative bool` |`negative` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var negative bool` |`bool` |`Bool` |`IsBoolean` |
| `var oldZero bool` |`oldZero` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldZero bool` |`bool` |`Bool` |`IsBoolean` |
| `var prec int` |`prec` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var prec int` |`int` |`Int` |`IsInteger` |
| `var u uint64` |`u` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var u uint64` |`uint64` |`Uint64` |`` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |

#### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var width int` |`width` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var width int` |`int` |`Int` |`IsInteger` |

###### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var oldZero bool` |`oldZero` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldZero bool` |`bool` |`Bool` |`IsBoolean` |

##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######## `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var next uint64` |`next` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var next uint64` |`uint64` |`Uint64` |`` |

##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`ForStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`ForStmt`/`BlockStmt`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######## `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


######### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`/`IfStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`IfStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/truncateString/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |

#### `Package`/fmt/truncateString/`BlockStmt`/`IfStmt`


##### `Package`/fmt/truncateString/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |

###### `Package`/fmt/truncateString/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/truncateString/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`


######## `Package`/fmt/truncateString/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/truncateString/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/truncate/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |

#### `Package`/fmt/truncate/`BlockStmt`/`IfStmt`


##### `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |

###### `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |

######## `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/truncate/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtS/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

### `Package`/fmt/fmtBs/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

### `Package`/fmt/fmtSbx/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c byte` |`c` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c byte` |`byte` |`Uint8` |`` |
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var length int` |`length` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var length int` |`int` |`Int` |`IsInteger` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
| `var width int` |`width` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var width int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var buf main.buffer` |`buf` |`main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf main.buffer` |`%!s(int=4)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`IfStmt`


######## `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`


###### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtSbx/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtSbx/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtSx/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |

### `Package`/fmt/fmtBx/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |

### `Package`/fmt/fmtQ/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

#### `Package`/fmt/fmtQ/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtQ/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtQ/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtQ/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtQ/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtC/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c uint64` |`c` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c uint64` |`uint64` |`Uint64` |`` |
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |

#### `Package`/fmt/fmtC/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtC/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtQc/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c uint64` |`c` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c uint64` |`uint64` |`Uint64` |`` |
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |

#### `Package`/fmt/fmtQc/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtQc/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtQc/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtQc/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtQc/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/fmtFloat/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var prec int` |`prec` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var prec int` |`int` |`Int` |`IsInteger` |
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
| `var v float64` |`v` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v float64` |`float64` |`Float64` |`IsFloat` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var num []byte` |`num` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var num []byte` |`byte` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f *main.fmt` |`f` |`*main.fmt` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f *main.fmt` |`main.fmt` |

#### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var oldZero bool` |`oldZero` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldZero bool` |`bool` |`Bool` |`IsBoolean` |

###### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var tail []byte` |`tail` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var tail []byte` |`byte` |
**Vars/Array**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem | Len |
|---|---|---|---|---|---|---|---|---|
| `var tailBuf [6]byte` |`tailBuf` |`[6]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var tailBuf [6]byte` |`byte` |`%!s(int64=6)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var digits int` |`digits` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits int` |`int` |`Int` |`IsInteger` |
| `var hasDecimalPoint bool` |`hasDecimalPoint` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var hasDecimalPoint bool` |`bool` |`Bool` |`IsBoolean` |
| `var sawNonzeroDigit bool` |`sawNonzeroDigit` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var sawNonzeroDigit bool` |`bool` |`Bool` |`IsBoolean` |

###### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`


####### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######## `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


######### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


######## `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`


######### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


########## `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


########### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


########## `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


########### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


########### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`


####### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


#### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/fmtFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


## `Package`/fmt

**PkgNames/Basic**

| String | Name | Type | Name | Kind | Info |
|---|---|---|---|---|---|
| `package fmtsort ("internal/fmtsort")` |`fmtsort` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package io` |`io` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package os` |`os` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package reflect` |`reflect` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package strconv` |`strconv` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package sync` |`sync` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package utf8 ("unicode/utf8")` |`utf8` |`invalid type` |`invalid type` |`Invalid` |`` |

### `Package`/fmt/`GenDecl`/State/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/`GenDecl`/State/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |

### `Package`/fmt/`GenDecl`/State/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var prec int` |`prec` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var prec int` |`int` |`Int` |`IsInteger` |

### `Package`/fmt/`GenDecl`/State/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c int` |`c` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c int` |`int` |`Int` |`IsInteger` |

### `Package`/fmt/`GenDecl`/Formatter/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var f main.State` |`f` |`main.State` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f main.State` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/`GenDecl`/Stringer/`InterfaceType`/`FieldList`/`Field`/`FuncType`


### `Package`/fmt/`GenDecl`/GoStringer/`InterfaceType`/`FieldList`/`Field`/`FuncType`


### `Package`/fmt/parsenum/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var end int` |`end` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var end int` |`int` |`Int` |`IsInteger` |
| `var isnum bool` |`isnum` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var isnum bool` |`bool` |`Bool` |`IsBoolean` |
| `var newi int` |`newi` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newi int` |`int` |`Int` |`IsInteger` |
| `var num int` |`num` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var num int` |`int` |`Int` |`IsInteger` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
| `var start int` |`start` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var start int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/parsenum/`BlockStmt`/`IfStmt`


##### `Package`/fmt/parsenum/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/parsenum/`BlockStmt`/`ForStmt`


##### `Package`/fmt/parsenum/`BlockStmt`/`ForStmt`/`BlockStmt`


###### `Package`/fmt/parsenum/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/parsenum/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/FormatString/`FuncType`

**Vars/Array**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem | Len |
|---|---|---|---|---|---|---|---|---|
| `var tmp [16]byte` |`tmp` |`[16]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var tmp [16]byte` |`byte` |`%!s(int64=16)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var state main.State` |`state` |`main.State` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var state main.State` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |

#### `Package`/fmt/FormatString/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c rune` |`c` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c rune` |`rune` |`Int32` |`IsInteger` |

##### `Package`/fmt/FormatString/`BlockStmt`/`RangeStmt`/`BlockStmt`


###### `Package`/fmt/FormatString/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/FormatString/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/FormatString/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var w int` |`w` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/FormatString/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/FormatString/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var p int` |`p` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/FormatString/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/write/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p []byte` |`p` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p []byte` |`byte` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b *main.buffer` |`b` |`*main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b *main.buffer` |`main.buffer` |

### `Package`/fmt/writeString/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b *main.buffer` |`b` |`*main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b *main.buffer` |`main.buffer` |

### `Package`/fmt/writeByte/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c byte` |`c` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c byte` |`byte` |`Uint8` |`` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b *main.buffer` |`b` |`*main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b *main.buffer` |`main.buffer` |

### `Package`/fmt/writeRune/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b *main.buffer` |`b` |`*main.buffer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b *main.buffer` |`main.buffer` |

### `Package`/fmt/`GenDecl`/`ValueSpec`/`CompositeLit`/`KeyValueExpr`/`FuncLit`/`FuncType`


### `Package`/fmt/newPrinter/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

### `Package`/fmt/free/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/free/`BlockStmt`/`IfStmt`


##### `Package`/fmt/free/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/free/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/free/`BlockStmt`/`IfStmt`


##### `Package`/fmt/free/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/Width/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

### `Package`/fmt/Precision/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var prec int` |`prec` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var prec int` |`int` |`Int` |`IsInteger` |

### `Package`/fmt/Flag/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var b int` |`b` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/Flag/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/Flag/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/Flag/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/Flag/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/Flag/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/Flag/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/Write/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ret int` |`ret` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ret int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |

### `Package`/fmt/WriteString/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ret int` |`ret` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ret int` |`int` |`Int` |`IsInteger` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/Fprintf/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var w io.Writer` |`w` |`io.Writer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w io.Writer` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

### `Package`/fmt/Printf/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Sprintf/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

### `Package`/fmt/Appendf/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |

### `Package`/fmt/Fprint/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var w io.Writer` |`w` |`io.Writer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w io.Writer` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Print/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Sprint/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Append/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |

### `Package`/fmt/Fprintln/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var w io.Writer` |`w` |`io.Writer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w io.Writer` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

### `Package`/fmt/Println/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/Sprintln/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |

### `Package`/fmt/Appendln/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |

### `Package`/fmt/getField/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v reflect.Value` |`v` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var val reflect.Value` |`val` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var val reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/getField/`BlockStmt`/`IfStmt`


##### `Package`/fmt/getField/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/tooLarge/`FuncType`

**Consts/Basic**

| String | Name | Type | Val | Name | Kind | Info |
|---|---|---|---|---|---|---|
| `const max int` |`max` |`int` |`1000000` |`int` |`Int` |`IsInteger` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var x int` |`x` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var x int` |`int` |`Int` |`IsInteger` |

### `Package`/fmt/unknownType/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v reflect.Value` |`v` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/unknownType/`BlockStmt`/`IfStmt`


##### `Package`/fmt/unknownType/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/badVerb/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/badVerb/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/badVerb/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/badVerb/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/badVerb/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmtBool/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var v bool` |`v` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v bool` |`bool` |`Bool` |`IsBoolean` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

#### `Package`/fmt/fmtBool/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmt0x64/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var leading0x bool` |`leading0x` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var leading0x bool` |`bool` |`Bool` |`IsBoolean` |
| `var sharp bool` |`sharp` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var sharp bool` |`bool` |`Bool` |`IsBoolean` |
| `var v uint64` |`v` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v uint64` |`uint64` |`Uint64` |`` |

### `Package`/fmt/fmtInteger/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var isSigned bool` |`isSigned` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var isSigned bool` |`bool` |`Bool` |`IsBoolean` |
| `var v uint64` |`v` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v uint64` |`uint64` |`Uint64` |`` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtInteger/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmtFloat/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
| `var v float64` |`v` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v float64` |`float64` |`Float64` |`IsFloat` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/fmtFloat/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtFloat/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmtComplex/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
| `var v complex128` |`v` |`complex128` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v complex128` |`complex128` |`Complex128` |`IsComplex` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/fmtComplex/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtComplex/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var oldPlus bool` |`oldPlus` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldPlus bool` |`bool` |`Bool` |`IsBoolean` |

##### `Package`/fmt/fmtComplex/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmtString/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var v string` |`v` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmtBytes/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v []byte` |`v` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var typeString string` |`typeString` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var typeString string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

#### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c byte` |`c` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c byte` |`byte` |`Uint8` |`` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

######### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`


########## `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c byte` |`c` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c byte` |`byte` |`Uint8` |`` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

######### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`


########## `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtBytes/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/fmtPointer/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var u uintptr` |`u` |`uintptr` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var u uintptr` |`uintptr` |`Uintptr` |`` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var value reflect.Value` |`value` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var value reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


#### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/fmtPointer/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/catchPanic/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var method string` |`method` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var method string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/catchPanic/`BlockStmt`/`IfStmt`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var err interface{}` |`err` |`interface{}` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err interface{}` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

##### `Package`/fmt/catchPanic/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var oldFlags main.fmtFlags` |`oldFlags` |`main.fmtFlags` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var oldFlags main.fmtFlags` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

###### `Package`/fmt/catchPanic/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v reflect.Value` |`v` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

####### `Package`/fmt/catchPanic/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/catchPanic/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/catchPanic/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/handleMethods/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var handled bool` |`handled` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var handled bool` |`bool` |`Bool` |`IsBoolean` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`


##### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`


##### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

###### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var formatter main.Formatter` |`formatter` |`main.Formatter` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var formatter main.Formatter` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

##### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`


##### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var stringer main.GoStringer` |`stringer` |`main.GoStringer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var stringer main.GoStringer` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

####### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`


####### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######## `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`TypeSwitchStmt`


######### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v error` |`v` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

######### `Package`/fmt/handleMethods/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v main.Stringer` |`v` |`main.Stringer` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v main.Stringer` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/printArg/`FuncType`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/printArg/`BlockStmt`/`IfStmt`


##### `Package`/fmt/printArg/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/printArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`


####### `Package`/fmt/printArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/printArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


#### `Package`/fmt/printArg/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/printArg/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printArg/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


#### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`


##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f bool` |`f` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f bool` |`bool` |`Bool` |`IsBoolean` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f float32` |`f` |`float32` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f float32` |`float32` |`Float32` |`IsFloat` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f float64` |`f` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f float64` |`float64` |`Float64` |`IsFloat` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f complex64` |`f` |`complex64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f complex64` |`complex64` |`Complex64` |`IsComplex` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f complex128` |`f` |`complex128` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f complex128` |`complex128` |`Complex128` |`IsComplex` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f int` |`f` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f int8` |`f` |`int8` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f int8` |`int8` |`Int8` |`IsInteger` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f int16` |`f` |`int16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f int16` |`int16` |`Int16` |`IsInteger` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f int32` |`f` |`int32` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f int32` |`int32` |`Int32` |`IsInteger` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f int64` |`f` |`int64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f int64` |`int64` |`Int64` |`IsInteger` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f uint` |`f` |`uint` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f uint` |`uint` |`Uint` |`` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f uint8` |`f` |`uint8` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f uint8` |`uint8` |`Uint8` |`` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f uint16` |`f` |`uint16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f uint16` |`uint16` |`Uint16` |`` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f uint32` |`f` |`uint32` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f uint32` |`uint32` |`Uint32` |`` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f uint64` |`f` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f uint64` |`uint64` |`Uint64` |`` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f uintptr` |`f` |`uintptr` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f uintptr` |`uintptr` |`Uintptr` |`` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f string` |`f` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f string` |`string` |`String` |`IsString` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var f []byte` |`f` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f []byte` |`byte` |

##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var f reflect.Value` |`f` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

###### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var f any` |`f` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

###### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printArg/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


### `Package`/fmt/printValue/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var value reflect.Value` |`value` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var value reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var depth int` |`depth` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var depth int` |`int` |`Int` |`IsInteger` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/printValue/`BlockStmt`/`IfStmt`


##### `Package`/fmt/printValue/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/printValue/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var f reflect.Value` |`f` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`SwitchStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var sorted internal/fmtsort.SortedMap` |`sorted` |`internal/fmtsort.SortedMap` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var sorted internal/fmtsort.SortedMap` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var m internal/fmtsort.KeyValue` |`m` |`internal/fmtsort.KeyValue` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var m internal/fmtsort.KeyValue` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var name string` |`name` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var name string` |`string` |`String` |`IsString` |

########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var value reflect.Value` |`value` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var value reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var t reflect.Type` |`t` |`reflect.Type` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var t reflect.Type` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var bytes []byte` |`bytes` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var bytes []byte` |`byte` |

########## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


########## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


########## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`SwitchStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var a reflect.Value` |`a` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

######### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/printValue/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/intFromArg/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var argNum int` |`argNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var argNum int` |`int` |`Int` |`IsInteger` |
| `var isInt bool` |`isInt` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var isInt bool` |`bool` |`Bool` |`IsBoolean` |
| `var newArgNum int` |`newArgNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newArgNum int` |`int` |`Int` |`IsInteger` |
| `var num int` |`num` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var num int` |`int` |`Int` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

#### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`


##### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v reflect.Value` |`v` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

######### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int64` |`n` |`int64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int64` |`int64` |`Int64` |`IsInteger` |

########## `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


########### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n uint64` |`n` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n uint64` |`uint64` |`Uint64` |`` |

########## `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


########### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/intFromArg/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/parseArgNumber/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var index int` |`index` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var index int` |`int` |`Int` |`IsInteger` |
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/parseArgNumber/`BlockStmt`/`IfStmt`


##### `Package`/fmt/parseArgNumber/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/parseArgNumber/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/parseArgNumber/`BlockStmt`/`ForStmt`/`BlockStmt`


###### `Package`/fmt/parseArgNumber/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/parseArgNumber/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var newi int` |`newi` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newi int` |`int` |`Int` |`IsInteger` |
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var width int` |`width` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var width int` |`int` |`Int` |`IsInteger` |

######## `Package`/fmt/parseArgNumber/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/parseArgNumber/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/argNumber/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var argNum int` |`argNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var argNum int` |`int` |`Int` |`IsInteger` |
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var found bool` |`found` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var found bool` |`bool` |`Bool` |`IsBoolean` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |
| `var index int` |`index` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var index int` |`int` |`Int` |`IsInteger` |
| `var newArgNum int` |`newArgNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newArgNum int` |`int` |`Int` |`IsInteger` |
| `var newi int` |`newi` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newi int` |`int` |`Int` |`IsInteger` |
| `var numArgs int` |`numArgs` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var numArgs int` |`int` |`Int` |`IsInteger` |
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/argNumber/`BlockStmt`/`IfStmt`


##### `Package`/fmt/argNumber/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/argNumber/`BlockStmt`/`IfStmt`


##### `Package`/fmt/argNumber/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/badArgNum/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

### `Package`/fmt/missingArg/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

### `Package`/fmt/doPrintf/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var afterIndex bool` |`afterIndex` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var afterIndex bool` |`bool` |`Bool` |`IsBoolean` |
| `var argNum int` |`argNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var argNum int` |`int` |`Int` |`IsInteger` |
| `var end int` |`end` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var end int` |`int` |`Int` |`IsInteger` |
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |

#### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var lasti int` |`lasti` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var lasti int` |`int` |`Int` |`IsInteger` |
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`ForStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c byte` |`c` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c byte` |`byte` |`Uint8` |`` |

######## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


########## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


########### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/doPrintf/`BlockStmt`/`LabeledStmt`/`ForStmt`/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


#### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`


##### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |
**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

####### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/doPrintf/`BlockStmt`/`IfStmt`/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/doPrint/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var prevString bool` |`prevString` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var prevString bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

#### `Package`/fmt/doPrint/`BlockStmt`/`RangeStmt`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var argNum int` |`argNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var argNum int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/doPrint/`BlockStmt`/`RangeStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var isString bool` |`isString` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var isString bool` |`bool` |`Bool` |`IsBoolean` |

###### `Package`/fmt/doPrint/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrint/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/doPrintln/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var p *main.pp` |`p` |`*main.pp` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p *main.pp` |`main.pp` |

#### `Package`/fmt/doPrintln/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var argNum int` |`argNum` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var argNum int` |`int` |`Int` |`IsInteger` |
**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

##### `Package`/fmt/doPrintln/`BlockStmt`/`RangeStmt`/`BlockStmt`


###### `Package`/fmt/doPrintln/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doPrintln/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


## `Package`/fmt

**PkgNames/Basic**

| String | Name | Type | Name | Kind | Info |
|---|---|---|---|---|---|
| `package errors` |`errors` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package io` |`io` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package math` |`math` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package os` |`os` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package reflect` |`reflect` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package strconv` |`strconv` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package sync` |`sync` |`invalid type` |`invalid type` |`Invalid` |`` |
| `package utf8 ("unicode/utf8")` |`utf8` |`invalid type` |`invalid type` |`Invalid` |`` |

### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`


### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`


### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Signature**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | RecvTypeParams | Recv | Params | Results |
|---|---|---|---|---|---|---|---|---|---|---|
| `var f func(rune) bool` |`f` |`func(rune) bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f func(rune) bool` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(rune)` |`(bool)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var token []byte` |`token` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var token []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var skipSpace bool` |`skipSpace` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var skipSpace bool` |`bool` |`Bool` |`IsBoolean` |

#### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`/`FieldList`/`Field`/`FuncType`


### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |

### `Package`/fmt/`GenDecl`/ScanState/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/`GenDecl`/Scanner/`InterfaceType`/`FieldList`/`Field`/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var state main.ScanState` |`state` |`main.ScanState` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var state main.ScanState` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/isSpace/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
| `var rx uint16` |`rx` |`uint16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var rx uint16` |`uint16` |`Uint16` |`` |

#### `Package`/fmt/isSpace/`BlockStmt`/`IfStmt`


##### `Package`/fmt/isSpace/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/isSpace/`BlockStmt`/`RangeStmt`

**Vars/Array**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem | Len |
|---|---|---|---|---|---|---|---|---|
| `var rng [2]uint16` |`rng` |`[2]uint16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var rng [2]uint16` |`uint16` |`%!s(int64=2)` |

##### `Package`/fmt/isSpace/`BlockStmt`/`RangeStmt`/`BlockStmt`


###### `Package`/fmt/isSpace/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/isSpace/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/isSpace/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/isSpace/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/Scan/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Scanln/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Scanf/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Read/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var b []byte` |`b` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b []byte` |`byte` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var r *main.stringReader` |`r` |`*main.stringReader` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r *main.stringReader` |`main.stringReader` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/Read/`BlockStmt`/`IfStmt`


##### `Package`/fmt/Read/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/Sscan/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var str string` |`str` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var str string` |`string` |`String` |`IsString` |

### `Package`/fmt/Sscanln/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var str string` |`str` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var str string` |`string` |`String` |`IsString` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/Sscanf/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var str string` |`str` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var str string` |`string` |`String` |`IsString` |

### `Package`/fmt/Fscan/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var old main.ssave` |`old` |`main.ssave` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var old main.ssave` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var r io.Reader` |`r` |`io.Reader` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r io.Reader` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

### `Package`/fmt/Fscanln/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var old main.ssave` |`old` |`main.ssave` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var old main.ssave` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var r io.Reader` |`r` |`io.Reader` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r io.Reader` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Fscanf/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var old main.ssave` |`old` |`main.ssave` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var old main.ssave` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var r io.Reader` |`r` |`io.Reader` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r io.Reader` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

### `Package`/fmt/Read/`FuncType`

**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var buf []byte` |`buf` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var buf []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/ReadRune/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`IfStmt`


###### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/Width/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var wid int` |`wid` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var wid int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/Width/`BlockStmt`/`IfStmt`


##### `Package`/fmt/Width/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/getRune/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/getRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/getRune/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/getRune/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/getRune/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/mustReadRune/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

#### `Package`/fmt/mustReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/mustReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/UnreadRune/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

### `Package`/fmt/error/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

### `Package`/fmt/errorString/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var err string` |`err` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err string` |`string` |`String` |`IsString` |

### `Package`/fmt/Token/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Signature**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | RecvTypeParams | Recv | Params | Results |
|---|---|---|---|---|---|---|---|---|---|---|
| `var f func(rune) bool` |`f` |`func(rune) bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f func(rune) bool` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(rune)` |`(bool)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var tok []byte` |`tok` |`[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var tok []byte` |`byte` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var skipSpace bool` |`skipSpace` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var skipSpace bool` |`bool` |`Bool` |`IsBoolean` |

#### `Package`/fmt/Token/`FuncType`/`FieldList`/`Field`/`FuncType`


#### `Package`/fmt/Token/`BlockStmt`/`DeferStmt`/`CallExpr`/`FuncLit`/`FuncType`


##### `Package`/fmt/Token/`BlockStmt`/`DeferStmt`/`CallExpr`/`FuncLit`/`BlockStmt`/`IfStmt`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var e interface{}` |`e` |`interface{}` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e interface{}` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

###### `Package`/fmt/Token/`BlockStmt`/`DeferStmt`/`CallExpr`/`FuncLit`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/Token/`BlockStmt`/`DeferStmt`/`CallExpr`/`FuncLit`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var se main.scanError` |`se` |`main.scanError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var se main.scanError` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

######## `Package`/fmt/Token/`BlockStmt`/`DeferStmt`/`CallExpr`/`FuncLit`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/Token/`BlockStmt`/`DeferStmt`/`CallExpr`/`FuncLit`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/Token/`BlockStmt`/`IfStmt`


##### `Package`/fmt/Token/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/notSpace/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

### `Package`/fmt/readByte/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var b byte` |`b` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b byte` |`byte` |`Uint8` |`` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var r *main.readRune` |`r` |`*main.readRune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r *main.readRune` |`main.readRune` |

#### `Package`/fmt/readByte/`BlockStmt`/`IfStmt`


##### `Package`/fmt/readByte/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/readByte/`BlockStmt`/`IfStmt`


##### `Package`/fmt/readByte/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/ReadRune/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var rr rune` |`rr` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var rr rune` |`rune` |`Int32` |`IsInteger` |
| `var size int` |`size` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var size int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var r *main.readRune` |`r` |`*main.readRune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r *main.readRune` |`main.readRune` |

#### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/ReadRune/`BlockStmt`/`ForStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`ForStmt`/`BlockStmt`


###### `Package`/fmt/ReadRune/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/ReadRune/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/ReadRune/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/ReadRune/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/ReadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/UnreadRune/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var r *main.readRune` |`r` |`*main.readRune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r *main.readRune` |`main.readRune` |

#### `Package`/fmt/UnreadRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/UnreadRune/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/`GenDecl`/`ValueSpec`/`CompositeLit`/`KeyValueExpr`/`FuncLit`/`FuncType`


### `Package`/fmt/newScanState/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var old main.ssave` |`old` |`main.ssave` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var old main.ssave` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var r io.Reader` |`r` |`io.Reader` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r io.Reader` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var nlIsEnd bool` |`nlIsEnd` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var nlIsEnd bool` |`bool` |`Bool` |`IsBoolean` |
| `var nlIsSpace bool` |`nlIsSpace` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var nlIsSpace bool` |`bool` |`Bool` |`IsBoolean` |

#### `Package`/fmt/newScanState/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var rs io.RuneScanner` |`rs` |`io.RuneScanner` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var rs io.RuneScanner` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

##### `Package`/fmt/newScanState/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/newScanState/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/free/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var old main.ssave` |`old` |`main.ssave` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var old main.ssave` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/free/`BlockStmt`/`IfStmt`


##### `Package`/fmt/free/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/free/`BlockStmt`/`IfStmt`


##### `Package`/fmt/free/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/SkipSpace/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`


##### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

###### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/SkipSpace/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/token/`FuncType`

**Vars/Signature**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | RecvTypeParams | Recv | Params | Results |
|---|---|---|---|---|---|---|---|---|---|---|
| `var f func(rune) bool` |`f` |`func(rune) bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f func(rune) bool` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(rune)` |`(bool)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var skipSpace bool` |`skipSpace` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var skipSpace bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/token/`FuncType`/`FieldList`/`Field`/`FuncType`


#### `Package`/fmt/token/`BlockStmt`/`IfStmt`


##### `Package`/fmt/token/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/token/`BlockStmt`/`ForStmt`


##### `Package`/fmt/token/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

###### `Package`/fmt/token/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/token/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/token/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/token/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/indexRune/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |

#### `Package`/fmt/indexRune/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c rune` |`c` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c rune` |`rune` |`Int32` |`IsInteger` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/indexRune/`BlockStmt`/`RangeStmt`/`BlockStmt`


###### `Package`/fmt/indexRune/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/indexRune/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/consume/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var accept bool` |`accept` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var accept bool` |`bool` |`Bool` |`IsBoolean` |
| `var ok string` |`ok` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok string` |`string` |`String` |`IsString` |
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

#### `Package`/fmt/consume/`BlockStmt`/`IfStmt`


##### `Package`/fmt/consume/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/consume/`BlockStmt`/`IfStmt`


##### `Package`/fmt/consume/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/consume/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/consume/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/consume/`BlockStmt`/`IfStmt`


##### `Package`/fmt/consume/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/peek/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok string` |`ok` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok string` |`string` |`String` |`IsString` |
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/peek/`BlockStmt`/`IfStmt`


##### `Package`/fmt/peek/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/notEOF/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/notEOF/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

##### `Package`/fmt/notEOF/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/accept/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok string` |`ok` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok string` |`string` |`String` |`IsString` |

### `Package`/fmt/okVerb/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var okVerbs string` |`okVerbs` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var okVerbs string` |`string` |`String` |`IsString` |
| `var typ string` |`typ` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var typ string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

#### `Package`/fmt/okVerb/`BlockStmt`/`RangeStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var v rune` |`v` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v rune` |`rune` |`Int32` |`IsInteger` |

##### `Package`/fmt/okVerb/`BlockStmt`/`RangeStmt`/`BlockStmt`


###### `Package`/fmt/okVerb/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/okVerb/`BlockStmt`/`RangeStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/scanBool/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/scanBool/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanBool/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/scanBool/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


### `Package`/fmt/getBase/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var base int` |`base` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var base int` |`int` |`Int` |`IsInteger` |
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/getBase/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/getBase/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/getBase/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/getBase/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/scanNumber/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var haveDigits bool` |`haveDigits` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var haveDigits bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/scanNumber/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanNumber/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/scanNumber/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/scanNumber/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanNumber/`BlockStmt`/`ForStmt`


##### `Package`/fmt/scanNumber/`BlockStmt`/`ForStmt`/`BlockStmt`


### `Package`/fmt/scanRune/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var bitSize int` |`bitSize` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var bitSize int` |`int` |`Int` |`IsInteger` |
| `var n uint` |`n` |`uint` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n uint` |`uint` |`Uint` |`` |
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |
| `var x int64` |`x` |`int64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var x int64` |`int64` |`Int64` |`IsInteger` |

#### `Package`/fmt/scanRune/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanRune/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/scanBasePrefix/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var base int` |`base` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var base int` |`int` |`Int` |`IsInteger` |
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var zeroFound bool` |`zeroFound` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var zeroFound bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/scanBasePrefix/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanBasePrefix/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanBasePrefix/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/scanBasePrefix/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/scanBasePrefix/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/scanBasePrefix/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/scanBasePrefix/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/scanInt/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var base int` |`base` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var base int` |`int` |`Int` |`IsInteger` |
| `var bitSize int` |`bitSize` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var bitSize int` |`int` |`Int` |`IsInteger` |
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var haveDigits bool` |`haveDigits` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var haveDigits bool` |`bool` |`Bool` |`IsBoolean` |
| `var i int64` |`i` |`int64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int64` |`int64` |`Int64` |`IsInteger` |
| `var n uint` |`n` |`uint` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n uint` |`uint` |`Uint` |`` |
| `var tok string` |`tok` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var tok string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
| `var x int64` |`x` |`int64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var x int64` |`int64` |`Int64` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanInt/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/scanUint/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var base int` |`base` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var base int` |`int` |`Int` |`IsInteger` |
| `var bitSize int` |`bitSize` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var bitSize int` |`int` |`Int` |`IsInteger` |
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var haveDigits bool` |`haveDigits` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var haveDigits bool` |`bool` |`Bool` |`IsBoolean` |
| `var i uint64` |`i` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i uint64` |`uint64` |`Uint64` |`` |
| `var n uint` |`n` |`uint` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n uint` |`uint` |`Uint` |`` |
| `var tok string` |`tok` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var tok string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
| `var x uint64` |`x` |`uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var x uint64` |`uint64` |`Uint64` |`` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`IfStmt`


###### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanUint/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/floatToken/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var digits string` |`digits` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digits string` |`string` |`String` |`IsString` |
| `var exp string` |`exp` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var exp string` |`string` |`String` |`IsString` |

#### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`


##### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`


##### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`


##### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/floatToken/`BlockStmt`/`ForStmt`


##### `Package`/fmt/floatToken/`BlockStmt`/`ForStmt`/`BlockStmt`


#### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`


##### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`


####### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


#### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`


##### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`


####### `Package`/fmt/floatToken/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


### `Package`/fmt/complexTokens/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var imag string` |`imag` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var imag string` |`string` |`String` |`IsString` |
| `var imagSign string` |`imagSign` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var imagSign string` |`string` |`String` |`IsString` |
| `var parens bool` |`parens` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var parens bool` |`bool` |`Bool` |`IsBoolean` |
| `var real string` |`real` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var real string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/complexTokens/`BlockStmt`/`IfStmt`


##### `Package`/fmt/complexTokens/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/complexTokens/`BlockStmt`/`IfStmt`


##### `Package`/fmt/complexTokens/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/complexTokens/`BlockStmt`/`IfStmt`


##### `Package`/fmt/complexTokens/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/hasX/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var s string` |`s` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s string` |`string` |`String` |`IsString` |

#### `Package`/fmt/hasX/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/hasX/`BlockStmt`/`ForStmt`/`BlockStmt`


###### `Package`/fmt/hasX/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/hasX/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/convertFloat/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f float64` |`f` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f float64` |`float64` |`Float64` |`IsFloat` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var str string` |`str` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var str string` |`string` |`String` |`IsString` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

#### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var p int` |`p` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var p int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f float64` |`f` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f float64` |`float64` |`Float64` |`IsFloat` |
| `var m int` |`m` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var m int` |`int` |`Int` |`IsInteger` |

###### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var e *strconv.NumError` |`e` |`*strconv.NumError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e *strconv.NumError` |`strconv.NumError` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

######### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var e *strconv.NumError` |`e` |`*strconv.NumError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e *strconv.NumError` |`strconv.NumError` |

######### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`


##### `Package`/fmt/convertFloat/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/scanComplex/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var imag float64` |`imag` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var imag float64` |`float64` |`Float64` |`IsFloat` |
| `var n int` |`n` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var n int` |`int` |`Int` |`IsInteger` |
| `var real float64` |`real` |`float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var real float64` |`float64` |`Float64` |`IsFloat` |
| `var simag string` |`simag` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var simag string` |`string` |`String` |`IsString` |
| `var sreal string` |`sreal` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var sreal string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |

#### `Package`/fmt/scanComplex/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanComplex/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/convertString/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var str string` |`str` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var str string` |`string` |`String` |`IsString` |
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/convertString/`BlockStmt`/`IfStmt`


##### `Package`/fmt/convertString/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/convertString/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/convertString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/convertString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/convertString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/quotedString/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var quote rune` |`quote` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var quote rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


###### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`


####### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

######## `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var result string` |`result` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var result string` |`string` |`String` |`IsString` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

###### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`


####### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

######## `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`IfStmt`


########## `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/quotedString/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/hexDigit/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var d rune` |`d` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var d rune` |`rune` |`Int32` |`IsInteger` |
| `var digit int` |`digit` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var digit int` |`int` |`Int` |`IsInteger` |

#### `Package`/fmt/hexDigit/`BlockStmt`/`SwitchStmt`


##### `Package`/fmt/hexDigit/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/hexDigit/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


##### `Package`/fmt/hexDigit/`BlockStmt`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/hexByte/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var b byte` |`b` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b byte` |`byte` |`Uint8` |`` |
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |
| `var rune1 rune` |`rune1` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var rune1 rune` |`rune` |`Int32` |`IsInteger` |
| `var value1 int` |`value1` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var value1 int` |`int` |`Int` |`IsInteger` |
| `var value2 int` |`value2` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var value2 int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/hexByte/`BlockStmt`/`IfStmt`


##### `Package`/fmt/hexByte/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/hexByte/`BlockStmt`/`IfStmt`


##### `Package`/fmt/hexByte/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/hexByte/`BlockStmt`/`IfStmt`


##### `Package`/fmt/hexByte/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/hexString/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/hexString/`BlockStmt`/`ForStmt`


##### `Package`/fmt/hexString/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var b byte` |`b` |`byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var b byte` |`byte` |`Uint8` |`` |
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

###### `Package`/fmt/hexString/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/hexString/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/hexString/`BlockStmt`/`IfStmt`


##### `Package`/fmt/hexString/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/scanPercent/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/scanPercent/`BlockStmt`/`IfStmt`


##### `Package`/fmt/scanPercent/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/scanOne/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var verb rune` |`verb` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var verb rune` |`rune` |`Int32` |`IsInteger` |
**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/scanOne/`BlockStmt`/`IfStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v main.Scanner` |`v` |`main.Scanner` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v main.Scanner` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

##### `Package`/fmt/scanOne/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/scanOne/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/scanOne/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/scanOne/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/scanOne/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`


##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *bool` |`v` |`*bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *bool` |`bool` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *complex64` |`v` |`*complex64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *complex64` |`complex64` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *complex128` |`v` |`*complex128` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *complex128` |`complex128` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *int` |`v` |`*int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *int` |`int` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *int8` |`v` |`*int8` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *int8` |`int8` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *int16` |`v` |`*int16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *int16` |`int16` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *int32` |`v` |`*int32` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *int32` |`int32` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *int64` |`v` |`*int64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *int64` |`int64` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *uint` |`v` |`*uint` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *uint` |`uint` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *uint8` |`v` |`*uint8` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *uint8` |`uint8` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *uint16` |`v` |`*uint16` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *uint16` |`uint16` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *uint32` |`v` |`*uint32` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *uint32` |`uint32` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *uint64` |`v` |`*uint64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *uint64` |`uint64` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *uintptr` |`v` |`*uintptr` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *uintptr` |`uintptr` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *float32` |`v` |`*float32` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *float32` |`float32` |

###### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *float64` |`v` |`*float64` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *float64` |`float64` |

###### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *string` |`v` |`*string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *string` |`string` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var v *[]byte` |`v` |`*[]byte` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v *[]byte` |`[]byte` |

##### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var v any` |`v` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var ptr reflect.Value` |`ptr` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ptr reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
| `var val reflect.Value` |`val` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var val reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

###### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var v reflect.Value` |`v` |`reflect.Value` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var v reflect.Value` |`%!s(int=93)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |

####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var typ reflect.Type` |`typ` |`reflect.Type` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var typ reflect.Type` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var str string` |`str` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var str string` |`string` |`String` |`IsString` |

######## `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`


######### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

######### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`/`ForStmt`/`BlockStmt`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


####### `Package`/fmt/scanOne/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`SwitchStmt`/`BlockStmt`/`CaseClause`


### `Package`/fmt/errorHandler/`FuncType`

**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var errp *error` |`errp` |`*error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var errp *error` |`error` |

#### `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var e interface{}` |`e` |`interface{}` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var e interface{}` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

##### `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var se main.scanError` |`se` |`main.scanError` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var se main.scanError` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

####### `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


####### `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`IfStmt`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var eof error` |`eof` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var eof error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var ok bool` |`ok` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var ok bool` |`bool` |`Bool` |`IsBoolean` |

######## `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/errorHandler/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/doScan/`FuncType`

**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |
**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var numProcessed int` |`numProcessed` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var numProcessed int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/doScan/`BlockStmt`/`RangeStmt`

**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

##### `Package`/fmt/doScan/`BlockStmt`/`RangeStmt`/`BlockStmt`


#### `Package`/fmt/doScan/`BlockStmt`/`IfStmt`


##### `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`


####### `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var r rune` |`r` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var r rune` |`rune` |`Int32` |`IsInteger` |

######## `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doScan/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/advance/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |

#### `Package`/fmt/advance/`BlockStmt`/`ForStmt`


##### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var fmtc rune` |`fmtc` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var fmtc rune` |`rune` |`Int32` |`IsInteger` |
| `var inputc rune` |`inputc` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var inputc rune` |`rune` |`Int32` |`IsInteger` |
| `var w int` |`w` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w int` |`int` |`Int` |`IsInteger` |

###### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var newlines int` |`newlines` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var newlines int` |`int` |`Int` |`IsInteger` |
| `var trailingSpace bool` |`trailingSpace` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var trailingSpace bool` |`bool` |`Bool` |`IsBoolean` |

######## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`


######### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


########## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var j int` |`j` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var j int` |`int` |`Int` |`IsInteger` |

######### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var inputc rune` |`inputc` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var inputc rune` |`rune` |`Int32` |`IsInteger` |

########## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`ForStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


########## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var inputc rune` |`inputc` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var inputc rune` |`rune` |`Int32` |`IsInteger` |

########## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


########## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`ForStmt`/`BlockStmt`


########## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


########### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var nextc rune` |`nextc` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var nextc rune` |`rune` |`Int32` |`IsInteger` |

######## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/advance/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


### `Package`/fmt/doScanf/`FuncType`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var end int` |`end` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var end int` |`int` |`Int` |`IsInteger` |
| `var format string` |`format` |`string` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var format string` |`string` |`String` |`IsString` |
| `var numProcessed int` |`numProcessed` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var numProcessed int` |`int` |`Int` |`IsInteger` |
**Vars/Pointer**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var s *main.ss` |`s` |`*main.ss` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var s *main.ss` |`main.ss` |
**Vars/Named**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|---|---|---|
| `var err error` |`err` |`error` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var err error` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
**Vars/Slice**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Elem |
|---|---|---|---|---|---|---|---|
| `var a []any` |`a` |`[]any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var a []any` |`any` |

#### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var i int` |`i` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var i int` |`int` |`Int` |`IsInteger` |

##### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var c rune` |`c` |`rune` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var c rune` |`rune` |`Int32` |`IsInteger` |
| `var w int` |`w` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var w int` |`int` |`Int` |`IsInteger` |
| `var widPresent bool` |`widPresent` |`bool` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var widPresent bool` |`bool` |`Bool` |`IsBoolean` |
**Vars/Interface**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | IsComparable | NumEmbeddeds | NumMethods |
|---|---|---|---|---|---|---|---|---|---|
| `var arg any` |`arg` |`any` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var arg any` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(int=0)` |

###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


######## `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`


######### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`

**Vars/Basic**

| String | Name | Type | Anonymous | Embedded | IsField | Origin | Name | Kind | Info |
|---|---|---|---|---|---|---|---|---|---|
| `var f int` |`f` |`int` |`%!s(bool=false)` |`%!s(bool=false)` |`%!s(bool=false)` |`var f int` |`int` |`Int` |`IsInteger` |

####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


###### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`


####### `Package`/fmt/doScanf/`BlockStmt`/`ForStmt`/`BlockStmt`/`IfStmt`/`BlockStmt`


#### `Package`/fmt/doScanf/`BlockStmt`/`IfStmt`


##### `Package`/fmt/doScanf/`BlockStmt`/`IfStmt`/`BlockStmt`


