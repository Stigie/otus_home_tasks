package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

type dictionaryItem struct {
	Key   string
	Value int
}

const topCounts = 10

func Top10(input string) []string {
	re := regexp.MustCompile(`(?i)([^\W_]*([\wа-яА-я]+(-[\wа-яА-я]+)+)*[^\s,\-.!?'"\x60]*)`)

	dictionary := make(map[string]int)

	for _, match := range re.FindAllString(input, -1) {
		index := strings.ToLower(match)
		count, ok := dictionary[index]
		if ok {
			count++
			dictionary[index] = count
		}
		if !ok && index != "" {
			dictionary[index] = 1
		}
	}

	dictionarySlice := make([]dictionaryItem, len(dictionary))

	i := 0
	for k, v := range dictionary {
		dictionarySlice[i] = dictionaryItem{
			Key:   k,
			Value: v,
		}
		i++
	}
	sort.Slice(dictionarySlice, func(i, j int) bool {
		return dictionarySlice[i].Value > dictionarySlice[j].Value
	})

	resultLength := topCounts
	if length := len(dictionarySlice); length < topCounts {
		resultLength = length
	}

	dictionarySlice = dictionarySlice[:resultLength]

	res := make([]string, resultLength)

	for i := range res {
		res[i] = dictionarySlice[i].Key
	}

	return res
}
