#

**TypeNames/Basic**

| String            | Name         | Type         | IsAlias           | Name         | Kind         | Info        |
| ----------------- | ------------ | ------------ | ----------------- | ------------ | ------------ | ----------- |
| `type bool`       | `bool`       | `bool`       | `%!s(bool=false)` | `bool`       | `Bool`       | `IsBoolean` |
| `type byte`       | `byte`       | `byte`       | `%!s(bool=true)`  | `byte`       | `Uint8`      | ``          |
| `type complex128` | `complex128` | `complex128` | `%!s(bool=false)` | `complex128` | `Complex128` | `IsComplex` |
| `type complex64`  | `complex64`  | `complex64`  | `%!s(bool=false)` | `complex64`  | `Complex64`  | `IsComplex` |
| `type float32`    | `float32`    | `float32`    | `%!s(bool=false)` | `float32`    | `Float32`    | `IsFloat`   |
| `type float64`    | `float64`    | `float64`    | `%!s(bool=false)` | `float64`    | `Float64`    | `IsFloat`   |
| `type int`        | `int`        | `int`        | `%!s(bool=false)` | `int`        | `Int`        | `IsInteger` |
| `type int16`      | `int16`      | `int16`      | `%!s(bool=false)` | `int16`      | `Int16`      | `IsInteger` |
| `type int32`      | `int32`      | `int32`      | `%!s(bool=false)` | `int32`      | `Int32`      | `IsInteger` |
| `type int64`      | `int64`      | `int64`      | `%!s(bool=false)` | `int64`      | `Int64`      | `IsInteger` |
| `type int8`       | `int8`       | `int8`       | `%!s(bool=false)` | `int8`       | `Int8`       | `IsInteger` |
| `type rune`       | `rune`       | `rune`       | `%!s(bool=true)`  | `rune`       | `Int32`      | `IsInteger` |
| `type string`     | `string`     | `string`     | `%!s(bool=false)` | `string`     | `String`     | `IsString`  |
| `type uint`       | `uint`       | `uint`       | `%!s(bool=false)` | `uint`       | `Uint`       | ``          |
| `type uint16`     | `uint16`     | `uint16`     | `%!s(bool=false)` | `uint16`     | `Uint16`     | ``          |
| `type uint32`     | `uint32`     | `uint32`     | `%!s(bool=false)` | `uint32`     | `Uint32`     | ``          |
| `type uint64`     | `uint64`     | `uint64`     | `%!s(bool=false)` | `uint64`     | `Uint64`     | ``          |
| `type uint8`      | `uint8`      | `uint8`      | `%!s(bool=false)` | `uint8`      | `Uint8`      | ``          |
| `type uintptr`    | `uintptr`    | `uintptr`    | `%!s(bool=false)` | `uintptr`    | `Uintptr`    | ``          |

**TypeNames/Interface**

| String                   | Name  | Type  | IsAlias          | IsComparable      | NumEmbeddeds | NumMethods   |
| ------------------------ | ----- | ----- | ---------------- | ----------------- | ------------ | ------------ |
| `type any = interface{}` | `any` | `any` | `%!s(bool=true)` | `%!s(bool=false)` | `%!s(int=0)` | `%!s(int=0)` |

**TypeNames/Named**

| String                                  | Name         | Type         | IsAlias           | NumMethods   | TypeArgs                     | TypeParams                        |
| --------------------------------------- | ------------ | ------------ | ----------------- | ------------ | ---------------------------- | --------------------------------- |
| `type comparable interface{comparable}` | `comparable` | `comparable` | `%!s(bool=false)` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |
| `type error interface{Error() string}`  | `error`      | `error`      | `%!s(bool=false)` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

**Consts/Basic**

| String                     | Name    | Type           | Val     | Name           | Kind          | Info |
| -------------------------- | ------- | -------------- | ------- | -------------- | ------------- | ---- |
| `const false untyped bool` | `false` | `untyped bool` | `false` | `untyped bool` | `UntypedBool` | ``   |
| `const iota untyped int`   | `iota`  | `untyped int`  | `0`     | `untyped int`  | `UntypedInt`  | ``   |
| `const true untyped bool`  | `true`  | `untyped bool` | `true`  | `untyped bool` | `UntypedBool` | ``   |

**Builtins/Basic**

| String            | Name      | Type           | Name           | Kind      | Info |
| ----------------- | --------- | -------------- | -------------- | --------- | ---- |
| `builtin append`  | `append`  | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin cap`     | `cap`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin clear`   | `clear`   | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin close`   | `close`   | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin complex` | `complex` | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin copy`    | `copy`    | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin delete`  | `delete`  | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin imag`    | `imag`    | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin len`     | `len`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin make`    | `make`    | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin max`     | `max`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin min`     | `min`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin new`     | `new`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin panic`   | `panic`   | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin print`   | `print`   | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin println` | `println` | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin real`    | `real`    | `invalid type` | `invalid type` | `Invalid` | ``   |
| `builtin recover` | `recover` | `invalid type` | `invalid type` | `Invalid` | ``   |

**Nils/Basic**

| String | Name  | Type          | Name          | Kind         | Info        |
| ------ | ----- | ------------- | ------------- | ------------ | ----------- |
| `nil`  | `nil` | `untyped nil` | `untyped nil` | `UntypedNil` | `IsUntyped` |
