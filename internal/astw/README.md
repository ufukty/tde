# ast - ast wrapper

## Subpackages of astw

-   astw
    -   traced
        -   Contains: InspectWithTrace() and WalkWithNil()
    -   traverse
        -   Contains: Traverse() and TraversableNode{}, suits the need accessing, storing nodes on AST including nil valued struct fields
    -   types:
        -   Contains: Enums for structs defined in standard library ast, used by both, exported.
        -   Used by first two package above and from outside packages
    -   utilities:
        -   Contains: loading, exporting ast from/to file, dir. listing subtree of a node, find etc.
        -   Used by first two package above and from outside packages
