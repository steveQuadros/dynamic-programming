package gridtraveler

import (
	"strconv"
)

func GridTraveler(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 && n == 1 {
		return 1
	}
	return GridTraveler(m-1, n) + GridTraveler(m, n-1)
}

func GridTravelerDP(m, n int) int {
	memo := make([][]int, m + 1, m + 1)
	for i, _ := range memo {
		memo[i] = make([]int, n + 1, n + 1)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}

	return gtr(m, n, memo)
}

func gtr(m, n int, memo [][]int) int {
	var paths int
	if paths = memo[m][n]; paths != -1 {
		return paths
	}

	if m == 0 || n == 0 {
		paths = 0
	} else if m == 1 && n == 1 {
		paths = 1
	} else {
		paths = gtr(m-1, n, memo) + gtr(m, n-1, memo)
	}
	memo[m][n] = paths
	return paths
}

func GridTravelerTab(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	}

	tab := make([][]int, m+1, m+1)
	for i := 0; i < m+1; i++ {
		tab[i] = make([]int, n+1, n+1)
	}

	tab[1][1] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i+1 <= m {
				tab[i+1][j] += tab[i][j]
			}
			if j+1 <= n {
				tab[i][j+1] += tab[i][j]
			}
		}
	}

	return tab[m][n]
}

func GridTravelerMap(m, n int) int {
	mp := map[string]int{}
	return gtrm(m, n, mp)
}

func gtrm(m, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if path, ok := memo[key]; ok {
		return path
	}

	if m == 1 && n == 1 {
		return 1
	} else if m == 0 || n == 0 {
		return 0
	}

	path := gtrm(m-1, n, memo) + gtrm(m, n-1, memo)
	memo[key] = path
	return path
}
