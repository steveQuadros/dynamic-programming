package fib

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n - 1) + fib(n - 2)
}

func FibM(n int) int {
	memo := make([]int, n+1)
	for i, _ := range memo {
		memo[i] = -1
	}
	return fibm(n, memo)
}

func fibm(n int, memo []int) int {
	if v := memo[n]; v != -1 {
		return v
	}

	if n == 0 || n == 1 {
		memo[n] = n
		return n
	}

	result := fibm(n-2, memo) + fibm(n-1, memo)
	memo[n] = result
	return result
}

func FibT(n int) int {
	tab := make([]int, n + 1)

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			tab[i] = i
			continue
		}
		tab[i] = tab[i-2] + tab[i-1]
	}
	return tab[n]
}

func FibNT(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var lastlast, last, cur = 0, 1, 1
	for i := 2; i < n; i++ {
		lastlast =  last
		last = cur
		cur = last + lastlast
	}

	return cur
}