package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	p3 := l3
	p1 := l1
	p2 := l2

	total := 0
	carry := 0

	for p1 != nil || p2 != nil {
		number1 := 0
		if p1 != nil {
			number1 = p1.Val
			p1 = p1.Next
		}

		number2 := 0
		if p2 != nil {
			number2 = p2.Val
			p2 = p2.Next
		}

		total = number1 + number2 + carry
		carry = 0

		if total >= 10 {
			total = total % 10
			carry = 1
		}

		p3.Val = total
		if p1 != nil || p2 != nil {
			p3.Next = &ListNode{}
		} else if carry == 1 {
			p3.Next = &ListNode{1, nil}
		}
		p3 = p3.Next
	}

	return l3
}

//func main() {
//	l1 := &ListNode{Val: 2, Next: nil}
//	l1.Next = &ListNode{Val: 4, Next: nil}
//	l1.Next.Next = &ListNode{Val: 3, Next: nil}
//
//	l2 := &ListNode{Val: 5, Next: nil}
//	l2.Next = &ListNode{Val: 6, Next: nil}
//	l2.Next.Next = &ListNode{Val: 4, Next: nil}
//	addTwoNumbers(l1, l2)
//}
