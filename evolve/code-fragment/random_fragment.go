package codefragment

import "math/rand"

var allowedCharacters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:;\n ")

func ProduceRandomFragment() []byte {
	fragment := []byte{}
	length := len(allowedCharacters) // last one is for termination
	for i := 0; i != length; i = rand.Intn(length + 1) {
		fragment = append(fragment, allowedCharacters[i])
	}
	return fragment
}
