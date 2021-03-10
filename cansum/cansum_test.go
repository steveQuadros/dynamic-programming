package cansum

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var cstests = []struct{
	ts int
	nums []int
	expected bool
} {
	{7, []int{5,3,4,7}, true},
	{7, []int{2,4}, false},
	{2, []int{2,4}, true},
	{20, []int{2,4,5,10,20}, true},
}

func TestCanSum(t *testing.T) {
	testCanSum(t, CanSum)
}

func BenchmarkCanSum(b *testing.B) {
	benchmarkCanSum(b, CanSum)
}

func TestCanSumTab(t *testing.T) {
	testCanSum(t, CanSumTab)
}

func BenchmarkCanSumTab(b *testing.B) {
	benchmarkCanSum(b, CanSumTab)
}

func TestCanSumDP(t *testing.T) {
	testCanSum(t, CanSumDP)
}

func BenchmarkCanSumDP(b *testing.B) {
	benchmarkCanSum(b, CanSumDP)
}

func TestHasPath(t *testing.T) {
	tests := []struct{
		ts int
		nums []int
		path []int
	} {
		{
			ts: 7, nums: []int{5,3,4,7}, path: []int{4,3},
		},
		{ts: 20, nums: []int{10, 2}, path: []int{10,10}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("sum: %d, nums: %v", tc.ts, tc.nums), func(t *testing.T) {
			assert.Equal(t, tc.path, HasSum(tc.ts, tc.nums))
		})
	}
}

func TestHowSumTab(t *testing.T) {
	tests := []struct{
		ts int
		nums []int
		all [][]int
	} {
		{ts: 7, nums: []int{5,3,4,7}, all: [][]int{{4,3}, {3,4}}},
		{ts: 20, nums: []int{10, 2}, all: [][]int{{10,10}, {2,2,2,2,2,2,2,2,2,2}}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("sum: %d, nums: %v", tc.ts, tc.nums), func(t *testing.T) {
			assert.Contains(t, tc.all, HowSumTab(tc.ts, tc.nums))
		})
	}
}

func TestBestSum(t *testing.T) {
	tests := []struct{
		ts int
		nums []int
		path []int
	} {
		{
			ts: 7, 
			nums: []int{5,3,4,7}, 
			path: []int{7},
		},
		{
			ts: 20, 
			nums: []int{2, 10}, 
			path: []int{10,10},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.ts), func(t *testing.T) {
			assert.Equal(t, tc.path, BestSum(tc.ts, tc.nums))
		})
	}
}

func TestBestSumTab(t *testing.T) {
	tests := []struct{
		ts int
		nums []int
		path []int
	} {
		{
			ts: 7, 
			nums: []int{5,3,4,7}, 
			path: []int{7},
		},
		{
			ts: 20, 
			nums: []int{2, 10}, 
			path: []int{10,10},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.ts), func(t *testing.T) {
			assert.Equal(t, tc.path, BestSumTab(tc.ts, tc.nums))
		})
	}
}

func TestAllSum(t *testing.T) {
	tests := []struct{
		ts int
		nums []int
		all [][]int
	} {
		{
			ts: 7, 
			nums: []int{5,3,4,7}, 
			all: [][]int{{4, 3}, {3,4}, {7}},
		},
		{
			ts: 20, 
			nums: []int{2, 10}, 
			all: [][]int{{2,2,2,2,2,2,2,2,2,2}, {2,2,2,2,2,10}, {10,10}, {2,2,2,2,10,2}, {2,2,2,10,2,2}, {2,2,10,2,2,2}, {2,10,2,2,2,2}, {10,2,2,2,2,2}},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.ts), func(t *testing.T) {
			assert.ElementsMatch(t, tc.all, AllSum(tc.ts, tc.nums))
		})
	}
}

func testCanSum(t *testing.T, cs func(int, []int)bool) {
	for _, tc := range cstests {
		t.Run(fmt.Sprintf("sum: %d, nums: %v", tc.ts, tc.nums), func(t *testing.T) {
			assert.Equal(t, tc.expected, cs(tc.ts, tc.nums))
		})
	}
}

func benchmarkCanSum(b *testing.B, cs func(int, []int)bool) {
	for i := 0; i < b.N; i++ {
		cs(20, []int{2,4,5,10,20})
	}
}