package main

func ToLower(text string) string {
	runes := []rune(text)
	for i, char := range runes {
		runes[i] = unicode.ToLower(char)
	}
	return string(runes)
}
