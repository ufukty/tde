# `cfg` - Code Fragment Generators

## Interface

```go
type CFG interface {
    Develop(code any)
    Modify(code any)
}
```

`Develop(code any)` method expected to add 1 character/token/node to existing valid or invalid code. Not all type of CFGs required to create valid code at each `Develop` call. Type of code can be `string` or `ast.Node` depending on the CFG used in program.

## Types of CFGs

-   **Character based CFG**  
    Adds random characters to code.
-   **Token based CFG**  
    Adds random tokens to code. Tokens are keywords and operators that belong to the language syntax.
-   **AST based CFG**  
    Operates on abstract syntax tree of the code. All results expected to have valid syntax. Compile could fail, cause the produced code can contain access to undeclared variables, functions.
-   **Context-Aware AST based CFG**  
    In addition to AST based CFG, it accounts of previously declared variables, defined functions and imported libraries. When new statements and expressions needs to create; that stack is used as reference for the accessible variables, functions, methods. Less compile errors are expected from the results of this CFG.

From top type to bottom;

-   Longer execution period for `Develop` method, more time and memory complexity.
-   More care must be taken to implement a generator that has no unreachable spots on the solution space.
-   Less probability to produce gibberish code.
