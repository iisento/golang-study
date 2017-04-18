package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount ...
// 単語を数える関数
func WordCount(s string) map[string]int {
	sSlice := strings.Fields(s)
	m := make(map[string]int)
	for _, key := range sSlice {
		m[key]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
