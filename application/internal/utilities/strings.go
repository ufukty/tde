package utilities

func StringFold(str string, partLength int) (splitted []string) {
	for i := 0; i < len(str); i += partLength {
		splitted = append(splitted, str[i:Min(i+partLength, len(str))])
	}
	return
}
