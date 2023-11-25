# Universe

**TypeName / Basic**

| String            | Name       | Type       | Kind       | Info      | Name       |
| ----------------- | ---------- | ---------- | ---------- | --------- | ---------- |
| `type bool`       | bool       | bool       | Bool       | IsBoolean | bool       |
| `type byte`       | byte       | byte       | Uint8      |           | byte       |
| `type complex128` | complex128 | complex128 | Complex128 | IsComplex | complex128 |
| `type complex64`  | complex64  | complex64  | Complex64  | IsComplex | complex64  |
| `type float32`    | float32    | float32    | Float32    | IsFloat   | float32    |
| `type float64`    | float64    | float64    | Float64    | IsFloat   | float64    |
| `type int`        | int        | int        | Int        | IsInteger | int        |
| `type int16`      | int16      | int16      | Int16      | IsInteger | int16      |
| `type int32`      | int32      | int32      | Int32      | IsInteger | int32      |
| `type int64`      | int64      | int64      | Int64      | IsInteger | int64      |
| `type int8`       | int8       | int8       | Int8       | IsInteger | int8       |
| `type rune`       | rune       | rune       | Int32      | IsInteger | rune       |
| `type string`     | string     | string     | String     | IsString  | string     |
| `type uint`       | uint       | uint       | Uint       |           | uint       |
| `type uint16`     | uint16     | uint16     | Uint16     |           | uint16     |
| `type uint32`     | uint32     | uint32     | Uint32     |           | uint32     |
| `type uint64`     | uint64     | uint64     | Uint64     |           | uint64     |
| `type uint8`      | uint8      | uint8      | Uint8      |           | uint8      |
| `type uintptr`    | uintptr    | uintptr    | Uintptr    |           | uintptr    |

**TypeName / Named**

| String                                  | Name       | Type       |
| --------------------------------------- | ---------- | ---------- |
| `type comparable interface{comparable}` | comparable | comparable |
| `type error interface{Error() string}`  | error      | error      |

**TypeName / Interface**

| String                   | Name | Type |
| ------------------------ | ---- | ---- |
| `type any = interface{}` | any  | any  |

**Const**

| String                     | Name  | Type         | Val   |
| -------------------------- | ----- | ------------ | ----- |
| `const false untyped bool` | false | untyped bool | false |
| `const iota untyped int`   | iota  | untyped int  | 0     |
| `const true untyped bool`  | true  | untyped bool | true  |

**Builtin**

| String            | Name    | Type         |
| ----------------- | ------- | ------------ |
| `builtin append`  | append  | invalid type |
| `builtin cap`     | cap     | invalid type |
| `builtin clear`   | clear   | invalid type |
| `builtin close`   | close   | invalid type |
| `builtin complex` | complex | invalid type |
| `builtin copy`    | copy    | invalid type |
| `builtin delete`  | delete  | invalid type |
| `builtin imag`    | imag    | invalid type |
| `builtin len`     | len     | invalid type |
| `builtin make`    | make    | invalid type |
| `builtin max`     | max     | invalid type |
| `builtin min`     | min     | invalid type |
| `builtin new`     | new     | invalid type |
| `builtin panic`   | panic   | invalid type |
| `builtin print`   | print   | invalid type |
| `builtin println` | println | invalid type |
| `builtin real`    | real    | invalid type |
| `builtin recover` | recover | invalid type |

**Nil**

| String | Name | Type        |
| ------ | ---- | ----------- |
| `nil`  | nil  | untyped nil |

## The Package

**TypeName / Named**

| String                                                                                                      | Name                 | Type                      |
| ----------------------------------------------------------------------------------------------------------- | -------------------- | ------------------------- |
| `type main.WalkCallbackFunction func(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int) bool` | WalkCallbackFunction | main.WalkCallbackFunction |

**Func**

