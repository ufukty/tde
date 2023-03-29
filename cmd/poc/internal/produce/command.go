package produce

import (
	"fmt"
	"tde/internal/command"
	"tde/internal/folders/discovery"

	"github.com/kr/pretty"
)

type Command struct {
	Timeout    int                 `long:"timeout" default:"10"`
	Runner     string              `long:"runner"`
	Continue   string              `long:"continue" short:"c" default:"10"`
	Model      string              `long:"model" default:"0.1"`
	Ratios     string              `long:"ratios" default:"10/1"`
	Population int                 `long:"population" default:"1000"`
	Iterate    int                 `long:"iterate" default:"10"`
	Size       int                 `long:"size" default:"1000"`
	Package    command.MultiString `long:"package" short:"p"`
	TestName   string              `precedence:"0"`
}

func (c *Command) Run() {

	discovery.FindModulePath()


	
	fmt.Println("hello world")
	pretty.Println(c)
}
