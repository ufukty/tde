package help

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed help.yaml
var helpYamlFileContent string

func Run() error {
	topic := "main"
	if len(os.Args) > 0 {
		topic = os.Args[0]
	}

	helpFileContent := map[string]string{}
	err := yaml.NewDecoder(strings.NewReader(helpYamlFileContent)).Decode(&helpFileContent)
	if err != nil {
		return fmt.Errorf("decoding yaml: %w", err)
	}

	msg, ok := helpFileContent[topic]
	if !ok {
		return fmt.Errorf(`Unrecognized command for help. Run "tde help"`)
	}
	fmt.Println(msg)

	return nil
}
