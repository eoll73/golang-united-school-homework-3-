package main

import (
	"fmt"
	"sort"
)

func sortMapValues(input *map[int]string) []string {
	var output []string
	var keys []int
	for k := range *input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		output = append(output, (*input)[k])
	}
	return output
}

func main() {
	sortMapValuesInput := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	sortMapValuesRes := sortMapValues(&sortMapValuesInput)
	fmt.Printf("Sort Map Values result: %v\n", sortMapValuesRes)

	fmt.Println("done")
}
