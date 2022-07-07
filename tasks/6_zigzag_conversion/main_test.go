package main

import "testing"

type TestCase struct {
	input  string
	result string
	rows   int
}

func TestZigzagConversion(t *testing.T) {
	tests := []TestCase{
		{"PAYPALISHIRING", "PAHNAPLSIIGYIR", 3},
		{"PAYPALISHIRING", "PINALSIGYAHRPI", 4},
		{"A", "A", 1},
		{"AB", "AB", 1},
	}

	for index, test := range tests {
		result := convert(test.input, test.rows)
		if result != test.result {
			t.Errorf("%d Test failed! Expected value: %s, result is: %s", index, test.result, result)
		}
	}
}