| String                                                                                                                                | Name                   | FullName                    |
| ------------------------------------------------------------------------------------------------------------------------------------- | ---------------------- | --------------------------- |
| `func main.WalkWithNils(root go/ast.Node, callback main.WalkCallbackFunction)`                                                        | WalkWithNils           | main.WalkWithNils           |
| `func main.increaseLastChildIndex(childIndexTrace []int)`                                                                             | increaseLastChildIndex | main.increaseLastChildIndex |
| `func main.isNodeNil(n go/ast.Node) bool`                                                                                             | isNodeNil              | main.isNodeNil              |
| `func main.walkAstTypeFieldsIfSet(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)` | walkAstTypeFieldsIfSet | main.walkAstTypeFieldsIfSet |
| `func main.walkHelper(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)`           | walkHelper             | main.walkHelper             |

### package "main"

**PkgName**

| String                   | Name    | Type         |
| ------------------------ | ------- | ------------ |
| `package ast ("go/ast")` | ast     | invalid type |
| `package fmt`            | fmt     | invalid type |
| `package reflect`        | reflect | invalid type |

#### testdata/evolution/walk.go scope

**Var**

| String                          | Name            | Type          | Anonymous | Embedded | IsField |
| ------------------------------- | --------------- | ------------- | --------- | -------- | ------- |
| `var childIndexTrace []int`     | childIndexTrace | []int         | false     | false    | false   |
| `var n go/ast.Node`             | n               | go/ast.Node   | false     | false    | false   |
| `var parentTrace []go/ast.Node` | parentTrace     | []go/ast.Node | false     | false    | false   |

#### testdata/evolution/walk.go scope

**Var**

| String                      | Name            | Type  | Anonymous | Embedded | IsField |
| --------------------------- | --------------- | ----- | --------- | -------- | ------- |
| `var childIndexTrace []int` | childIndexTrace | []int | false     | false    | false   |

#### testdata/evolution/walk.go scope

**Var**

| String              | Name | Type        | Anonymous | Embedded | IsField |
| ------------------- | ---- | ----------- | --------- | -------- | ------- |
| `var n go/ast.Node` | n    | go/ast.Node | false     | false    | false   |

##### function scope

#### testdata/evolution/walk.go scope

**Var**

| String                                   | Name            | Type                      | Anonymous | Embedded | IsField |
| ---------------------------------------- | --------------- | ------------------------- | --------- | -------- | ------- |
| `var callback main.WalkCallbackFunction` | callback        | main.WalkCallbackFunction | false     | false    | false   |
| `var childIndexTrace []int`              | childIndexTrace | []int                     | false     | false    | false   |
| `var parentTrace []go/ast.Node`          | parentTrace     | []go/ast.Node             | false     | false    | false   |
| `var vars []any`                         | vars            | []any                     | false     | false    | false   |

##### function scope

**Var**

| String          | Name  | Type | Anonymous | Embedded | IsField |
| --------------- | ----- | ---- | --------- | -------- | ------- |
| `var field any` | field | any  | false     | false    | false   |

#### testdata/evolution/walk.go scope

**Var**

| String                                   | Name            | Type                      | Anonymous | Embedded | IsField |
| ---------------------------------------- | --------------- | ------------------------- | --------- | -------- | ------- |
| `var callback main.WalkCallbackFunction` | callback        | main.WalkCallbackFunction | false     | false    | false   |
| `var childIndexTrace []int`              | childIndexTrace | []int                     | false     | false    | false   |
| `var n go/ast.Node`                      | n               | go/ast.Node               | false     | false    | false   |
| `var parentTrace []go/ast.Node`          | parentTrace     | []go/ast.Node             | false     | false    | false   |

##### function scope

##### function scope

#### testdata/evolution/walk.go scope

**Var**

| String                                   | Name     | Type                      | Anonymous | Embedded | IsField |
| ---------------------------------------- | -------- | ------------------------- | --------- | -------- | ------- |
| `var callback main.WalkCallbackFunction` | callback | main.WalkCallbackFunction | false     | false    | false   |
| `var root go/ast.Node`                   | root     | go/ast.Node               | false     | false    | false   |
