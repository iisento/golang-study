package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	s_slice := strings.Fields(s)
	m := make(map[string]int)
	for _, key := range s_slice {
		m[key] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
