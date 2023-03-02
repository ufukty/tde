package help

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"tde/internal/utilities"

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

type Command struct {
	Topic string `precedence:"0"`
}

func (c *Command) Run() {
	if c.Topic == "" {
		c.Topic = "help"
	}

	parseHelpFileContent()
	if msg, ok := helpFileContent[c.Topic]; ok {
		fmt.Println(msg)
	} else {
		utilities.Terminate("Unrecognized command for help. Run \"tde help\"")
	}
}
