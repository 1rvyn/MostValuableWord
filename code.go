package main

import (
	_ "fmt"
	"sort"
	"strings"
	_ "strings"
)

func main() {

	// get the largest value word
	// eg "way" = 23+1+25=49
	s := string("there is no way to do this asdasdsdzxczxczvzxc")

	store := string("_abcdefghijklmnopqrstuvwxyz")

	var m map[string]int
	m = make(map[string]int)

	words := strings.Fields(s)
	for _, word := range words {
		totals := 0
		for _, letter := range string(word) {
			for f, g := range store {
				if string(letter) == string(g) {
					totals = totals + f
				}
			}
		}
		m[word] = totals

	}

	counts := m
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] > counts[keys[j]] })
	println(keys[0])

}
