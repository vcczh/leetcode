package main

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	frequency := make(map[string]int)
	banMap := make(map[string]bool)
	max, maxWord := 0, ""

	for _, word := range banned {
		banMap[word] = true
	}

	words := strings.Split(paragraph, " ")
	for _, word := range words {
		word = removePunctuation(word)
		if _, ok := banMap[word]; !ok {
			f, contains := frequency[word]
			if !contains {
				f = 1
			} else {
				f += 1
			}
			frequency[word] = f
			if f > max {
				max, maxWord = f, word
			}
		}
	}
	return maxWord
}

func removePunctuation(word string) string {
	return strings.ToLower(strings.TrimFunc(word, func(c rune) bool {
		return !unicode.IsLetter(c)
	}))
}
