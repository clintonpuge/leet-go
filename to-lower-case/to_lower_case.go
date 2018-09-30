package tolowercase

// ToLowerCase lower case string
func ToLowerCase(word string) string {
	lowerStrings := ""
	for _, char := range word {
		if char <= 90 && char >= 65 {
			lowerStrings += string(char + 32)
		} else {
			lowerStrings += string(char)
		}
	}
	return lowerStrings
	// r := []rune(word)
	// for i := 0; i < len(r); i++ {
	// 	if r[i] <= 90 && r[i] >= 65 {
	// 		r[i] += 32
	// 	}
	// }
	// return string(r)
}
