package main

import (
	"fmt"
)

func average(input *[]int) float64 {
	if len(*input) == 0 {
		return float64(0)
	}
	sum := 0
	for _, v := range *input {
		sum += v
	}
	return float64(sum) / float64(len(*input))
}

func main() {
	averageInput := []int{1, 2, 3, 4, 5, 6}
	averageRes := average(&averageInput)
	fmt.Printf("Average result: %f\n", averageRes)
}
