package main

import (
	"errors"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

// GetUTFLength возвращает длину строки UTF8 и ошибку ErrInvalidUTF8 (в случае возникновения).
func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
