package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "A quick brown fox jumps over the lazy dog. The dog barks, and the fox runs away."

	wordCount := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		word := strings.ToLower(word)
		word = strings.Trim(word, ".,!?")
		wordCount[word]++
	}
	fmt.Println("wordCount...")
	for word, count := range wordCount {
		fmt.Printf("%v = %v\n", word, count)
	}

}
