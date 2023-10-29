package reverse_linked_list_206

import (
	"slices"
	"testing"

	s "github.com/austingebauer/go-leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	t.Run("helpers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}

		assert.Equal(t, in, unmakeLL(makeLL(in)))
	})

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "reverse linked list",
			args: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "reverse linked list",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "reverse linked list",
			args: []int{},
			want: []int{},
		},
		{
			name: "reverse linked list",
			args: []int{1, 2},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r1 := unmakeLL(reverseList(makeLL(tt.args)))
			assert.Equal(t, tt.want, r1)
			r2 := unmakeLL(reverseListRecur(makeLL(tt.args)))
			assert.Equal(t, tt.want, r2)
		})
	}
}

func makeLL(nums []int) *s.ListNode {
	cp := make([]int, len(nums))
	copy(cp, nums)
	slices.Reverse(cp)

	var tmp *s.ListNode
	for _, n := range cp {
		n := s.ListNode{
			Val:  n,
			Next: tmp,
		}
		tmp = &n
	}

	return tmp
}

func unmakeLL(n *s.ListNode) []int {
	arr := []int{}

	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}

	return arr
}
