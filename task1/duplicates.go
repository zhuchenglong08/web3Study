package task1

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}

	}
	fmt.Println(nums)
	return i + 1
}
