func topKFrequent(nums []int, k int) []int {
	// build freq map of nums O(N)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// make a list of (n, freq) O(N)
	var pairs []pair
	for n, f := range freq {
		pairs = append(pairs, [2]int{n, f})
	}

	// sort the list desc by freq O(N log N)
	sort.Slice(pairs, func (i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	// return top K	O(N)
	var result []int
	for _, v := range pairs[:k] {
		result = append(result, v[0])
	}
	return result
}

type pair [2]int
