# 

 **TypeNames/Named**

| String | Name | Type | IsAlias | NumMethods | TypeArgs | TypeParams |
|---|---|---|---|---|---|---|
| `type main.WalkCallbackFunction func(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int) bool` |`WalkCallbackFunction` |`main.WalkCallbackFunction` |`%!s(bool=false)` |`%!s(int=0)` |`%!s(*types.TypeList=<nil>)` |`%!s(*types.TypeParamList=<nil>)` |
 **Funcs/Signature**

| String | Name | Type | FullName | Origin | Pkg | RecvTypeParams | Recv | Params | Results |
|---|---|---|---|---|---|---|---|---|---|
| `func main.WalkWithNils(root go/ast.Node, callback main.WalkCallbackFunction)` |`WalkWithNils` |`func(root go/ast.Node, callback main.WalkCallbackFunction)` |`main.WalkWithNils` |`func main.WalkWithNils(root go/ast.Node, callback main.WalkCallbackFunction)` |`package test_package ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(root go/ast.Node, callback main.WalkCallbackFunction)` |`()` |
| `func main.increaseLastChildIndex(childIndexTrace []int)` |`increaseLastChildIndex` |`func(childIndexTrace []int)` |`main.increaseLastChildIndex` |`func main.increaseLastChildIndex(childIndexTrace []int)` |`package test_package ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(childIndexTrace []int)` |`()` |
| `func main.isNodeNil(n go/ast.Node) bool` |`isNodeNil` |`func(n go/ast.Node) bool` |`main.isNodeNil` |`func main.isNodeNil(n go/ast.Node) bool` |`package test_package ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(n go/ast.Node)` |`(bool)` |
| `func main.walkAstTypeFieldsIfSet(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)` |`walkAstTypeFieldsIfSet` |`func(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)` |`main.walkAstTypeFieldsIfSet` |`func main.walkAstTypeFieldsIfSet(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars ...any)` |`package test_package ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction, vars []any)` |`()` |
| `func main.walkHelper(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)` |`walkHelper` |`func(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)` |`main.walkHelper` |`func main.walkHelper(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)` |`package test_package ("main")` |`%!s(*types.TypeParamList=<nil>)` |`<nil>` |`(n go/ast.Node, parentTrace []go/ast.Node, childIndexTrace []int, callback main.WalkCallbackFunction)` |`()` |

