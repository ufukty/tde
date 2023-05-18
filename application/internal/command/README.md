# Command

Organizer for flag parser

## Usage from main

```go
package main

import (
    "tde/internal/command"

    "tde/cmd/client/help"
    "tde/cmd/client/produce"
)

func main() {
    command.RegisterCommand("help", help.Command{})
    command.RegisterCommand("produce", produce.Command{})

    command.Route()
}
```

## Command interface

```go
type Command interface {
	Run()
}
```

## Example command implementation

```go
type Command struct {
    PositionalArgument1 string `precedence="0"`
    PositionalArgument2 int    `precedence="1"`
    NamedArgument1      string `long="named1" default="foobar"`
    NamedArgument2      int    `long="named2" short="n"`
}

func (c Command) Run() {}
```

## Notes

-   Not all positional arguments has to be passed from CLI.
-   When more than enough positional arguments are passed, it is an error.
-   Default values will be used if specified for the positional arguments that enough value are not passed from CLI.
-   Precedences are not required to be successive.
-   Named arguments are not required to have default values specified. When no value passed, zero-values for related type will be used.
-   For multi-use named flags, `command.MultiString` can be used as type to related struct field.


