package main

import "testing"

func TestLongestSubstr(t *testing.T) {
	testCases := []struct {
		str string
		len int
	}{
		{
			"abcabcbb",
			3,
		}, {
			"bbbbb",
			1,
		}, {
			"pwwkew",
			3,
		}, {
			" ",
			1,
		}, {
			"dvdf",
			3,
		}, {
			"dvdfgrvsk",
			7,
		}, {
			"",
			0,
		},
	}

	for i, tc := range testCases {
		result := longestSubstr(tc.str)
		if result != tc.len {
			t.Errorf("Wrong answer in test case #%d (%s)! Expected: %d, Actual: %d\n", i+1, tc.str, tc.len, result)
		}
	}
}
