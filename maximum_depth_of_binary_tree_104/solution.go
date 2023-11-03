package maximum_depth_of_binary_tree_104

import (
	. "github.com/austingebauer/go-leetcode/structures"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
