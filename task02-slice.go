package main

import "fmt"

func reverse(input *[]int64) []int64 {
	var output []int64
	if len(*input) == 0 {
		return output
	}
	for i := len(*input) - 1; i >= 0; i-- {
		output = append(output, (*input)[i])
	}
	return output
}

func main() {
	reverseInput := []int64{1, 2, 5, 15}
	reverseRes := reverse(&reverseInput)
	fmt.Printf("Reverce result: %v\n", reverseRes)

	fmt.Println()
}
