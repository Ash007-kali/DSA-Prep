package main

func containsDuplicate(nums []int) bool {

	isDuplicate := make(map[int]bool) //init an empty map

	for _, item := range nums {
		_, exists := isDuplicate[item]
		if exists {
			return true
		}
		isDuplicate[item] = true
	}
	return false
}
