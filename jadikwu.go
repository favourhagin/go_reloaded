package main

// for go_reloaded project
//converting uppercase to lowercase
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
