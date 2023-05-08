package config_reader

import (
	"flag"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

type TestConfig struct {
	Test string
}

func Test_ReadConfig(t *testing.T) {
	flag.Set("config", "test.yml")

	var config = FillAndReturn(&TestConfig{})
	if config.Test != "49af1177-1cf5-506c-b99d-1fb90ba26c62" {
		t.Error("validation")
	}

	yaml.NewEncoder(os.Stdout).Encode(config)
}
