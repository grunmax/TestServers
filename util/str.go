package util

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func RegSplit(text string, delimeter string) []string {
	text = strings.TrimSpace(text)
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	if len(text) == 0 {
		return []string{}
	}
	laststart := 0
	list := make([]string, len(indexes)+1)
	for i, element := range indexes {
		list[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	list[len(indexes)] = text[laststart:len(text)] // + last word
	return list
}

func RuneCountFilterList(list []string, minCount int) []string {
	var result []string
	for _, word := range list {
		if utf8.RuneCountInString(word) >= minCount {
			result = append(result, word)
		}
	}
	return result
}
