package config_reader

import (
	"flag"
	"os"
	"reflect"
	"tde/internal/microservices/logger"
	"tde/internal/utilities"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var log = logger.NewLogger("ConfigReader")

type flags struct {
	Config      string
	GracePeriod time.Duration
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
	var flags = &flags{}

	flag.StringVar(&flags.Config, "config", "", "")

	flag.DurationVar(&flags.GracePeriod, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	log.Println("Parsing CLI args")
	flag.Parse()

	checkZeroValuedFields(flags)
	return flags
}

func FillAndReturn[T any](target *T) *T {
	var flags = getFlags()

	log.Printf("Reading '%s' as config file\n", flags.Config)
	var fileReadHandler, err = os.Open(flags.Config)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not open config file"))
	}
	err = yaml.NewDecoder(fileReadHandler).Decode(target)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Could not decode config file"))
	}

	checkZeroValuedFields(target)
	return target
}
