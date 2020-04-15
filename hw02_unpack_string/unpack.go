package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) { //nolint:gocognit,funlen
	var b strings.Builder
	var prevRune rune
	isScreenPrevRune := false

	if len(input) == 0 {
		return "", nil
	}

	for i, char := range input {
		if prevRune == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(prevRune) && unicode.IsDigit(char) && !isScreenPrevRune {
			return "", ErrInvalidString
		}

		if i == len(input)-1 && char == '\\' && isScreenPrevRune {
			return "", ErrInvalidString
		}

		if prevRune == '\\' && !isScreenPrevRune {
			if char == '\\' || unicode.IsDigit(char) {
				_, err := fmt.Fprintf(&b, "%s", string(char))
				isScreenPrevRune = true
				prevRune = char
				if err != nil {
					log.Println(err)
				}
				continue
			}
			return "", ErrInvalidString
		}

		if prevRune == '\\' && isScreenPrevRune {
			if char == '\\' {
				isScreenPrevRune = false
				continue
			}
		}

		if unicode.IsDigit(char) {
			repeat, err := strconv.Atoi(string(char))
			if err != nil {
				log.Println(err)
			}

			if repeat == 0 {
				temp := b.String()[0 : len(b.String())-1]
				b.Reset()
				_, err := fmt.Fprintf(&b, "%s", temp)
				if err != nil {
					log.Println(err)
				}
				prevRune = char
				continue
			}

			repeat--
			_, err = fmt.Fprintf(&b, "%s", strings.Repeat(string(prevRune), repeat))
			if err != nil {
				log.Println(err)
			}
			prevRune = char
			if isScreenPrevRune {
				isScreenPrevRune = false
			}
			continue
		}

		if char != '\\' {
			_, err := fmt.Fprintf(&b, "%s", string(char))
			if err != nil {
				log.Println(err)
			}
		}
		if char == '\\' {
			isScreenPrevRune = false
		}

		prevRune = char
	}
	return b.String(), nil
}
