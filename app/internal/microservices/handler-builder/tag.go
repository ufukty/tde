package handler_builder

import (
	"regexp"
)

func ParseTag(tag string) map[string]string {
	var pairs = map[string]string{}
	var keyValueScanner = regexp.MustCompile("[ ]*([A-Za-z0-9\\-_]+):\"([A-Za-z0-9\\-/]*)\"")
	var matches = keyValueScanner.FindAllStringSubmatch(tag, -1)
	for _, match := range matches {
		pairs[match[1]] = match[2]
	}
	return pairs
}
