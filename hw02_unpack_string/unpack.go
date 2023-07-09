package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var lastRune rune = 0
	r0 := rune('0')
	r9 := rune('9')
	sb := strings.Builder{}

	for _, k := range s {
		if k >= r0 && k <= r9 {
			if (lastRune >= r0 && lastRune <= r9 || lastRune == 0) && !(k == r0 && lastRune == r0) {
				return "", ErrInvalidString
			} else if !(lastRune >= r0 && lastRune <= r9) {
				sb.WriteString(strings.Repeat(string(lastRune), int(k)-int('0')))
			}
		} else if !(lastRune >= r0 && lastRune <= r9 || lastRune == 0) {
			sb.WriteString(string(lastRune))
		}
		lastRune = k
	}

	if !(lastRune >= r0 && lastRune <= r9) && lastRune > 0 {
		sb.WriteString(string(lastRune))
	}
	return sb.String(), nil
}
