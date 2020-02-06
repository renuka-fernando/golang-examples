package main

import (
	"fmt"
	"sort"
)

func main() {
	myMap := make(map[int]string)

	myMap[2] = "Two"
	myMap[4] = "Four"
	myMap[3] = "Three"
	myMap[1] = "One"
	myMap[5] = "Five"

	keys := make([]int, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, myMap[k])
	}
}
