package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getVal(ln *ListNode) int {
	if ln == nil {
		return 0
	} else {
		return ln.Val
	}
}

func getNext(ln *ListNode) *ListNode {
	if ln == nil || ln.Next == nil {
		return &ListNode{}
	} else {
		return ln.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	carry := 0
	current := result
	for {
		nodesSum := getVal(l1) + getVal(l2) + carry
		remainder := nodesSum % 10
		carry = nodesSum / 10
		current.Val = remainder
		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		} else {
			l1, l2 = getNext(l1), getNext(l2)
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return result
}
