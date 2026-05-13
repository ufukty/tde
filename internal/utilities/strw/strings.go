package strw

import (
	"strings"

	"tde/internal/utilities/functional"
)

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
