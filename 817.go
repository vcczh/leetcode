package main

func numComponents(head *ListNode, G []int) int {
	count, inConnect := 0, false
	set := make(map[int]bool)
	for _, v := range G {
		set[v] = true
	}
	curr := head
	for curr != nil {
		if _, ok := set[curr.Val]; ok {
			if !inConnect {
				inConnect = true
				count += 1
			}
		} else {
			if inConnect {
				inConnect = false
			}
		}
		curr = curr.Next
	}
	return count
}
