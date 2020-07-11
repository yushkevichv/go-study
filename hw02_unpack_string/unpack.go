package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var tmp = rune('0')
	var s strings.Builder
	for _, runeValue := range str {
		if unicode.IsDigit(runeValue) && unicode.IsDigit(tmp) {
			return "", ErrInvalidString
		}

		if tmp == '0' {
			tmp = runeValue
			continue
		}

		if unicode.IsDigit(runeValue) {
			count, _ := strconv.Atoi(string(runeValue))
			s.WriteString(strings.Repeat(string(tmp), count))
			tmp = rune('0')
			continue
		} else {
			s.WriteRune(tmp)
			tmp = runeValue
			continue
		}
	}

	if tmp != '0' {
		s.WriteRune(tmp)
	}

	return s.String(), nil
}
