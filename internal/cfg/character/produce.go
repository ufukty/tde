package character

import "math/rand"

var AllowedCharacters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789()[]{}_.=&!+-*/%:;\n ")

func ProduceRandomFragment() []byte {
	fragment := []byte{}
	length := len(AllowedCharacters) // last one is for termination
	for i := 0; i != length; i = rand.Intn(length + 1) {
		fragment = append(fragment, AllowedCharacters[i])
	}
	return fragment
}
