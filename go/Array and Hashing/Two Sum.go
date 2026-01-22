package main

func twoSum(nums []int, target int) []int {

	counter := make(map[int]int)

	for i, val := range nums {

		diff := target - val
		if idx, exists := counter[diff]; exists {
			return []int{idx, i}
		}

		counter[val] = i
	}

	return []int{}
}
