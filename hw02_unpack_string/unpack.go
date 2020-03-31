package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	//s := strings.Split(input, )
	var b strings.Builder
	b.Grow(32)
	for i, rune := range input{
		var a strings.Builder
		a.Grow(4)
		fmt.Fprintf(&a, string(rune))
		fmt.Println(i, " ", rune, a.String())
		fmt.Fprintf(&b, "%s", rune)
		fmt.Println(unicode.IsDigit(rune))
	}
	fmt.Println(b.String())

	
	return "", nil
}


