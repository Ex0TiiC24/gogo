package main

import "fmt"

var nums = []int{3, 2, 4}

func main() {
	fmt.Print(twoSum(nums, 6))
}
func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); {
		for k := i + 1; k < len(nums); {
			if nums[i]+nums[k] == target {
				result = append(result, i, k)

			}
			k++
		}
		i++
	}
	return result
}
