package words

func WordReverseSemantic(w string) string {
	return ""
}

func WordReverse(w string) string {
	runes := []rune(w)
	reversed := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}

func WordReverseEnchanced(w string) string {
	inputRunes := []rune(w)
	length := len(inputRunes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		inputRunes[i], inputRunes[j] = inputRunes[j], inputRunes[i]
	}
	return string(inputRunes)
}
