package config_reader

import (
	"flag"
	"os"
	"reflect"
	"tde/internal/microservices/logger"
	"tde/internal/utilities"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var log = logger.NewLogger("ConfigReader")

type flags struct {
	Config string
	// GracePeriod time.Duration
}

func checkZeroValuedFields(subject any) {
	var valueOf = reflect.Indirect(reflect.ValueOf(subject))
	var typeOf = valueOf.Type()

	var nFields = typeOf.NumField()
	for i := range utilities.Range(nFields) {
		var fieldValue = valueOf.Field(i)
		var fieldType = typeOf.Field(i)

		if fieldValue.IsZero() {
			log.Fatalf("Field '%s' is set to zero-value: '%s'\n", fieldType.Name, fieldValue)
		}
	}
}

func getFlags() *flags {
	log.Println("Parsing CLI args")
	var flags = &flags{}
	flag.StringVar(&flags.Config, "config", "", "")
	flag.Parse()
	checkZeroValuedFields(flags)
	return flags
}

func GetConfig() *Config {
	var (
		flags           = getFlags()
		fileReadHandler *os.File
		err             error
		config          = &Config{}
	)
	log.Printf("Reading '%s' as config file\n", flags.Config)
	fileReadHandler, err = os.Open(flags.Config)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not open config file"))
	}
	err = yaml.NewDecoder(fileReadHandler).Decode(config)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not decode config file"))
	}
	checkZeroValuedFields(config)
	return config
}
