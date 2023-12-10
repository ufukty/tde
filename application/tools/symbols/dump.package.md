#

**TypeNames/Named**

| String                                                                                                      | Name                   | Type                        | IsAlias           | NumMethods   | TypeArgs                     | TypeParams                        |
| ----------------------------------------------------------------------------------------------------------- | ---------------------- | --------------------------- | ----------------- | ------------ | ---------------------------- | --------------------------------- |
| `type main.WalkCallbackFunction func(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int) bool` | `WalkCallbackFunction` | `main.WalkCallbackFunction` | `%!s(bool=false)` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

**Funcs/Signature**

| String                                                                                                                                | Name                     | Type                                                                                                        | FullName                      | Origin                                                                                                                                | Pkg                             | RecvTypeParams                    | Recv    | Params                                                                                                  | Results  |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------ | ----------------------------------------------------------------------------------------------------------- | ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------- | --------------------------------- | ------- | ------------------------------------------------------------------------------------------------------- | -------- |
| `func main.WalkWithNils(root go/ast.Node, callback main.WalkCallbackFunction)`                                                        | `WalkWithNils`           | `func(root go/ast.Node, callback main.WalkCallbackFunction)`                                                | `main.WalkWithNils`           | `func main.WalkWithNils(root go/ast.Node, callback main.WalkCallbackFunction)`                                                        | `package test_package ("main")` | `%!s(*types.TypeParamList=<nil>)` | `<nil>` | `(root go/ast.Node, callback main.WalkCallbackFunction)`                                                | `()`     |
| `func main.increaseLastChildIndex(childIndexTrace []int)`                                                                             | `increaseLastChildIndex` | `func(childIndexTrace []int)`                                                                               | `main.increaseLastChildIndex` | `func main.increaseLastChildIndex(childIndexTrace []int)`                                                                             | `package test_package ("main")` | `%!s(*types.TypeParamList=<nil>)` | `<nil>` | `(childIndexTrace []int)`                                                                               | `()`     |
| `func main.isNodeNil(n go/ast.Node) bool`                                                                                             | `isNodeNil`              | `func(n go/ast.Node) bool`                                                                                  | `main.isNodeNil`              | `func main.isNodeNil(n go/ast.Node) bool`                                                                                             | `package test_package ("main")` | `%!s(*types.TypeParamList=<nil>)` | `<nil>` | `(n go/ast.Node)`                                                                                       | `(bool)` |
| `func main.walkAstTypeFieldsIfSet(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)` | `walkAstTypeFieldsIfSet` | `func(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)`   | `main.walkAstTypeFieldsIfSet` | `func main.walkAstTypeFieldsIfSet(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)` | `package test_package ("main")` | `%!s(*types.TypeParamList=<nil>)` | `<nil>` | `(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars []any)`    | `()`     |
| `func main.walkHelper(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)`           | `walkHelper`             | `func(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)` | `main.walkHelper`             | `func main.walkHelper(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)`           | `package test_package ("main")` | `%!s(*types.TypeParamList=<nil>)` | `<nil>` | `(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)` | `()`     |

## `Package`/test_package

**PkgNames/Basic**

| String                   | Name      | Type           | Name           | Kind      | Info |
| ------------------------ | --------- | -------------- | -------------- | --------- | ---- |
| `package ast ("go/ast")` | `ast`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `package fmt`            | `fmt`     | `invalid type` | `invalid type` | `Invalid` | ``   |
| `package reflect`        | `reflect` | `invalid type` | `invalid type` | `Invalid` | ``   |

### `Package`/test_package/`GenDecl`/WalkCallbackFunction/`FuncType`

**Vars/Slice**

| String                          | Name              | Type            | Anonymous         | Embedded          | IsField           | Origin                          | Elem          |
| ------------------------------- | ----------------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------------- | ------------- |
| `var childIndexTrace []int`     | `childIndexTrace` | `[]int`         | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var childIndexTrace []int`     | `int`         |
| `var parentTrace []go/ast.Node` | `parentTrace`     | `[]go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var parentTrace []go/ast.Node` | `go/ast.Node` |

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Node` | `n`  | `go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Node` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

### `Package`/test_package/increaseLastChildIndex/`FuncType`

**Vars/Slice**

| String                      | Name              | Type    | Anonymous         | Embedded          | IsField           | Origin                      | Elem  |
| --------------------------- | ----------------- | ------- | ----------------- | ----------------- | ----------------- | --------------------------- | ----- |
| `var childIndexTrace []int` | `childIndexTrace` | `[]int` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var childIndexTrace []int` | `int` |

### `Package`/test_package/isNodeNil/`FuncType`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Node` | `n`  | `go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Node` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

#### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.ArrayType` | `n`  | `*go/ast.ArrayType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ArrayType` | `go/ast.ArrayType` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.AssignStmt` | `n`  | `*go/ast.AssignStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.AssignStmt` | `go/ast.AssignStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.BadDecl` | `n`  | `*go/ast.BadDecl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BadDecl` | `go/ast.BadDecl` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.BadExpr` | `n`  | `*go/ast.BadExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BadExpr` | `go/ast.BadExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.BadStmt` | `n`  | `*go/ast.BadStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BadStmt` | `go/ast.BadStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.BasicLit` | `n`  | `*go/ast.BasicLit` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BasicLit` | `go/ast.BasicLit` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.BinaryExpr` | `n`  | `*go/ast.BinaryExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BinaryExpr` | `go/ast.BinaryExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.BlockStmt` | `n`  | `*go/ast.BlockStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BlockStmt` | `go/ast.BlockStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.BranchStmt` | `n`  | `*go/ast.BranchStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BranchStmt` | `go/ast.BranchStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.CallExpr` | `n`  | `*go/ast.CallExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CallExpr` | `go/ast.CallExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.CaseClause` | `n`  | `*go/ast.CaseClause` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CaseClause` | `go/ast.CaseClause` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.ChanType` | `n`  | `*go/ast.ChanType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ChanType` | `go/ast.ChanType` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.CommClause` | `n`  | `*go/ast.CommClause` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CommClause` | `go/ast.CommClause` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.Comment` | `n`  | `*go/ast.Comment` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Comment` | `go/ast.Comment` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.CommentGroup` | `n`  | `*go/ast.CommentGroup` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CommentGroup` | `go/ast.CommentGroup` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.CompositeLit` | `n`  | `*go/ast.CompositeLit` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CompositeLit` | `go/ast.CompositeLit` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.DeclStmt` | `n`  | `*go/ast.DeclStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.DeclStmt` | `go/ast.DeclStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.DeferStmt` | `n`  | `*go/ast.DeferStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.DeferStmt` | `go/ast.DeferStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.Ellipsis` | `n`  | `*go/ast.Ellipsis` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Ellipsis` | `go/ast.Ellipsis` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.EmptyStmt` | `n`  | `*go/ast.EmptyStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.EmptyStmt` | `go/ast.EmptyStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.ExprStmt` | `n`  | `*go/ast.ExprStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ExprStmt` | `go/ast.ExprStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                | Name | Type            | Anonymous         | Embedded          | IsField           | Origin                | Elem           |
| --------------------- | ---- | --------------- | ----------------- | ----------------- | ----------------- | --------------------- | -------------- |
| `var n *go/ast.Field` | `n`  | `*go/ast.Field` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Field` | `go/ast.Field` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.FieldList` | `n`  | `*go/ast.FieldList` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FieldList` | `go/ast.FieldList` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String               | Name | Type           | Anonymous         | Embedded          | IsField           | Origin               | Elem          |
| -------------------- | ---- | -------------- | ----------------- | ----------------- | ----------------- | -------------------- | ------------- |
| `var n *go/ast.File` | `n`  | `*go/ast.File` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.File` | `go/ast.File` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.ForStmt` | `n`  | `*go/ast.ForStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ForStmt` | `go/ast.ForStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.FuncDecl` | `n`  | `*go/ast.FuncDecl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FuncDecl` | `go/ast.FuncDecl` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.FuncLit` | `n`  | `*go/ast.FuncLit` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FuncLit` | `go/ast.FuncLit` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.FuncType` | `n`  | `*go/ast.FuncType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FuncType` | `go/ast.FuncType` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.GenDecl` | `n`  | `*go/ast.GenDecl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.GenDecl` | `go/ast.GenDecl` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                 | Name | Type             | Anonymous         | Embedded          | IsField           | Origin                 | Elem            |
| ---------------------- | ---- | ---------------- | ----------------- | ----------------- | ----------------- | ---------------------- | --------------- |
| `var n *go/ast.GoStmt` | `n`  | `*go/ast.GoStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.GoStmt` | `go/ast.GoStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                | Name | Type            | Anonymous         | Embedded          | IsField           | Origin                | Elem           |
| --------------------- | ---- | --------------- | ----------------- | ----------------- | ----------------- | --------------------- | -------------- |
| `var n *go/ast.Ident` | `n`  | `*go/ast.Ident` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Ident` | `go/ast.Ident` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                 | Name | Type             | Anonymous         | Embedded          | IsField           | Origin                 | Elem            |
| ---------------------- | ---- | ---------------- | ----------------- | ----------------- | ----------------- | ---------------------- | --------------- |
| `var n *go/ast.IfStmt` | `n`  | `*go/ast.IfStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IfStmt` | `go/ast.IfStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.ImportSpec` | `n`  | `*go/ast.ImportSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ImportSpec` | `go/ast.ImportSpec` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.IncDecStmt` | `n`  | `*go/ast.IncDecStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IncDecStmt` | `go/ast.IncDecStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.IndexExpr` | `n`  | `*go/ast.IndexExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IndexExpr` | `go/ast.IndexExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                        | Name | Type                    | Anonymous         | Embedded          | IsField           | Origin                        | Elem                   |
| ----------------------------- | ---- | ----------------------- | ----------------- | ----------------- | ----------------- | ----------------------------- | ---------------------- |
| `var n *go/ast.IndexListExpr` | `n`  | `*go/ast.IndexListExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IndexListExpr` | `go/ast.IndexListExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                        | Name | Type                    | Anonymous         | Embedded          | IsField           | Origin                        | Elem                   |
| ----------------------------- | ---- | ----------------------- | ----------------- | ----------------- | ----------------- | ----------------------------- | ---------------------- |
| `var n *go/ast.InterfaceType` | `n`  | `*go/ast.InterfaceType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.InterfaceType` | `go/ast.InterfaceType` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.KeyValueExpr` | `n`  | `*go/ast.KeyValueExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.KeyValueExpr` | `go/ast.KeyValueExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                      | Name | Type                  | Anonymous         | Embedded          | IsField           | Origin                      | Elem                 |
| --------------------------- | ---- | --------------------- | ----------------- | ----------------- | ----------------- | --------------------------- | -------------------- |
| `var n *go/ast.LabeledStmt` | `n`  | `*go/ast.LabeledStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.LabeledStmt` | `go/ast.LabeledStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.MapType` | `n`  | `*go/ast.MapType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.MapType` | `go/ast.MapType` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.Package` | `n`  | `*go/ast.Package` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Package` | `go/ast.Package` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.ParenExpr` | `n`  | `*go/ast.ParenExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ParenExpr` | `go/ast.ParenExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.RangeStmt` | `n`  | `*go/ast.RangeStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.RangeStmt` | `go/ast.RangeStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.ReturnStmt` | `n`  | `*go/ast.ReturnStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ReturnStmt` | `go/ast.ReturnStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.SelectorExpr` | `n`  | `*go/ast.SelectorExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SelectorExpr` | `go/ast.SelectorExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.SelectStmt` | `n`  | `*go/ast.SelectStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SelectStmt` | `go/ast.SelectStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.SendStmt` | `n`  | `*go/ast.SendStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SendStmt` | `go/ast.SendStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.SliceExpr` | `n`  | `*go/ast.SliceExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SliceExpr` | `go/ast.SliceExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.StarExpr` | `n`  | `*go/ast.StarExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.StarExpr` | `go/ast.StarExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.StructType` | `n`  | `*go/ast.StructType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.StructType` | `go/ast.StructType` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.SwitchStmt` | `n`  | `*go/ast.SwitchStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SwitchStmt` | `go/ast.SwitchStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                         | Name | Type                     | Anonymous         | Embedded          | IsField           | Origin                         | Elem                    |
| ------------------------------ | ---- | ------------------------ | ----------------- | ----------------- | ----------------- | ------------------------------ | ----------------------- |
| `var n *go/ast.TypeAssertExpr` | `n`  | `*go/ast.TypeAssertExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.TypeAssertExpr` | `go/ast.TypeAssertExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.TypeSpec` | `n`  | `*go/ast.TypeSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.TypeSpec` | `go/ast.TypeSpec` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                         | Name | Type                     | Anonymous         | Embedded          | IsField           | Origin                         | Elem                    |
| ------------------------------ | ---- | ------------------------ | ----------------- | ----------------- | ----------------- | ------------------------------ | ----------------------- |
| `var n *go/ast.TypeSwitchStmt` | `n`  | `*go/ast.TypeSwitchStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.TypeSwitchStmt` | `go/ast.TypeSwitchStmt` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.UnaryExpr` | `n`  | `*go/ast.UnaryExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.UnaryExpr` | `go/ast.UnaryExpr` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.ValueSpec` | `n`  | `*go/ast.ValueSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ValueSpec` | `go/ast.ValueSpec` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Expr` | `n`  | `go/ast.Expr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Expr` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Stmt` | `n`  | `go/ast.Stmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Stmt` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Decl` | `n`  | `go/ast.Decl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Decl` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Spec` | `n`  | `go/ast.Spec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Spec` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

##### `Package`/test_package/isNodeNil/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Node` | `n`  | `go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Node` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

