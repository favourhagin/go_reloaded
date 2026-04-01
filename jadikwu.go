package main

func ToLower(text string) string {
	runes := []rune(text)
	for i, char := range runes {
		runes[i] = unicode.ToLower(char)
	}
	return string(runes)
}

// for go_reloaded 
func ToLower(text string) string {
	words := strings.Fields(text)
	char := []string{}
	for _, i := range words {
		if i == "(low)" && len(char) > 0 {
			char[len(char)-1] = strings.ToLower(char[len(char)-1])
		} else if i != "(low)" {
			char = append(char, i)
		}
	}
	return strings.Join(char, " ")
}
