package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var last_rune rune = 0
	var r0 rune = '0'
	var r9 rune = '9'
	var sb strings.Builder
	sb.WriteString("")
	for _, k := range s {
		if k >= r0 && k <= r9 {
			if (last_rune >= r0 && last_rune <= r9 || last_rune == 0) && (!(k == r0 && last_rune == r0)) {
				return "", ErrInvalidString
			} else if !(last_rune >= r0 && last_rune <= r9) {
				sb.WriteString(strings.Repeat(string(last_rune), int(k)-int('0')))
			}
		} else if !(last_rune >= r0 && last_rune <= r9 || last_rune == 0) {
			sb.WriteString(string(last_rune))
		}
		last_rune = k
	}

	if !(last_rune >= r0 && last_rune <= r9) {
		sb.WriteString(string(last_rune))
	}

	return sb.String(), nil
}
