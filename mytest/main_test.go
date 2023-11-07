// main_test.go
package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -15},
		{"Mixed numbers", []int{-1, 2, -3, 4, -5}, -3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Sum(test.numbers)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
