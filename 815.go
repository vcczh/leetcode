package main

import (
	"sort"
)

type bus struct {
	Key       int
	Route     []int
	Adjacency []int
}

func createBus(g map[int]*bus, key int, route []int) {
	var b bus
	b.Key = key
	sort.Ints(route)
	b.Route = route
	b.Adjacency = make([]int, 0)
	g[b.Key] = &b
}

func isIntersect(a []int, b []int) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			return true
		}
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return false
}

func passBy(a []int, x int) bool {
	i := sort.SearchInts(a, x)
	return i != len(a) && a[i] == x
}

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	g := make(map[int]*bus)
	visited := make(map[int]int)
	queue := make([]int, 0)

	for i, route := range routes {
		createBus(g, i, route)
		for j := 0; j < i; j++ {
			if isIntersect(g[j].Route, route) {
				g[j].Adjacency = append(g[j].Adjacency, i)
				g[i].Adjacency = append(g[i].Adjacency, j)
			}
		}
		if passBy(route, S) {
			queue = append(queue, i)
			visited[i] = 1
			if passBy(route, T) {
				return 1
			}
		}
	}
	for len(queue) > 0 {
		current := g[queue[0]]
		queue = queue[1:]
		for _, v := range current.Adjacency {
			if passBy(g[v].Route, T) {
				return visited[current.Key] + 1
			}
			if _, contains := visited[v]; !contains {
				queue = append(queue, v)
				visited[v] = visited[current.Key] + 1
			}
		}
	}
	return -1
}
