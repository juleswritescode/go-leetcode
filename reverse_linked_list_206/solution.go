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

func reverseListRecur(head *s.ListNode) *s.ListNode {
	if head.Next == nil {
		return head
	}
}
