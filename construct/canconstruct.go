package construct

import (
	"strings"
)

func CanConstruct(str string, substrings []string) []string {
	if str == "" {
		return []string{}
	}

	for _, substring := range substrings {
		idx := strings.Index(str, substring)
		if idx == -1 {
			continue
		}
		res1 := CanConstruct(str[0:idx], substrings)
		res2 := CanConstruct(str[idx+len(substring):], substrings)
		if  res1 != nil && res2 != nil {
			res := append(res1, substring)
			res = append(res, res2...)
			return res
		}
	}

	return nil
}

func CanConstructIndex(str string, substrings []string) []string {
	memo := map[string][]string{}
	return cci(str, substrings, memo)
}

func cci(str string, substrings []string, memo map[string][]string) []string {
	if v, seen := memo[str]; seen {
		return v
	}

	if str == "" {
		return []string{}
	}

	for _, substring := range substrings {
		idx := strings.Index(str, substring)
		if idx != 0 {
			continue
		}

		res := cci(str[len(substring):], substrings, memo)
		if  res != nil {
			memo[str] = append(res, substring)
			return memo[str]
		}
	}

	memo[str] = nil
	return memo[str]
}

func CountConstruct(target string, wordbank []string) int {
	memo := map[string]int{}
	return cc(target, wordbank, memo)
}

func cc(target string, wordbank []string, memo map[string]int) int {
	if v, seen := memo[target]; seen {
		return v
	}

	if target == "" {
		return 1
	}

	var totalCount int
	for _, word := range wordbank {
		if strings.Index(target, word) == 0 {
			totalCount += cc(target[len(word):], wordbank, memo)
		}
	}

	memo[target] = totalCount
	return totalCount
}

func AllConstruct(target string, wordbank []string) [][]string {
	return ac(target, wordbank)
}

func ac(target string, wordbank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	var all [][]string
	for _, word := range wordbank {
		if strings.Index(target, word) == 0 {
			paths := ac(target[len(word):], wordbank)
			for i, _ := range paths {
				paths[i] = append(paths[i], word)
				all = append(all, paths[i])
			}
		}
	}

	return all
}

func AllConstructMemo(target string, wordbank []string) [][]string {
	memo := map[string][][]string{}
	return acmemo(target, wordbank, memo)
}

func acmemo(target string, wordbank []string, memo map[string][][]string) [][]string {
	if v, seen := memo[target]; seen {
		return v
	}

	if target == "" {
		return [][]string{{}}
	}

	var all [][]string
	for _, word := range wordbank {
		if strings.Index(target, word) == 0 {
			paths := acmemo(target[len(word):], wordbank, memo)
			for i, _ := range paths {
				paths[i] = append(paths[i], word)
			}
			all = append(all, paths...)
		}
	}

	memo[target] = all
	return all
}

func CanConstructTab(target string, wordbank []string) bool {
	tab := make([]bool, len(target)+1)
	tab[0] = true

	for i := 0; i < len(tab); i++ {
		if !tab[i] {
			continue
		}
		for _, w := range wordbank {
			if i+len(w) > len(target) {
				continue
			}

			if target[i:i+len(w)] == w {
				tab[i+len(w)] = true
			}
		}
	}
	return tab[len(target)]
}

func CountConstructTab(target string, wordbank []string) int {
	table := make([]int, len(target) + 1)
	table[0] = 1

	for i := 0; i < len(table); i++ {
		if table[i] > 0 {
			for _, w := range wordbank {
				start, end := i, i + len(w)
				if end < len(table) && target[start:end] == w {
					table[end] += table[i]
				}
			}
		}
	}

	return table[len(target)]
}

type Results [][]string
func AllConstructTab(target string, wordbank []string) Results {
	table := make([]Results, len(target)+1)
	table[0] = Results{[]string{}}

	for i := 0; i < len(table); i++ {
		if len(table[i]) > 0 {
			for _, w := range wordbank {
				start, end := i, i+len(w)
				if end < len(table) && target[start:end] == w {
					for _, combo := range table[i] {
						newCombo := make([]string, len(combo))
						copy(newCombo, combo)
						newCombo = append(newCombo, w)
						table[end] = append(table[end], newCombo)
					}
				}
			}
		}
	}

	return table[len(target)]
}