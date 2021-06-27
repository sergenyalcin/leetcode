package easy

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		index, ok := m[target-v]

		if ok {
			return []int{index, i}
		}

		m[v] = i
	}

	return nil
}
