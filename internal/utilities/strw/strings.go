package strw

import (
	"strings"
	"tde/internal/utilities/functional"
)

func Fold(str string, partLength int) (splitted []string) {
	for i := 0; i < len(str); i += partLength {
		splitted = append(splitted, str[i:min(i+partLength, len(str))])
	}
	return
}

func Fill(c string, repeat int) string {
	ret := ""
	for i := 0; i < repeat; i++ {
		ret += c
	}
	return ret
}

func IndentLines(str string, indent int) string {
	indentation := Fill(" ", indent)
	return strings.Join(functional.Map(strings.Split(str, "\n"), func(i int, line string) string {
		return indentation + line
	}), "\n")
}

func EscapeLinefeeds(str string) string {
	return strings.Join(strings.Split(str, "\n"), " (\\n) ")
}
