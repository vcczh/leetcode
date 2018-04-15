package main

type Status struct {
	Pos, Speed, Step int
}

func racecarBFS(target int) int {
	n, max := 0, 1
	for max < target {
		max = max << 1
		n += 1
	}
	queue := make([]Status, 0)
	start := Status{0, 1, 0}
	visited := make(map[Status]bool, 0)
	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Pos == target {
			return curr.Step
		}
		queue = append(queue, acc(curr), rev(curr))
	}
	return -1
}

func acc(s Status) (r Status) {
	r.Pos, r.Speed, r.Step = s.Pos+s.Speed, s.Speed*2, s.Step+1
	return
}

func rev(s Status) (r Status) {
	if s.Speed > 0 {
		r.Pos, r.Speed, r.Step = s.Pos, -1, s.Step+1
	} else {
		r.Pos, r.Speed, r.Step = s.Pos, 1, s.Step+1
	}
	return
}

func racecar(target int) int {
	// DP
	opt := make([]int, target+1)
	for i := 1; i <= target; i++ {
		opt[i] = maxInt()
	}
	for i, n := 1, uint(1); i <= target; i++ {
		if i == 1<<n-1 {
			opt[i] = int(n)
			n++
			continue
		}
		if 1<<n-1-i < i {
			opt[i] = min(opt[i], opt[int(1<<n-1-i)]+int(n)+1)
		}
		for m := uint(0); m < n-1; m++ {
			opt[i] = min(opt[i], opt[int(i-1<<(n-1)+1<<m)]+int(n+m)+1)
		}

	}
	return int(opt[target])
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxInt() int {
	return int(^uint(0) >> 1)
}
