package main

func getConcatenation(nums []int) []int {

	nums = append(nums, nums...) //syntax to append to slice
	return nums
}
