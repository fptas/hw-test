package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")


func Unpack(s string) (string, error) {
	var lastRune rune
	sb := strings.Builder{}
	var curRuneType int
	var lastRuneType int
	for _, k := range s {
		if k >= '0' && k <= '9' {
			curRuneType = 1
		} else {
			curRuneType = -1
		}

		if curRuneType > 0 {
			if (lastRuneType >= 0) && !(k == '0' && lastRune == '0') {
				return "", ErrInvalidString
			} else if lastRuneType < 0 {
				sb.WriteString(strings.Repeat(string(lastRune), int(k)-int('0')))
			}
		} else if lastRuneType < 0 {
			sb.WriteString(string(lastRune))
		}
		lastRune = k
		lastRuneType = curRuneType
	}
	if lastRuneType < 0 {
		sb.WriteString(string(lastRune))
	}
	return sb.String(), nil
}
