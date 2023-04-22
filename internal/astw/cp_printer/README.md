# `cp_printer`

-   `cp_printer` is for `copy & paste printer`

-   `ast_inspect` tool can be used to get `cp_printer` output for a file.

-   It prints an ast.Node in way that when the output pasted in a Go file, it will be compiled perfectly.

-   That paste can be used to have an in-code representation of a Go code.

```sh
$ cat <some_file> | ast_inspect cp
# output to copy to Go file
```

```Go
package main

import (
    "fmt"
    "go/ast"
    "go/token"
)

var f = // paste the output here

func main() {
    f.Decl... // use it interactively
}
```