### `Package`/test_package/walkAstTypeFieldsIfSet/`FuncType`

**Vars/Named**

| String                                   | Name       | Type                        | Anonymous         | Embedded          | IsField           | Origin                                   | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------------------------- | ---------- | --------------------------- | ----------------- | ----------------- | ----------------- | ---------------------------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var callback main.WalkCallbackFunction` | `callback` | `main.WalkCallbackFunction` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var callback main.WalkCallbackFunction` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

**Vars/Slice**

| String                          | Name              | Type            | Anonymous         | Embedded          | IsField           | Origin                          | Elem          |
| ------------------------------- | ----------------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------------- | ------------- |
| `var childIndexTrace []int`     | `childIndexTrace` | `[]int`         | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var childIndexTrace []int`     | `int`         |
| `var parentTrace []go/ast.Node` | `parentTrace`     | `[]go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var parentTrace []go/ast.Node` | `go/ast.Node` |
| `var vars []any`                | `vars`            | `[]any`         | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var vars []any`                | `any`         |

#### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`

**Vars/Interface**

| String          | Name    | Type  | Anonymous         | Embedded          | IsField           | Origin          | IsComparable      | NumEmbeddeds | NumMethods   |
| --------------- | ------- | ----- | ----------------- | ----------------- | ----------------- | --------------- | ----------------- | ------------ | ------------ |
| `var field any` | `field` | `any` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field any` | `%!s(bool=false)` | `%!s(int=0)` | `%!s(int=0)` |

##### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`

###### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                        | Name    | Type                | Anonymous         | Embedded          | IsField           | Origin                        | Elem              |
| ----------------------------- | ------- | ------------------- | ----------------- | ----------------- | ----------------- | ----------------------------- | ----------------- |
| `var field []*go/ast.Comment` | `field` | `[]*go/ast.Comment` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []*go/ast.Comment` | `*go/ast.Comment` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Pointer**

| String                     | Name   | Type              | Anonymous         | Embedded          | IsField           | Origin                     | Elem             |
| -------------------------- | ------ | ----------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ---------------- |
| `var item *go/ast.Comment` | `item` | `*go/ast.Comment` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item *go/ast.Comment` | `go/ast.Comment` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                             | Name    | Type                     | Anonymous         | Embedded          | IsField           | Origin                             | Elem                   |
| ---------------------------------- | ------- | ------------------------ | ----------------- | ----------------- | ----------------- | ---------------------------------- | ---------------------- |
| `var field []*go/ast.CommentGroup` | `field` | `[]*go/ast.CommentGroup` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []*go/ast.CommentGroup` | `*go/ast.CommentGroup` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Pointer**

| String                          | Name   | Type                   | Anonymous         | Embedded          | IsField           | Origin                          | Elem                  |
| ------------------------------- | ------ | ---------------------- | ----------------- | ----------------- | ----------------- | ------------------------------- | --------------------- |
| `var item *go/ast.CommentGroup` | `item` | `*go/ast.CommentGroup` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item *go/ast.CommentGroup` | `go/ast.CommentGroup` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                           | Name    | Type                   | Anonymous         | Embedded          | IsField           | Origin                           | Elem                 |
| -------------------------------- | ------- | ---------------------- | ----------------- | ----------------- | ----------------- | -------------------------------- | -------------------- |
| `var field []*go/ast.ImportSpec` | `field` | `[]*go/ast.ImportSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []*go/ast.ImportSpec` | `*go/ast.ImportSpec` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Pointer**

| String                        | Name   | Type                 | Anonymous         | Embedded          | IsField           | Origin                        | Elem                |
| ----------------------------- | ------ | -------------------- | ----------------- | ----------------- | ----------------- | ----------------------------- | ------------------- |
| `var item *go/ast.ImportSpec` | `item` | `*go/ast.ImportSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item *go/ast.ImportSpec` | `go/ast.ImportSpec` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                      | Name    | Type              | Anonymous         | Embedded          | IsField           | Origin                      | Elem            |
| --------------------------- | ------- | ----------------- | ----------------- | ----------------- | ----------------- | --------------------------- | --------------- |
| `var field []*go/ast.Ident` | `field` | `[]*go/ast.Ident` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []*go/ast.Ident` | `*go/ast.Ident` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Pointer**

| String                   | Name   | Type            | Anonymous         | Embedded          | IsField           | Origin                   | Elem           |
| ------------------------ | ------ | --------------- | ----------------- | ----------------- | ----------------- | ------------------------ | -------------- |
| `var item *go/ast.Ident` | `item` | `*go/ast.Ident` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item *go/ast.Ident` | `go/ast.Ident` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                      | Name    | Type              | Anonymous         | Embedded          | IsField           | Origin                      | Elem            |
| --------------------------- | ------- | ----------------- | ----------------- | ----------------- | ----------------- | --------------------------- | --------------- |
| `var field []*go/ast.Field` | `field` | `[]*go/ast.Field` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []*go/ast.Field` | `*go/ast.Field` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Pointer**

| String                   | Name   | Type            | Anonymous         | Embedded          | IsField           | Origin                   | Elem           |
| ------------------------ | ------ | --------------- | ----------------- | ----------------- | ----------------- | ------------------------ | -------------- |
| `var item *go/ast.Field` | `item` | `*go/ast.Field` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item *go/ast.Field` | `go/ast.Field` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                    | Name    | Type            | Anonymous         | Embedded          | IsField           | Origin                    | Elem          |
| ------------------------- | ------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------- |
| `var field []go/ast.Stmt` | `field` | `[]go/ast.Stmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []go/ast.Stmt` | `go/ast.Stmt` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Named**

