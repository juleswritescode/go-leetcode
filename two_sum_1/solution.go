package two_sum_1

func twoSum(nums []int, target int) []int {
	checked := map[int]int{}

	for idx, n := range nums {
		missing := target - n

		missingIdx, ok := checked[missing]
		if ok {
			return []int{missingIdx, idx}
		} else {
			checked[n] = idx
		}
	}

	return []int{}
}
