package cypher

import (
	"strings"
	"bytes"
)

const (
	maxLetterInt = 122
	minLetterInt = 96
)

var lettersToInt = make(map[string]int)
var intToLetters = make(map[int]string)

func init() {
	letterString := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z"
	lettersSplit := strings.Split(letterString, ",")

	for idx, letter := range lettersSplit {
		lettersToInt[letter] = idx + 97
		intToLetters[idx + 97] = letter
	}
}

// Consume a message string, convert to lowercase and then to a slice of
// runes. From here we can use basic operators to shift the characters the
// desired amount.
func Shift(message string, shift int) (string, error) {
	originalSlice := bytes.Runes([]byte(strings.ToLower(message)))
	convertedSlice := []rune{}

	for _, char := range originalSlice {
		straightShift := int(char) + shift
		if letter, ok := intToLetters[straightShift]; ok {
			newChar := rune(lettersToInt[letter])
			convertedSlice = append(convertedSlice, newChar)
		} else {
			letter := intToLetters[(straightShift - maxLetterInt) + minLetterInt]
			newChar := rune(lettersToInt[letter])
			convertedSlice = append(convertedSlice, newChar)
		}
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