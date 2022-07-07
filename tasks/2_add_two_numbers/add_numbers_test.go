package main

import (
	"fmt"
	"strconv"
	"testing"
)

type TestCase struct {
	name     string
	x        []int
	y        []int
	expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []TestCase{
		{
			"Test Case 1",
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		}, {
			"Test Case 2",
			[]int{0},
			[]int{0},
			[]int{0},
		}, {
			"Test Case 3",
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tc := range testCases {
		x := createListNode(tc.x, nil)
		y := createListNode(tc.y, nil)

		result := addTwoNumbers(x, y)
		if err := checkResult(result, tc.expected); err != nil {
			t.Errorf("In %v, error: %v", tc.name, err.Error())
		}
	}
}

func ListNodeToStr(ln ListNode) string {
	var result string
	for {
		result += strconv.Itoa(ln.Val)
		if ln.Next == nil {
			break
		}
		ln = *ln.Next
	}
	return result
}

func IntSliceToStr(sl []int) string {
	var result string
	var item int
	for {
		item, sl = sl[0], sl[1:]
		result += strconv.Itoa(item)
		if len(sl) == 0 {
			break
		}
	}
	return result
}

func checkResult(result *ListNode, x []int) error {
	resStr := ListNodeToStr(*result)
	expStr := IntSliceToStr(x)

	curr := result
	for {
		expectedNodeVal := x[0]
		x = x[1:]
		if curr.Val != expectedNodeVal {
			return fmt.Errorf("the result node %v is not equal expected node %v! Exp: %v, Res: %v", curr.Val, expectedNodeVal, expStr, resStr)
		}
		if len(x) == 0 && curr.Next != nil {
			return fmt.Errorf("the expected value digits amount is less then result! Exp: %v, Res: %v", expStr, resStr)
		}
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}
	if len(x) != 0 {
		return fmt.Errorf("the expected value digits amount is greater then result! Exp: %v, Res: %v", expStr, resStr)
	} else {
		return nil
	}
}

func createListNode(x []int, ln *ListNode) *ListNode {
	if ln == nil {
		ln = &ListNode{}
	}
	ln.Val, x = x[0], x[1:]
	if len(x) == 0 {
		ln.Next = nil
		return ln
	}
	ln.Next = createListNode(x, ln.Next)
	return ln
}
