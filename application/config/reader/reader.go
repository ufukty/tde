package reader

import (
	"tde/internal/microservices/logger"
	"tde/internal/utilities"

	"flag"
	"os"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var log = logger.NewLogger("ConfigReader")

type flags struct {
	Config string
	// GracePeriod time.Duration
}

func checkZeroValuedFieldsHelper(typeOf reflect.Type, valueOf reflect.Value) {
	var nFields = typeOf.NumField()
	for i := range utilities.Range(nFields) {
		var fieldValue = valueOf.Field(i)
		var fieldType = typeOf.Field(i)

		if fieldType.Type.Kind() == reflect.Struct {
			checkZeroValuedFieldsHelper(fieldType.Type, fieldValue)
		} else if fieldValue.IsZero() {
			log.Fatalf("Field '%s' is set to zero-value: '%s'\n", fieldType.Name, fieldValue)
		}
	}
}

func checkZeroValuedFields(subject any) {
	var valueOf = reflect.Indirect(reflect.ValueOf(subject))
	var typeOf = valueOf.Type()
	checkZeroValuedFieldsHelper(typeOf, valueOf)
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

func printConfigHelper(typeOf reflect.Type, valueOf reflect.Value, scopeStack []string) {
	var nFields = typeOf.NumField()
	for i := range utilities.Range(nFields) {
		var fieldValue = valueOf.Field(i)
		var fieldType = typeOf.Field(i)
		if fieldType.Type.Kind() == reflect.Struct {
			printConfigHelper(fieldType.Type, fieldValue, append(scopeStack, fieldType.Name))
		} else {
			log.Printf("%s/%s = %s\n", strings.Join(scopeStack, "/"), fieldType.Name, fieldValue)
		}
	}
}

func Print(subject any) {
	log.Println("")
	var valueOf = reflect.Indirect(reflect.ValueOf(subject))
	var typeOf = valueOf.Type()
	printConfigHelper(typeOf, valueOf, []string{})
	log.Println("")
}
