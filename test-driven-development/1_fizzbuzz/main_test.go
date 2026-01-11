package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"3の倍数", 3, "Fizz"},
		{"5の倍数", 5, "Buzz"},
		{"15の倍数", 15, "FizzBuzz"},
		{"3の倍数(6)", 6, "Fizz"},
		{"5の倍数(10)", 10, "Buzz"},
		{"15の倍数(30)", 30, "FizzBuzz"},
		{"3でも5でもない倍数(1)", 1, "1"},
		{"3でも5でもない倍数(2)", 2, "2"},
		{"3でも5でもない倍数(7)", 7, "7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fizzbuzz(tt.input)
			if result != tt.expected {
				t.Errorf("fizzbuzz(%d) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}