| String                 | Name   | Type          | Anonymous         | Embedded          | IsField           | Origin                 | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------- | ------ | ------------- | ----------------- | ----------------- | ----------------- | ---------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var item go/ast.Stmt` | `item` | `go/ast.Stmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item go/ast.Stmt` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                    | Name    | Type            | Anonymous         | Embedded          | IsField           | Origin                    | Elem          |
| ------------------------- | ------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------- |
| `var field []go/ast.Decl` | `field` | `[]go/ast.Decl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []go/ast.Decl` | `go/ast.Decl` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Named**

| String                 | Name   | Type          | Anonymous         | Embedded          | IsField           | Origin                 | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------- | ------ | ------------- | ----------------- | ----------------- | ----------------- | ---------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var item go/ast.Decl` | `item` | `go/ast.Decl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item go/ast.Decl` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                    | Name    | Type            | Anonymous         | Embedded          | IsField           | Origin                    | Elem          |
| ------------------------- | ------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------- |
| `var field []go/ast.Spec` | `field` | `[]go/ast.Spec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []go/ast.Spec` | `go/ast.Spec` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Named**

| String                 | Name   | Type          | Anonymous         | Embedded          | IsField           | Origin                 | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------- | ------ | ------------- | ----------------- | ----------------- | ----------------- | ---------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var item go/ast.Spec` | `item` | `go/ast.Spec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item go/ast.Spec` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                    | Name    | Type            | Anonymous         | Embedded          | IsField           | Origin                    | Elem          |
| ------------------------- | ------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------- |
| `var field []go/ast.Expr` | `field` | `[]go/ast.Expr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []go/ast.Expr` | `go/ast.Expr` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Named**

| String                 | Name   | Type          | Anonymous         | Embedded          | IsField           | Origin                 | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------- | ------ | ------------- | ----------------- | ----------------- | ----------------- | ---------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var item go/ast.Expr` | `item` | `go/ast.Expr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item go/ast.Expr` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Slice**

| String                    | Name    | Type            | Anonymous         | Embedded          | IsField           | Origin                    | Elem          |
| ------------------------- | ------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------- |
| `var field []go/ast.Node` | `field` | `[]go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field []go/ast.Node` | `go/ast.Node` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Named**

| String                 | Name   | Type          | Anonymous         | Embedded          | IsField           | Origin                 | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------- | ------ | ------------- | ----------------- | ----------------- | ----------------- | ---------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var item go/ast.Node` | `item` | `go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var item go/ast.Node` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String                  | Name    | Type          | Anonymous         | Embedded          | IsField           | Origin                  | NumMethods   | TypeArgs                     | TypeParams                        |
| ----------------------- | ------- | ------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var field go/ast.Stmt` | `field` | `go/ast.Stmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field go/ast.Stmt` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String                  | Name    | Type          | Anonymous         | Embedded          | IsField           | Origin                  | NumMethods   | TypeArgs                     | TypeParams                        |
| ----------------------- | ------- | ------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var field go/ast.Decl` | `field` | `go/ast.Decl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field go/ast.Decl` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String                  | Name    | Type          | Anonymous         | Embedded          | IsField           | Origin                  | NumMethods   | TypeArgs                     | TypeParams                        |
| ----------------------- | ------- | ------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var field go/ast.Spec` | `field` | `go/ast.Spec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field go/ast.Spec` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String                  | Name    | Type          | Anonymous         | Embedded          | IsField           | Origin                  | NumMethods   | TypeArgs                     | TypeParams                        |
| ----------------------- | ------- | ------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var field go/ast.Expr` | `field` | `go/ast.Expr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field go/ast.Expr` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String                  | Name    | Type          | Anonymous         | Embedded          | IsField           | Origin                  | NumMethods   | TypeArgs                     | TypeParams                        |
| ----------------------- | ------- | ------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var field go/ast.Node` | `field` | `go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field go/ast.Node` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

####### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Interface**

