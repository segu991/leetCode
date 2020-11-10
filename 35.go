package main

import "fmt"

func main() {
	var nums = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
}
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
