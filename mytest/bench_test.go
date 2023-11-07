package main

import "testing"

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F(100, 10)
	}
}
func F(a, b int) int {
	return a + b
}
