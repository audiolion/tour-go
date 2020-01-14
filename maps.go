/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
	word_count := make(map[string]int)
	for _, word := range words {
		_, ok := word_count[word]
		if ok {
		  word_count[word] += 1
		} else {
		  word_count[word] = 1
		}
	}
	return word_count
}

func main() {
	wc.Test(WordCount)
}
