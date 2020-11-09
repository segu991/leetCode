package main

import "fmt"

func main() {
	var nums = []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {

	var sum, leftSum int
	for _, v := range nums {
		sum += v
	}

	for i := 0; i < len(nums); i++ {

		if leftSum == (sum - leftSum - nums[i]) {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
