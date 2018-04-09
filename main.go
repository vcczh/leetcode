package main

import (
	"fmt"
)

func main() {
	// A, K := []int{9, 1, 2, 3, 9}, 3

	A, K := []int{1, 2, 3, 4, 5, 6, 7}, 4

	fmt.Println(largestSumOfAverages(A, K))
}
