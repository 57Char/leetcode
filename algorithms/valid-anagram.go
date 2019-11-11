package algorithms

// https://leetcode.com/problems/valid-anagram/

// ASCII
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[byte]int8)
	for i, v := range s {
		counter[byte(v)]++
		counter[byte(t[i])]--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}

// Unicode
func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[rune]int8)
	rs, rt := []rune(s), []rune(t)
	for i, v := range rs {
		counter[rune(v)]++
		counter[rune(rt[i])]--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}