| String          | Name    | Type  | Anonymous         | Embedded          | IsField           | Origin          | IsComparable      | NumEmbeddeds | NumMethods   |
| --------------- | ------- | ----- | ----------------- | ----------------- | ----------------- | --------------- | ----------------- | ------------ | ------------ |
| `var field any` | `field` | `any` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var field any` | `%!s(bool=false)` | `%!s(int=0)` | `%!s(int=0)` |

######## `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`

**Vars/Named**

| String                 | Name | Type            | Anonymous         | Embedded          | IsField           | Origin                 | NumMethods    | TypeArgs                     | TypeParams                        |
| ---------------------- | ---- | --------------- | ----------------- | ----------------- | ----------------- | ---------------------- | ------------- | ---------------------------- | --------------------------------- |
| `var rv reflect.Value` | `rv` | `reflect.Value` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var rv reflect.Value` | `%!s(int=91)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`

######### `Package`/test_package/walkAstTypeFieldsIfSet/`BlockStmt`/`RangeStmt`/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`IfStmt`/`BlockStmt`

### `Package`/test_package/walkHelper/`FuncType`

**Vars/Named**

| String                                   | Name       | Type                        | Anonymous         | Embedded          | IsField           | Origin                                   | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------------------------- | ---------- | --------------------------- | ----------------- | ----------------- | ----------------- | ---------------------------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var callback main.WalkCallbackFunction` | `callback` | `main.WalkCallbackFunction` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var callback main.WalkCallbackFunction` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |
| `var n go/ast.Node`                      | `n`        | `go/ast.Node`               | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Node`                      | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

**Vars/Slice**

| String                          | Name              | Type            | Anonymous         | Embedded          | IsField           | Origin                          | Elem          |
| ------------------------------- | ----------------- | --------------- | ----------------- | ----------------- | ----------------- | ------------------------------- | ------------- |
| `var childIndexTrace []int`     | `childIndexTrace` | `[]int`         | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var childIndexTrace []int`     | `int`         |
| `var parentTrace []go/ast.Node` | `parentTrace`     | `[]go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var parentTrace []go/ast.Node` | `go/ast.Node` |

#### `Package`/test_package/walkHelper/`BlockStmt`/`IfStmt`

##### `Package`/test_package/walkHelper/`BlockStmt`/`IfStmt`/`BlockStmt`

##### `Package`/test_package/walkHelper/`BlockStmt`/`IfStmt`/`IfStmt`

###### `Package`/test_package/walkHelper/`BlockStmt`/`IfStmt`/`IfStmt`/`BlockStmt`

#### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.Comment` | `n`  | `*go/ast.Comment` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Comment` | `go/ast.Comment` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.CommentGroup` | `n`  | `*go/ast.CommentGroup` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CommentGroup` | `go/ast.CommentGroup` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                | Name | Type            | Anonymous         | Embedded          | IsField           | Origin                | Elem           |
| --------------------- | ---- | --------------- | ----------------- | ----------------- | ----------------- | --------------------- | -------------- |
| `var n *go/ast.Field` | `n`  | `*go/ast.Field` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Field` | `go/ast.Field` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.FieldList` | `n`  | `*go/ast.FieldList` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FieldList` | `go/ast.FieldList` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Named**

| String              | Name | Type          | Anonymous         | Embedded          | IsField           | Origin              | NumMethods   | TypeArgs                     | TypeParams                        |
| ------------------- | ---- | ------------- | ----------------- | ----------------- | ----------------- | ------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var n go/ast.Node` | `n`  | `go/ast.Node` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n go/ast.Node` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.Ellipsis` | `n`  | `*go/ast.Ellipsis` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Ellipsis` | `go/ast.Ellipsis` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.FuncLit` | `n`  | `*go/ast.FuncLit` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FuncLit` | `go/ast.FuncLit` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.CompositeLit` | `n`  | `*go/ast.CompositeLit` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CompositeLit` | `go/ast.CompositeLit` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.ParenExpr` | `n`  | `*go/ast.ParenExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ParenExpr` | `go/ast.ParenExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.SelectorExpr` | `n`  | `*go/ast.SelectorExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SelectorExpr` | `go/ast.SelectorExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.IndexExpr` | `n`  | `*go/ast.IndexExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IndexExpr` | `go/ast.IndexExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                        | Name | Type                    | Anonymous         | Embedded          | IsField           | Origin                        | Elem                   |
| ----------------------------- | ---- | ----------------------- | ----------------- | ----------------- | ----------------- | ----------------------------- | ---------------------- |
| `var n *go/ast.IndexListExpr` | `n`  | `*go/ast.IndexListExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IndexListExpr` | `go/ast.IndexListExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.SliceExpr` | `n`  | `*go/ast.SliceExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SliceExpr` | `go/ast.SliceExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                         | Name | Type                     | Anonymous         | Embedded          | IsField           | Origin                         | Elem                    |
| ------------------------------ | ---- | ------------------------ | ----------------- | ----------------- | ----------------- | ------------------------------ | ----------------------- |
| `var n *go/ast.TypeAssertExpr` | `n`  | `*go/ast.TypeAssertExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.TypeAssertExpr` | `go/ast.TypeAssertExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.CallExpr` | `n`  | `*go/ast.CallExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CallExpr` | `go/ast.CallExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.StarExpr` | `n`  | `*go/ast.StarExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.StarExpr` | `go/ast.StarExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.UnaryExpr` | `n`  | `*go/ast.UnaryExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.UnaryExpr` | `go/ast.UnaryExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.BinaryExpr` | `n`  | `*go/ast.BinaryExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BinaryExpr` | `go/ast.BinaryExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                       | Name | Type                   | Anonymous         | Embedded          | IsField           | Origin                       | Elem                  |
| ---------------------------- | ---- | ---------------------- | ----------------- | ----------------- | ----------------- | ---------------------------- | --------------------- |
| `var n *go/ast.KeyValueExpr` | `n`  | `*go/ast.KeyValueExpr` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.KeyValueExpr` | `go/ast.KeyValueExpr` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.ArrayType` | `n`  | `*go/ast.ArrayType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ArrayType` | `go/ast.ArrayType` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.StructType` | `n`  | `*go/ast.StructType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.StructType` | `go/ast.StructType` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.FuncType` | `n`  | `*go/ast.FuncType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FuncType` | `go/ast.FuncType` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                        | Name | Type                    | Anonymous         | Embedded          | IsField           | Origin                        | Elem                   |
| ----------------------------- | ---- | ----------------------- | ----------------- | ----------------- | ----------------- | ----------------------------- | ---------------------- |
| `var n *go/ast.InterfaceType` | `n`  | `*go/ast.InterfaceType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.InterfaceType` | `go/ast.InterfaceType` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.MapType` | `n`  | `*go/ast.MapType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.MapType` | `go/ast.MapType` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.ChanType` | `n`  | `*go/ast.ChanType` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ChanType` | `go/ast.ChanType` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.BadStmt` | `n`  | `*go/ast.BadStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BadStmt` | `go/ast.BadStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.DeclStmt` | `n`  | `*go/ast.DeclStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.DeclStmt` | `go/ast.DeclStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.EmptyStmt` | `n`  | `*go/ast.EmptyStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.EmptyStmt` | `go/ast.EmptyStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                      | Name | Type                  | Anonymous         | Embedded          | IsField           | Origin                      | Elem                 |
| --------------------------- | ---- | --------------------- | ----------------- | ----------------- | ----------------- | --------------------------- | -------------------- |
| `var n *go/ast.LabeledStmt` | `n`  | `*go/ast.LabeledStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.LabeledStmt` | `go/ast.LabeledStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.ExprStmt` | `n`  | `*go/ast.ExprStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ExprStmt` | `go/ast.ExprStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.SendStmt` | `n`  | `*go/ast.SendStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SendStmt` | `go/ast.SendStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.IncDecStmt` | `n`  | `*go/ast.IncDecStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IncDecStmt` | `go/ast.IncDecStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.AssignStmt` | `n`  | `*go/ast.AssignStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.AssignStmt` | `go/ast.AssignStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                 | Name | Type             | Anonymous         | Embedded          | IsField           | Origin                 | Elem            |
| ---------------------- | ---- | ---------------- | ----------------- | ----------------- | ----------------- | ---------------------- | --------------- |
| `var n *go/ast.GoStmt` | `n`  | `*go/ast.GoStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.GoStmt` | `go/ast.GoStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.DeferStmt` | `n`  | `*go/ast.DeferStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.DeferStmt` | `go/ast.DeferStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.ReturnStmt` | `n`  | `*go/ast.ReturnStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ReturnStmt` | `go/ast.ReturnStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.BranchStmt` | `n`  | `*go/ast.BranchStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BranchStmt` | `go/ast.BranchStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.BlockStmt` | `n`  | `*go/ast.BlockStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BlockStmt` | `go/ast.BlockStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                 | Name | Type             | Anonymous         | Embedded          | IsField           | Origin                 | Elem            |
| ---------------------- | ---- | ---------------- | ----------------- | ----------------- | ----------------- | ---------------------- | --------------- |
| `var n *go/ast.IfStmt` | `n`  | `*go/ast.IfStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.IfStmt` | `go/ast.IfStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.CaseClause` | `n`  | `*go/ast.CaseClause` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CaseClause` | `go/ast.CaseClause` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.SwitchStmt` | `n`  | `*go/ast.SwitchStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SwitchStmt` | `go/ast.SwitchStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                         | Name | Type                     | Anonymous         | Embedded          | IsField           | Origin                         | Elem                    |
| ------------------------------ | ---- | ------------------------ | ----------------- | ----------------- | ----------------- | ------------------------------ | ----------------------- |
| `var n *go/ast.TypeSwitchStmt` | `n`  | `*go/ast.TypeSwitchStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.TypeSwitchStmt` | `go/ast.TypeSwitchStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.CommClause` | `n`  | `*go/ast.CommClause` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.CommClause` | `go/ast.CommClause` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.SelectStmt` | `n`  | `*go/ast.SelectStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.SelectStmt` | `go/ast.SelectStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.ForStmt` | `n`  | `*go/ast.ForStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ForStmt` | `go/ast.ForStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.RangeStmt` | `n`  | `*go/ast.RangeStmt` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.RangeStmt` | `go/ast.RangeStmt` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                     | Name | Type                 | Anonymous         | Embedded          | IsField           | Origin                     | Elem                |
| -------------------------- | ---- | -------------------- | ----------------- | ----------------- | ----------------- | -------------------------- | ------------------- |
| `var n *go/ast.ImportSpec` | `n`  | `*go/ast.ImportSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ImportSpec` | `go/ast.ImportSpec` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                    | Name | Type                | Anonymous         | Embedded          | IsField           | Origin                    | Elem               |
| ------------------------- | ---- | ------------------- | ----------------- | ----------------- | ----------------- | ------------------------- | ------------------ |
| `var n *go/ast.ValueSpec` | `n`  | `*go/ast.ValueSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.ValueSpec` | `go/ast.ValueSpec` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.TypeSpec` | `n`  | `*go/ast.TypeSpec` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.TypeSpec` | `go/ast.TypeSpec` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.BadDecl` | `n`  | `*go/ast.BadDecl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.BadDecl` | `go/ast.BadDecl` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.GenDecl` | `n`  | `*go/ast.GenDecl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.GenDecl` | `go/ast.GenDecl` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                   | Name | Type               | Anonymous         | Embedded          | IsField           | Origin                   | Elem              |
| ------------------------ | ---- | ------------------ | ----------------- | ----------------- | ----------------- | ------------------------ | ----------------- |
| `var n *go/ast.FuncDecl` | `n`  | `*go/ast.FuncDecl` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.FuncDecl` | `go/ast.FuncDecl` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String               | Name | Type           | Anonymous         | Embedded          | IsField           | Origin               | Elem          |
| -------------------- | ---- | -------------- | ----------------- | ----------------- | ----------------- | -------------------- | ------------- |
| `var n *go/ast.File` | `n`  | `*go/ast.File` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.File` | `go/ast.File` |

##### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`

**Vars/Pointer**

| String                  | Name | Type              | Anonymous         | Embedded          | IsField           | Origin                  | Elem             |
| ----------------------- | ---- | ----------------- | ----------------- | ----------------- | ----------------- | ----------------------- | ---------------- |
| `var n *go/ast.Package` | `n`  | `*go/ast.Package` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var n *go/ast.Package` | `go/ast.Package` |

###### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`

**Vars/Pointer**

| String               | Name | Type           | Anonymous         | Embedded          | IsField           | Origin               | Elem          |
| -------------------- | ---- | -------------- | ----------------- | ----------------- | ----------------- | -------------------- | ------------- |
| `var f *go/ast.File` | `f`  | `*go/ast.File` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var f *go/ast.File` | `go/ast.File` |

####### `Package`/test_package/walkHelper/`BlockStmt`/`TypeSwitchStmt`/`BlockStmt`/`CaseClause`/`RangeStmt`/`BlockStmt`

### `Package`/test_package/WalkWithNils/`FuncType`

**Vars/Named**

| String                                   | Name       | Type                        | Anonymous         | Embedded          | IsField           | Origin                                   | NumMethods   | TypeArgs                     | TypeParams                        |
| ---------------------------------------- | ---------- | --------------------------- | ----------------- | ----------------- | ----------------- | ---------------------------------------- | ------------ | ---------------------------- | --------------------------------- |
| `var callback main.WalkCallbackFunction` | `callback` | `main.WalkCallbackFunction` | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var callback main.WalkCallbackFunction` | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |
| `var root go/ast.Node`                   | `root`     | `go/ast.Node`               | `%!s(bool=false)` | `%!s(bool=false)` | `%!s(bool=false)` | `var root go/ast.Node`                   | `%!s(int=0)` | `%!s(*types.TypeList=<nil>)` | `%!s(*types.TypeParamList=<nil>)` |
