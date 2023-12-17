**details**

CFG git hash: a337665

STG is non-CASTG non-ST-ASTGP

- **CFG** : Code fragment generator
- **STG** : Subtree Generator
- **CASTG** : Context-Aware STG
- **ST-ASTGP** : Strongly Typed ASTGP

**parent:**

```go
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, callback)
}
```

**code level candidates (28 out of 200 total)**

```go
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper((var870.), root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, callback, 0)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    select {
    goto BranchLabel4
    }
    walkHelper(root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, [var2610[((): 0)]]ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, ...var5541, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, &var8657, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, [...*]int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, var12003, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, [var12392.]int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, [var12393: *1]int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, [var15508]int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, [*var15761]ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    goto BranchLabel26
    walkHelper(root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    goto BranchLabel82
    walkHelper(root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, [0.794055288192076], callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, callback, (0.7387086491617374))
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, (......*var25839), []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, [&1]int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, (0), []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, callback, )
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, q, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    break BranchLabel51
    walkHelper(root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(var31120, root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(var31121[1]., root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(0, root, []ast.Node{}, []int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, callback, 0)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, [var35332]int{}, callback)
}
---
func WalkWithNils(root ast.Node, callback WalkCallbackFunction) {
    walkHelper(root, []ast.Node{}, []int{}, var35364, callback)
}
```