package main

import "testing"

func TestReactor(t *testing.T) {
	tests := map[string]string{
		"aA":               "",
		"abBA":             "",
		"abAB":             "abAB",
		"aabAAB":           "aabAAB",
		"dabAcCaCBAcCcaDA": "dabCBAcaDA",
	}

	for in, out := range tests {
		got := processChain(in)

		if got != out {
			t.Errorf("wrong output, expected '%v', got '%v' when processing '%v'", out, got, in)
		}
	}
}
