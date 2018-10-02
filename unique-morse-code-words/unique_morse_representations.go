package uniquemorsecodewords

// UniqueMorseRepresentations return the number of different transformations among all words we have.
func UniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	mappedMorse := mapMorseToLetter(morseCode)
	transformations := []string{}
	for _, char := range words {
		transformations = append(transformations, convertToMorse(string(char), mappedMorse))
	}
	length := len(removeDuplicates(transformations))
	return length
}

func removeDuplicates(words []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, word := range words {
		if _, value := keys[word]; !value {
			keys[word] = true
			list = append(list, string(word))
		}
	}
	return list
}

func convertToMorse(word string, code map[string]string) string {
	convertedStr := ""
	for _, char := range word {
		convertedStr += code[string(char)]
	}
	return convertedStr
}

func mapMorseToLetter(morse []string) map[string]string {
	morseMap := make(map[string]string)
	for i, code := range morse {
		morseMap[string(toChar(i+1))] = code
	}
	return morseMap
}

func transformToMorse() {}

func toChar(i int) rune {
	return rune('a' - 1 + i)
}
