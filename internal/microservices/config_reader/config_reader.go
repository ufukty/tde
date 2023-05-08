package config_reader

import (
	"flag"
	"log"
	"os"
	"reflect"
	"tde/internal/utilities"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type flags struct {
	Config string
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
	log.Println("Reading CLI args to learn which path the config file in")
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
