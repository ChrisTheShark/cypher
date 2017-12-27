package cypher

import (
	"strings"
	"bytes"
)

// Consume a message string, convert to lowercase and then to a slice of
// runes. From here we can use basic operators to shift the characters the
// desired amount.
func Shift(message string, shift int) (string, error) {
	originalSlice := bytes.Runes([]byte(strings.ToLower(message)))
	convertedSlice := []rune{}

	for _, char := range originalSlice {
		newChar := char + rune(shift % 26)
		convertedSlice = append(convertedSlice, newChar)
	}

	return string(convertedSlice), nil
}

func CountRunes(message string) (map[string]int, error) {
	characters := []byte(strings.ToLower(message))
	counts := map[string]int{}

	for _, char := range characters {
		charString := string(char)
		if _, ok := counts[charString]; ok {
			counts[charString]++
		} else {
			counts[charString] = 1
		}
	}

	return counts, nil
}