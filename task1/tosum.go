package task1

func tosum(nums []int, target int) []int {
	//if len(nums) == 0 {
	//	return nil
	//}
	//var res []int
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target && i != j {
	//			res = []int{i, j}
	//		}
	//	}
	//}
	//return res
	var res = map[int]int{}
	for i, num := range nums {
		tagetnum := target - num
		index, ok := res[tagetnum]
		if ok {
			return []int{index, i}
		} else {
			res[num] = i
		}
	}
	return []int{}
}
