package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	} else if left == 1 {
		end := head
		index := 1
		tmp := []int{}
		for end != nil && index <= right {
			tmp = append(tmp, end.Val)
			end = end.Next
			index++
		}

		for i := 0; i < len(tmp); i++ {
			end = &ListNode{Val: tmp[i], Next: end}
		}
		return end
	} else {
		start := head
		end := head
		index := 1
		tmp := []int{}

		for end != nil {
			if index == 1 {
				end = end.Next
			} else if index < left {
				start = start.Next
				end = end.Next
			} else if index >= left && index <= right {
				tmp = append(tmp, end.Val)
				end = end.Next
			} else {
				break
			}
			index++
		}

		for i := 0; i < len(tmp); i++ {
			end = &ListNode{Val: tmp[i], Next: end}
		}
		start.Next = end
		return head
	}
}
