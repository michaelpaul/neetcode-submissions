func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if j, ok := dict[complement]; ok {
			return []int{j, i}
		}
		dict[v] = i
	}

	return nil 
}
