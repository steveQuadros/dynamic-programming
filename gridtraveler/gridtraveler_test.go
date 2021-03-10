package gridtraveler

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var tests = []struct{
	r, c int
	paths int
} {
	{0, 1, 0},
	{1, 0, 0},
	{1, 1, 1},
	{1, 2, 1},
	{2, 2, 2},
	{3, 3, 6},
	{5, 5, 70},
}

func TestGridTravelerDP(t *testing.T) {	
	testGT(t, GridTravelerDP)
}

func BenchmarkGridTravelerDP(t *testing.B) {
	benchGT(t, GridTravelerDP)
}

func TestGridTraveler(t *testing.T) {
	testGT(t, GridTraveler)
}

func BenchmarkGridTraveler(t *testing.B) {
	benchGT(t, GridTraveler)
}

func TestGridTravelerTab(t *testing.T) {
	testGT(t, GridTravelerTab)
}

func BenchmarkGridTravelerTab(t *testing.B) {
	benchGT(t, GridTravelerTab)
}

func TestGridTravelerMap(t *testing.T) {
	testGT(t, GridTravelerMap)
}

func BenchmarkGridTravelerMap(t *testing.B) {
	benchGT(t, GridTravelerMap)
}

func testGT(t *testing.T, gt func(int, int) int) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d, %d should = %d", tc.r, tc.c, tc.paths), func(t *testing.T) {
			assert.Equal(t, tc.paths, gt(tc.r, tc.c))
		})
	}
}

func benchGT(t *testing.B, gt func(int, int) int) {
	for i := 0; i < t.N; i++ {
		gt(10, 10)
	}
}