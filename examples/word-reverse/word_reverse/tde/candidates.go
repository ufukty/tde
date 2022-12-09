//go:build tde
// +build tde

package main

func candidate1(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}
