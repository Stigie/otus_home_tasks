package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

var errorUnmarshalString = errors.New("unmarshal error")

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var user User
	result := make(DomainStat)

	re, err := regexp.Compile("\\." + domain)
	if err != nil {
		return result, err
	}

	for scanner.Scan() {
		if err = json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return result, errorUnmarshalString
		}

		matched := re.MatchString(user.Email)

		if matched {
			value := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[value]++
		}
	}

	return result, nil
}
