package words

import "strings"

// CountWords
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
