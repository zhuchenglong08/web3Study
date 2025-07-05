package task1

func dealNumber(number []int) []int {
	length := len(number)
	//从最后一位开始
	for i := length - 1; i >= 0; i-- {
		if number[i] < 9 {
			number[i]++
			return number
		}
		number[i] = 0
	}
	//如果全为9 最前面加1
	return append([]int{1}, number...)
}
