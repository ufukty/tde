package command

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type Command interface {
	Run()
}

var (
	commands = map[string]Command{}
	flagSets = map[string]*flag.FlagSet{}
)

func RegisterCommand(commandName string, cmd Command) {
	flagSet := flag.NewFlagSet(commandName, flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Printf("Run help: \"tde help %s\"\n", commandName)
	}

	commands[commandName] = cmd
	flagSets[commandName] = flagSet
}

func rawToInt(raw string) (int, error) {
	if raw == "" {
		return 0, nil
	}
	out, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errors.Wrap(err, "invalid value for integer")
	}
	return out, nil
}

func rawToFloat64(raw string) (float64, error) {
	if raw == "" {
		return 0.0, nil
	}
	out, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0.0, errors.Wrap(err, "invalid value for boolean")
	}
	return out, nil
}

func rawToBool(raw string) (bool, error) {
	if raw == "" {
		return false, nil
	}
	out, err := strconv.ParseBool(raw)
	if err != nil {
		return false, errors.Wrap(err, "invalid value for float64")
	}
	return out, nil
}

func resolveNamedFlag(f *flag.FlagSet, fieldType reflect.StructField, fieldValue reflect.Value) error {
	names := []string{}
	if short, ok := fieldType.Tag.Lookup("short"); ok {
		names = append(names, short)
	}
	if long, ok := fieldType.Tag.Lookup("long"); ok {
		names = append(names, long)
	}

	defaultValueRaw := fieldType.Tag.Get("default")
	for _, flagName := range names {

		switch fieldType.Type.Kind() {

		case reflect.Int:
			defaultValueFinal, err := rawToInt(defaultValueRaw)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to parse default value for field: %s", fieldType.Name))
			}
			f.IntVar(fieldValue.Addr().Interface().(*int), flagName, defaultValueFinal, "")

		case reflect.String:
			f.StringVar(fieldValue.Addr().Interface().(*string), flagName, defaultValueRaw, "")

		case reflect.Bool:
			defaultValueFinal, err := rawToBool(defaultValueRaw)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to parse default value for field: %s", fieldType.Name))
			}
			f.BoolVar(fieldValue.Addr().Interface().(*bool), flagName, defaultValueFinal, "")

		case reflect.Float64:
			defaultValueFinal, err := rawToFloat64(defaultValueRaw)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to parse default value for field: %s", fieldType.Name))
			}
			f.Float64Var(fieldValue.Addr().Interface().(*float64), flagName, defaultValueFinal, "")

		case reflect.TypeOf(MultiString{}).Kind():
			f.Var(fieldValue.Addr().Interface().(*MultiString), flagName, "")

		default:
			return errors.New(fmt.Sprint("unsupported type for argument struct field: ", fieldType.Name))
		}
	}

	return nil
}

type FlagKind int

const (
	Positional = FlagKind(iota)
	Named
)

func getFlagKind(fieldType reflect.StructField) FlagKind {
	if _, ok := fieldType.Tag.Lookup("precedence"); ok {
		return Positional
	}
	return Named
}

func registerStructFieldsAsFlags(f *flag.FlagSet, cmd any) {
	typeOf := reflect.TypeOf(cmd)
	valueOf := reflect.ValueOf(cmd)

	if typeOf.Kind() == reflect.Pointer {
		valueOf = reflect.Indirect(reflect.ValueOf(cmd))
		typeOf = valueOf.Type()
	}

	numberOfFields := typeOf.NumField()
	for i := 0; i < numberOfFields; i++ {
		fieldType := typeOf.Field(i)
		fieldValue := valueOf.Field(i)

		if getFlagKind(fieldType) == Named {
			resolveNamedFlag(f, fieldType, fieldValue)
		}
	}
}

func ordinalNumber(number int) string {
	str := fmt.Sprintf("%d", number)
	absolute := int(math.Abs(float64(number)))
	if i := absolute % 100; i == 11 || i == 12 || i == 13 {
		return str + "th"
	}
	switch absolute % 10 {
	case 1:
		return str + "st"
	case 2:
		return str + "nd"
	case 3:
		return str + "rd"
	default:
		return str + "th"
	}
}

