package help

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

//go:embed help.yaml
var helpYamlFileContent string

var helpFileContent map[string]string

func parseHelpFileContent() {
	err := yaml.NewDecoder(strings.NewReader(helpYamlFileContent)).Decode(&helpFileContent)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode help.yaml"))
	}
	// fmt.Printf("%#v\n", helpFileContent)
}

func terminate(msg any) {
	fmt.Println(msg)
	os.Exit(1)
}

type Command struct{}

func (c *Command) BindFlagsToArgs(f *flag.FlagSet) {}

func (c *Command) ManualParser(f *flag.FlagSet) error {
	return nil
}

func (c *Command) Run() {
	var cmdToRun string

	if len(os.Args) < 3 { // means running as "tde help"
		cmdToRun = "help"
	} else if cmdToRun == "help" { // avoid infinite recursion
		terminate("Help command helps you. Run \"tde help\" to learn more, instead \"tde help help\"")
	} else {
		cmdToRun = os.Args[2] // command part in "tde help <command>"
	}

	parseHelpFileContent()
	if msg, ok := helpFileContent[cmdToRun]; ok {
		fmt.Println(msg)
	} else {
		terminate("Unrecognized command for help. Run \"tde help\"")
	}
}
