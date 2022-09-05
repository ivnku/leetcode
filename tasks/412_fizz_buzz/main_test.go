package main

import (
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		num    int
		answer []string
	}{
		{
			3,
			[]string{"1", "2", "Fizz"},
		}, {
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		}, {
			15,
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for i, tc := range testCases {
		result := fizzBuzz(tc.num)
		if !equalTwoSlices(result, tc.answer) {
			t.Errorf("Wrong answer in test case #%d! Expected: %#v, Actual: %#v\n", i+1, tc.answer, result)
		}
	}
}

func equalTwoSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, el := range a {
		if strings.Compare(el, b[i]) != 0 {
			return false
		}
	}

	return true
}