func parsePositionalArguments(f *flag.FlagSet, cmd Command) error {
	args := f.Args()

	typeOf := reflect.TypeOf(cmd)
	valueOf := reflect.ValueOf(cmd)
	if typeOf.Kind() == reflect.Pointer {
		valueOf = reflect.Indirect(reflect.ValueOf(cmd))
		typeOf = valueOf.Type()
	}
	type ArgField struct {
		TypeOf     reflect.StructField
		ValueOf    reflect.Value
		Precedence int
	}
	positionalFields := []ArgField{}
	numberOfFields := typeOf.NumField()
	for i := 0; i < numberOfFields; i++ {
		fieldType := typeOf.Field(i)
		fieldValue := valueOf.Field(i)

		if getFlagKind(fieldType) == Positional {
			if precedence, ok := fieldType.Tag.Lookup("precedence"); ok {
				if precedenceInt, err := rawToInt(precedence); err == nil {
					positionalFields = append(positionalFields, ArgField{
						TypeOf:     fieldType,
						ValueOf:    fieldValue,
						Precedence: precedenceInt,
					})
				} else {
					return errors.New("Precedence value of field \"" + fieldType.Name + "\" is not valid integer.")
				}
			}
		}
	}
	slices.SortStableFunc(positionalFields, func(r, l ArgField) int {
		return l.Precedence - r.Precedence
	})

	if len(args) > len(positionalFields) {
		return errors.New(fmt.Sprintf("Too many positional arguments. Can accept %d values; got %d.",
			len(positionalFields), len(args)))
	}

	for i := 0; i < min(len(args), len(positionalFields)); i++ {
		var (
			positionalField = positionalFields[i]
			arg             = args[i]
		)

		switch positionalField.TypeOf.Type.Kind() {

		case reflect.Int:
			argInt, err := rawToInt(arg)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf(
					"%s positional argument (value: \"%s\") supposed to be used for argument \"%s\" but its type should be %s",
					ordinalNumber(i+1), arg, positionalField.TypeOf.Name, positionalField.TypeOf.Type.Kind(),
				))
			}
			positionalField.ValueOf.SetInt(int64(argInt))

		case reflect.String:
			positionalField.ValueOf.SetString(arg)

		case reflect.Bool:
			argBool, err := rawToBool(arg)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf(
					"%s positional argument (value: \"%s\") supposed to be used for argument \"%s\" but its type should be %s",
					ordinalNumber(i+1), arg, positionalField.TypeOf.Name, positionalField.TypeOf.Type.Kind(),
				))
			}
			positionalField.ValueOf.SetBool(argBool)

		case reflect.Float64:
			argFloat64, err := rawToFloat64(arg)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf(
					"%s positional argument (value: \"%s\") supposed to be used for argument \"%s\" but its type should be %s",
					ordinalNumber(i+1), arg, positionalField.TypeOf.Name, positionalField.TypeOf.Type.Kind(),
				))
			}
			positionalField.ValueOf.SetFloat(argFloat64)

		default:
			return errors.New(fmt.Sprintf("Given positional argument's type (value: %s) is unsupported for the precedenced argument: %s",
				arg, positionalField.TypeOf.Name))
		}
	}

	// if given args are not enough to fill every positional args, then start to use defaults
	for i := len(args); i < len(positionalFields); i++ {
		positionalField := positionalFields[i]
		defaultValueRaw, ok := positionalField.TypeOf.Tag.Lookup("default")
		if !ok {
			continue
		}

		switch positionalField.TypeOf.Type.Kind() {

		case reflect.Int:
			defaultValueFinal, err := rawToInt(defaultValueRaw)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to parse default value for field: %s", positionalField.TypeOf.Name))
			}
			positionalField.ValueOf.SetInt(int64(defaultValueFinal))

		case reflect.String:
			positionalField.ValueOf.SetString(defaultValueRaw)

		case reflect.Bool:
			defaultValueFinal, err := rawToBool(defaultValueRaw)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to parse default value for field: %s", positionalField.TypeOf.Name))
			}
			positionalField.ValueOf.SetBool(defaultValueFinal)

		case reflect.Float64:
			defaultValueFinal, err := rawToFloat64(defaultValueRaw)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to parse default value for field: %s", positionalField.TypeOf.Name))
			}
			positionalField.ValueOf.SetFloat(defaultValueFinal)

		default:
			return errors.New(fmt.Sprint("unsupported type for argument struct field: ", positionalField.TypeOf.Name))
		}
	}

	return nil
}

func Route() {
	if len(os.Args) < 2 {
		log.Fatalln("Expected command. Run \"tde help\"")
	}

	for name, cmd := range commands {
		registerStructFieldsAsFlags(flagSets[name], cmd)
	}

	cmdToRun := os.Args[1]
	if cmd, ok := commands[cmdToRun]; ok {
		f := flagSets[cmdToRun]

		if err := f.Parse(os.Args[2:]); err != nil {
			if cmdToRun == "help" {
				cmdToRun = ""
			}
			log.Fatalln(errors.Wrap(err, fmt.Sprintf("Could not parse arguments. Run \"tde help %s\"", cmdToRun)))
		}

		if err := parsePositionalArguments(f, cmd); err != nil {
			if cmdToRun == "help" {
				cmdToRun = ""
			}
			log.Fatalln(errors.Wrap(err, fmt.Sprintf("Could not parse arguments. Run \"tde help %s\"", cmdToRun)))
		}

		cmd.Run()
		os.Exit(0)
	} else {
		log.Fatalln("Unrecognized command. Run \"tde help\"")
	}

}
