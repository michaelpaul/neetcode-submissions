type Counts [26]byte

// Time: O(N*M) => N: len(strs), M: 100
// Space: O(N*M) => N: len(strs), M: 100
func groupAnagrams(strs []string) [][]string {
	groups := make(map[Counts][]string)
	
	for _, v := range strs {
		var key Counts
		for j := 0; j < len(v); j++ {
			pos := int(v[j]) - 'a'
			key[pos]++
		}

		if group, ok := groups[key]; ok {
			groups[key] = append(group, v)
		} else {
			groups[key] = []string{v}
		}
	}

	result := [][]string{}
	for _, v := range groups {
		result = append(result, v)
	}
	return result 
}
