package kata

import (
	"strings"
)

const (
	dash   = '-'
	lodash = '_'
)

func toCamelCase(s string) string {
	newString := ""

	currentIndex := 0
	lastIndex := len(s) - 1
	if lastIndex < 0 {
		return newString
	}

	if s[0] == dash || s[0] == lodash {
		s = s[1:]
		lastIndex--
	}
	if s[lastIndex] == dash || s[lastIndex] == lodash {
		s = s[:lastIndex]
		lastIndex--
	}
	for currentIndex <= lastIndex {
		letter := s[currentIndex]
		if letter == dash || letter == lodash {
			currentIndex++
			letter = s[currentIndex]
			newString += strings.ToUpper(string(letter))
		} else {
			newString += string(letter)

		}
		currentIndex++
	}
	return newString
}
