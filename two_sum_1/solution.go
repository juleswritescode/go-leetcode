package two_sum_1

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		a := nums[i]

		for j := i + 1; j < l; j++ {
			b := nums[j]
			if a+b == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
