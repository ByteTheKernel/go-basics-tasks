package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-10, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{7, true},
		{9, false},
		{29, true},
		{49, false},
		{97, true},
		{100, false},
	}

	for _, tt := range tests {
		got := isPrime(tt.n)

		if got != tt.want {
			t.Errorf(
				"isPrime(%d) = %t want %t",
				tt.n, got, tt.want,
			)
		}
	}
}
