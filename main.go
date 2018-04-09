package main

import (
	"fmt"
)

func main() {
	a := [][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{0, 2}, []int{2, 0}}
	fmt.Println(largestTriangleArea(a))
}
