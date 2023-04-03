package produce

import (
	"fmt"
	"tde/internal/command"
	"tde/internal/folders/discovery"
	"tde/internal/utilities"

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
	Package    command.MultiString `long:"package" short:"p"` // packages allowed to import
	TestName   string              `precedence:"0"`
}

func (c *Command) Run() {
	modulePath, err := discovery.GetModulePath()
	if err != nil {
		utilities.Terminate("Could not find module root. Are you in a Go module?")
	}

	if c.TestName == "" {
		utilities.Terminate("Test name is not specified. Run tde help")
	}

	testDetails, err := discovery.ResolveTestDetailsInCurrentDir(c.TestName)
	if err != nil {
		utilities.Terminate("Could not find the test named", c.TestName, err)
	}

	pretty.Println(c)
	fmt.Println(modulePath)
	pretty.Println(testDetails)

}
