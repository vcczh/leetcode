package main

import (
	"fmt"
	"math"
)

func main() {
	a := [][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{0, 2}, []int{2, 0}}
	fmt.Println(largestTriangleArea(a))
}

func largestTriangleArea(points [][]int) float64 {
	l, max := len(points), 0.0

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				area := calculateArea(points[i], points[j], points[k])
				max = math.Max(max, area)
			}
		}
	}
	return max
}

func calculateArea(a []int, b []int, c []int) float64 {
	// area = |Ax(By-Cy) + Bx(Cy-Ay) + Cx(Ay-By)| / 2
	return math.Abs(float64(a[0]*(b[1]-c[1])+b[0]*(c[1]-a[1])+c[0]*(a[1]-b[1]))) / 2
}
