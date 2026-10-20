func isAnagram(s string, t string) bool {
	sourceMap, destMap := make(map[rune]int), make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		if _, ok := sourceMap[c]; ok {
			sourceMap[c]++	
		} else {
			sourceMap[c] = 1
		}
	}
	for _, c := range t {
		if _, ok := destMap[c]; ok {
			destMap[c]++	
		} else {
			destMap[c] = 1
		}
	}

	for k, v := range sourceMap {
		if count, ok := destMap[k]; !ok || count != v {
			return false
		}
	}

	return true
}
