package cansum

import (
	// "fmt"
)

func CanSum(ts int, nums []int) bool {
	if ts == 0 {
		return true
	} else if ts < 0 {
		return false
	}

	for _, n := range nums {
		if CanSum(ts - n, nums) {
			return true
		}
	}
	return false
}

func CanSumDP(ts int, nums []int) bool {
	memo := map[int]bool{}
	return csr(ts, nums, memo)
}

func csr(ts int, nums []int, memo map[int]bool) bool {
	if v, seen := memo[ts]; seen {
		return v
	}

	if ts == 0 {
		return true
	} else if ts < 0 {
		return false
	}

	for _, n := range nums {
		if csr(ts-n, nums, memo) {
			memo[ts] = true
			return true
		}
	}

	memo[ts] = false
	return false
}

func HasSum(ts int, nums []int) []int {
	memo := map[int][]int{}
	path := hs(ts, nums, memo)
	return path
}

func hs(ts int, nums []int, memo map[int][]int) []int {
	if v, seen := memo[ts]; seen {
		return v
	}

	if ts == 0 {
		return []int{}
	} else if ts < 0 {
		return nil
	}

	for _, n := range nums {
		if path := hs(ts-n, nums, memo); path != nil {
			memo[ts] = append(path, n)
			return memo[ts]
		}
	}

	memo[ts] = nil
	return nil
}

func BestSum(ts int, nums []int) []int {
	var best []int
	return bs(ts, nums, best)
}

func bs(ts int, nums, best []int) []int {
	if ts == 0 {
		return []int{}
	} else if ts < 0 {
		return nil
	}

	for _, n := range nums {
		rc := bs(ts-n, nums, best)
		if rc != nil {
			path := append(rc, n)
			if best == nil || len(best) > len(path){
				best = path
			}			
		}
	}

	return best
}

func AllSum(ts int, nums []int) [][]int {
	return as(ts, nums)
} 

func as(s int, nums []int) [][]int {
	if s == 0 {
		return [][]int{{}}
	}
	
	var all [][]int
	for _, n := range nums {
		if s-n < 0{
			continue
		}

		paths := as(s-n, nums)
		for i, _ := range paths {
			paths[i] = append(paths[i], n)
			all = append(all, paths[i])
		}
	}

	return all
}

func CanSumTab(target int, nums []int) bool {
	tab := make([]bool, target + 1)
	for i, _ := range tab {
		if i == 0 {
			tab[i] = true
		} else {
			tab[i] = false
		}
	}

	for i, can := range tab {
		if can {
			for _, n := range nums {
				if i + n <= target {
					tab[i+n] = true
				}
			}
		}
		
	}
	return tab[target]
}

func HowSumTab(target int, nums []int) []int {
	tab := make([][]int, target + 1)
	for i, _ := range tab {
		if i == 0 {
			tab[i] = []int{}
		} else {
			tab[i] = nil
		}
	}

	for i := 0; i < len(tab); i++ {
		if tab[i] != nil {
			for _, n := range nums {
				if i + n < len(tab) {
					tab[i+n] = append(tab[i], n)
				}
			}
		}
	}

	return tab[target]
}

func BestSumTab(target int, nums []int) []int {
	tab := make([][]int, target + 1)
	for i := 0; i < len(tab); i++ {
		if i == 0 {
			tab[i] = []int{}
		} else {
			tab[i] = nil
		}
	}

	for i := 0; i < len(tab); i++ {
		if tab[i] != nil {
			for _, n := range nums {
				if i + n < len(tab) {
					curpath := append(tab[i], n)
					if tab[i+n] == nil || len(curpath) < len(tab[i+n]) {
						tab[i+n] = curpath
					}
				}
			}
		}
	}
	return tab[target]
}