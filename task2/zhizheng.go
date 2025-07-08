package main

import "fmt"

func main() {
	a := 5
	fmt.Println("调用前的值", a)
	addten(&a)
	fmt.Println("调用后的值", a)

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("调用前的值", nums)
	mul2(&nums)
	fmt.Println("调用后的值", nums)
}

func mul2(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}
func addten(a *int) {
	*a += 10
}
