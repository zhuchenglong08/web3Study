package task1

func findSingle(nums []int) int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}
	for k, v := range counts {
		if v == 1 {
			return k
		}
	}
	panic("no one")
}
