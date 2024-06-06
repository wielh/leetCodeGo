package main

func reverseList(head *ListNode) *ListNode {
	var answer *ListNode = nil
	for head != nil {
		answer = &ListNode{Val: head.Val, Next: answer}
		head = head.Next
	}
	return answer
}
