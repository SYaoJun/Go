// math_test.go

package main

import "testing"

func Add(a, b int) int {
	return a + b
}
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}
