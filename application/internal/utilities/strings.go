package utilities

import "strings"

func StringFold(str string, partLength int) (splitted []string) {
	for i := 0; i < len(str); i += partLength {
		splitted = append(splitted, str[i:Min(i+partLength, len(str))])
	}
	return
}

func IndentLines(str string, indent int) string {
	indentation := ""
	for i := 0; i < indent; i++ {
		indentation += " "
	}
	return strings.Join(Map(strings.Split(str, "\n"), func(i int, line string) string {
		return indentation + line
	}), "\n")
}
