func lengthOfLongestSubstring(s string) int {
	var longest, left int
	// duplicate detection
	charIndex := make(map[rune]int)
	
	for right, char := range s {
		// if char was seen inside current window 
		if pos, ok := charIndex[char]; ok && pos >= left {
			left = pos + 1
		} 
		charIndex[char] = right
		longest = max(longest, right - left + 1)
	}

	return longest

}
