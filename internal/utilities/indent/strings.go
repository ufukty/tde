package indent

import (
	"strings"

	"tde/internal/utilities/functional"
)

func Lines(str string, indent int) string {
	indentation := strings.Repeat(" ", indent)
	return strings.Join(functional.Map(strings.Split(str, "\n"), func(i int, line string) string {
		return indentation + line
	}), "\n")
}
