# ast - ast wrapper

## Subpackages of astw

-   astw
    -   traced
        -   Contains: InspectWithTrace() and WalkWithNil()
    -   traverse
        -   Contains: Once(), Twice() and Node{}, suits the need accessing, storing nodes on AST including nil valued struct fields and slice items
    -   types:
        -   Contains: Enums for structs defined in standard library ast, used by both, exported.
        -   Used by first two package above and from outside packages
    -   utilities:
        -   Contains: loading, exporting ast from/to file, dir. listing subtree of a node, find etc.
        -   Used by first two package above and from outside packages
