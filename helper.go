package main

func isVowel(c string) bool {
	switch c {
	case "a", "e", "i", "o", "u", "h",
		"A", "E", "I", "O", "U", "H":
		return true
	}
	return false
}

func isLower(c string) bool {
	switch c {
	case "a", "e", "i", "o", "u":
		return true
	}
	return false
}
