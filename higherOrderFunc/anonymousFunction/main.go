package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"apple", "banana", "orange", "grape"}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Println(words) // Output: [apple grape banana orange]
}
