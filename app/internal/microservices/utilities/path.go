package utilities

import "strings"

func isSlashRune(r rune) bool { return r == '/' || r == '\\' }

func IsEvilPath(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, isSlashRune) {
		if ent == ".." {
			return true
		}
	}
	return false
}
