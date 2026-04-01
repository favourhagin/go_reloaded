package main


var start int

func fixQuotes(input string) string {
	inQuotes := false
	result := ""
	quotes := rune(0)

	for _, item := range input {
		if !inQuotes && (item != '\'' && item != '"') {
			result += string(item)
		} else if !inQuotes && (item == '\'' || item == '"') {
			quotes = item
			result += string(item)
			inQuotes = true

		} else if inQuotes && item == quotes {
			result += string(item)
			inQuotes = false

		} else if inQuotes {
			if item != ' ' {
				result += string(item)
			}
		}

	}
	return string(result)

}


