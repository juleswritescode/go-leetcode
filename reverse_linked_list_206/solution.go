package reverse_linked_list_206

import (
	s "github.com/austingebauer/go-leetcode/structures"
)

/* prev => current => next */

func reverseList(current *s.ListNode) (head *s.ListNode) {
	for current != nil {
		next := current.Next
		current.Next = head
		head = current
		current = next
	}

	return head
}

func reverseListRecur(current *s.ListNode) *s.ListNode {
	if current == nil || current.Next == nil {
		return current
	}

	newHead := reverseListRecur(current.Next)
	current.Next.Next = current
	current.Next = nil

	return newHead
}
