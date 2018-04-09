package main

import (
	"math"
)

func largestSumOfAverages(A []int, K int) float64 {
	return largestSumOfAveragesByDP(A, K)
	// return largestSumOfAveragesSinceX(A, K, 0)
}

func largestSumOfAveragesSinceX(A []int, K int, X int) float64 {
	// Recursive solution
	if K == 1 {
		sum := 0
		for i := X; i < len(A); i++ {
			sum += A[i]
		}
		return float64(sum) / float64((len(A) - X))
	}
	max, sum := 0.0, 0
	for i := X; i < len(A)-K+1; i++ {
		sum += A[i]
		max = math.Max(max, float64(sum)/float64((i-X+1))+largestSumOfAveragesSinceX(A, K-1, i+1))
	}
	return max
}

func largestSumOfAveragesByDP(A []int, K int) float64 {
	length := len(A)
	opt := make([][]float64, K)
	for i := range opt {
		opt[i] = make([]float64, length)
	}
	for sum, i := 0.0, 0; i < length; i++ {
		sum += float64(A[i])
		opt[0][i] = sum / float64(i+1)
	}
	for k := 1; k < K; k++ {
		for i := k; i < length; i++ {
			sum := float64(A[i])
			for j := i - 1; j >= k-1; j-- {
				opt[k][i] = math.Max(opt[k][i], opt[k-1][j]+sum/float64(i-j))
				sum += float64(A[j])
			}
		}
	}
	return opt[K-1][length-1]
}
