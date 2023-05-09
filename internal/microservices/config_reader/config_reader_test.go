package config_reader

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func Test_ReadConfig(t *testing.T) {
	os.Args = []string{os.Args[0], "-config", "test_config.yml"}

	var config = GetConfig()
	if config.Evolver.ServiceDiscoveryConfig != "75c31fcc-6dca-5e99-9bad-ea82ad9fe1f6" {
		t.Error("validation")
	}

	yaml.NewEncoder(os.Stdout).Encode(config)
}
