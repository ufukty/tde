package symbols

import (
	"fmt"
	"go/types"
	"os"
)

func ExampleScopeContent_Markdown() {
	_, _, _, _, pkg, err := prepare()
	if err != nil {
		panic(fmt.Errorf("prep: %w", err))
	}

	f, err := os.Create("output.md")
	if err != nil {
		panic(fmt.Errorf("prep, file: %w", err))
	}
	defer f.Close()

	io.Copy(f, PrintPackageAsMarkdown(pkg, 3))
	// Output:
}

func ExampleScopeContent_String() {
	fmt.Println(NewScopeContent(types.Universe).String())
	// Output:
	// Found TypeName:
	//     Found Basic:
	//         Name=bool            | String=type bool       | Type=bool            | Kind=Bool            | Info=IsBoolean       | Name=bool
	//         Name=byte            | String=type byte       | Type=byte            | Kind=Uint8           | Info=                | Name=byte
	//         Name=complex128      | String=type complex128 | Type=complex128      | Kind=Complex128      | Info=IsComplex       | Name=complex128
	//         Name=complex64       | String=type complex64  | Type=complex64       | Kind=Complex64       | Info=IsComplex       | Name=complex64
	//         Name=float32         | String=type float32    | Type=float32         | Kind=Float32         | Info=IsFloat         | Name=float32
	//         Name=float64         | String=type float64    | Type=float64         | Kind=Float64         | Info=IsFloat         | Name=float64
	//         Name=int             | String=type int        | Type=int             | Kind=Int             | Info=IsInteger       | Name=int
	//         Name=int16           | String=type int16      | Type=int16           | Kind=Int16           | Info=IsInteger       | Name=int16
	//         Name=int32           | String=type int32      | Type=int32           | Kind=Int32           | Info=IsInteger       | Name=int32
	//         Name=int64           | String=type int64      | Type=int64           | Kind=Int64           | Info=IsInteger       | Name=int64
	//         Name=int8            | String=type int8       | Type=int8            | Kind=Int8            | Info=IsInteger       | Name=int8
	//         Name=rune            | String=type rune       | Type=rune            | Kind=Int32           | Info=IsInteger       | Name=rune
	//         Name=string          | String=type string     | Type=string          | Kind=String          | Info=IsString        | Name=string
	//         Name=uint            | String=type uint       | Type=uint            | Kind=Uint            | Info=                | Name=uint
	//         Name=uint16          | String=type uint16     | Type=uint16          | Kind=Uint16          | Info=                | Name=uint16
	//         Name=uint32          | String=type uint32     | Type=uint32          | Kind=Uint32          | Info=                | Name=uint32
	//         Name=uint64          | String=type uint64     | Type=uint64          | Kind=Uint64          | Info=                | Name=uint64
	//         Name=uint8           | String=type uint8      | Type=uint8           | Kind=Uint8           | Info=                | Name=uint8
	//         Name=uintptr         | String=type uintptr    | Type=uintptr         | Kind=Uintptr         | Info=                | Name=uintptr
	//     Found Named:
	//         Name=comparable      | String=type comparable interface{comparable}              | Type=comparable
	//         Name=error           | String=type error interface{Error() string}               | Type=error
	//     Found Interface:
	//         Name=any             | String=type any = interface{}                             | Type=any
	// Found Const:
	//     Name=false           | Val=false              | String=const false untyped bool
	//     Name=iota            | Val=0                  | String=const iota untyped int
	//     Name=true            | Val=true               | String=const true untyped bool
	// Found Func:
	// Found PkgName:
	// Found Var:
	// Found Label:
	// Found Builtin:
	//     Name=append          | String=builtin append  | Type=invalid type
	//     Name=cap             | String=builtin cap     | Type=invalid type
	//     Name=clear           | String=builtin clear   | Type=invalid type
	//     Name=close           | String=builtin close   | Type=invalid type
	//     Name=complex         | String=builtin complex | Type=invalid type
	//     Name=copy            | String=builtin copy    | Type=invalid type
	//     Name=delete          | String=builtin delete  | Type=invalid type
	//     Name=imag            | String=builtin imag    | Type=invalid type
	//     Name=len             | String=builtin len     | Type=invalid type
	//     Name=make            | String=builtin make    | Type=invalid type
	//     Name=max             | String=builtin max     | Type=invalid type
	//     Name=min             | String=builtin min     | Type=invalid type
	//     Name=new             | String=builtin new     | Type=invalid type
	//     Name=panic           | String=builtin panic   | Type=invalid type
	//     Name=print           | String=builtin print   | Type=invalid type
	//     Name=println         | String=builtin println | Type=invalid type
	//     Name=real            | String=builtin real    | Type=invalid type
	//     Name=recover         | String=builtin recover | Type=invalid type
	// Found Nil:
	//     Name=nil             | String=nil
}
