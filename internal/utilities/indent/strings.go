package indent

import (
	"strings"
)

func Lines(str string, indent int) string {
	indentation := strings.Repeat(" ", indent)
	lines := []string{}
	for line := range strings.Lines(str) {
		lines = append(lines, indentation+line)
	}
	return strings.Join(lines, "\n")
}
