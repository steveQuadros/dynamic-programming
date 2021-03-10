package construct

import (
	"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

var canConstructTests = []struct{
	str string
	substrings []string
	path []string
} {
	{"abcdef", []string{"ab", "abc", "de","def", "abcd"}, []string{"abc", "def"}},
	{"dedefdec", []string{"de", "f", "c"}, []string{"de", "de", "de", "f", "c"}},
	{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, []string{}},
}

func TestCanConstruct(t *testing.T) {
	for _, tc := range canConstructTests {
		t.Run(tc.str, func(t *testing.T) {
			assert.ElementsMatch(t, tc.path, CanConstruct(tc.str, tc.substrings))
		})
	}
}

func TestCanConstructTab(t *testing.T) {
	var tests = []struct{
		str string
		substrings []string
		path []string
		can bool
	} {
		{"abcdef", []string{"ab", "abc", "de","def", "abcd"}, []string{"abc", "def"}, true},
		{"dedefdec", []string{"de", "f", "c"}, []string{"de", "de", "de", "f", "c"}, true},
		{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, []string{}, false},
		{"purple", []string{"purp", "p", "ur", "le", "purpl", "e", "l"}, []string{}, true},
	}

	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			assert.Equal(t, tc.can, CanConstructTab(tc.str, tc.substrings))
		})
	}
}

func TestCanConstructIndex(t *testing.T) {
	for _, tc := range canConstructTests {
		t.Run(tc.str, func(t *testing.T) {
			assert.ElementsMatch(t, tc.path, CanConstructIndex(tc.str, tc.substrings))
		})
	}
}

func TestCountConstruct(t *testing.T) {
	var tests = []struct{
		target string
		wordbank []string
		numWays int
	} {
		{"abcdef", []string{"ab", "abc", "de","def", "abcd"}, 1},
		{"dedefdec", []string{"de", "f", "c"}, 1},
		{"skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, 0},
		{"purple", []string{"purp", "p", "ur", "le", "purpl", "e", "l"}, 5},
	}

	for _, tc := range tests {
		t.Run(tc.target, func(t *testing.T) {
			assert.Equal(t, tc.numWays, CountConstructTab(tc.target, tc.wordbank))
		})
	}
}

func TestAllConstruct(t *testing.T) {
	tests := []struct{
		target string
		wordbank []string
		res [][]string
	} {
		{"purple", []string{"purp", "p", "ur", "le", "purpl", "e", "l"}, [][]string{
			{"le", "purp"},
			{"e", "l", "purp"},
			{"le", "p", "p", "ur"},
			{"e", "l", "p", "p", "ur"},
			{"e", "purpl"},
		}},
		{"cat", []string{"dig", "mouse"}, [][]string{}},
	}

	for _, tc := range tests {
		t.Run(tc.target, func(t *testing.T) {
			expected := tc.res
			sort.Slice(expected, func(i, j int) bool {
				return len(expected[i]) < len(expected[j])
			})
			for _, combo := range expected {
				sort.Strings(combo)
			}

			res := AllConstructTab(tc.target, tc.wordbank)
			sort.Slice(res, func(i, j int) bool {
				return len(res[i]) < len(res[j])
			})
			for _, combo := range res {
				sort.Strings(combo)
			}
			assert.ElementsMatch(t, expected, res)
		})
	}

	for _, tc := range tests {
		t.Run("TAB " + tc.target, func(t *testing.T) {
			expected := tc.res
			sort.Slice(expected, func(i, j int) bool {
				return len(expected[i]) < len(expected[j])
			})
			for _, combo := range expected {
				sort.Strings(combo)
			}

			res := AllConstructTab(tc.target, tc.wordbank)
			sort.Slice(res, func(i, j int) bool {
				return len(res[i]) < len(res[j])
			})
			for _, combo := range res {
				sort.Strings(combo)
			}
			assert.ElementsMatch(t, expected, res)
		})
	}
}

func BenchmarkCanConstruct(b *testing.B) {
	tc := canConstructTests[2]
	for i := 0; i < b.N; i++ {
		CanConstruct(tc.str, tc.substrings)
	}
}

func BenchmarkCanConstructIndex(b *testing.B) {
	tc := canConstructTests[2]
	for i := 0; i < b.N; i++ {
		CanConstructIndex(tc.str, tc.substrings)
	}
}

func BenchmarkAllConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl", "e", "l"})
	}
}

func BenchmarkAllConstructMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllConstructMemo("purple", []string{"purp", "p", "ur", "le", "purpl", "e", "l"})
	}
}