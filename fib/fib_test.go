package fib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tests = []struct{
	in int
	out int
} {
	{40, 102334155},	
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
}

func TestFib(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.out, fib(tc.in))
	}
}

func TestFibM(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.out, FibM(tc.in))
	}
}

func TestFibT(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.out, FibT(tc.in))
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

func BenchmarkFibM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibM(50)
	}
}

func BenchmarkFibT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibT(50)
	}
}

func BenchmarkFibNT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibNT(50)
	}
}