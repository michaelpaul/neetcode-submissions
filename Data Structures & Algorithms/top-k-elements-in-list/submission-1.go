func topKFrequent(nums []int, k int) []int {
	// build freq map of nums	O(N)
	count := make(map[int]int) // n => count
	for _, v := range nums {
		count[v]++
	}

	// group by (freq, numbers)	O(N)
	groupByCount := make(map[int][]int) // count => []n
	var top int 
	for n, c := range count {
		groupByCount[c] = append(groupByCount[c], n)
		top = max(top, c)
	}

	// take K elements from the groups in reverse order	O(n)
	var result []int
	for i := top; i > 0; i-- {
		if group, ok := groupByCount[i]; ok {
			for _, v := range group {
				result = append(result, v)
				if len(result) == k {
					return result
				}
			}
		}
	}

	return result
}
