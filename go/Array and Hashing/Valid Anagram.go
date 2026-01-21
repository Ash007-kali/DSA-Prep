package main

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	seen := make(map[rune]int)
	for _, item := range s {
		seen[item]++
	}

	for _, item := range t {

		if val, exists := seen[item]; !exists || val == 0 {
			return false
		}
		seen[item]--
	}

	return true

}
