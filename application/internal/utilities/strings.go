package utilities

import "strings"

func StringFold(str string, partLength int) (splitted []string) {
	for i := 0; i < len(str); i += partLength {
		splitted = append(splitted, str[i:Min(i+partLength, len(str))])
	}
	return
}

func StringFill(c string, repeat int) string {
	ret := ""
	for i := 0; i < repeat; i++ {
		ret += c
	}
	return ret
}

func IndentLines(str string, indent int) string {
	indentation := StringFill(" ", indent)
	return strings.Join(Map(strings.Split(str, "\n"), func(i int, line string) string {
		return indentation + line
	}), "\n")
}
