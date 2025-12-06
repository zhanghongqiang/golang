package main

import (
	"fmt"
)

func main() {
	a := 5
	fmt.Println("改变前a的值：", a)
	num(&a)
	fmt.Println("改变后a的值：", a)

	b := []int{1, 2, 3, 4, 5}

	fmt.Println("改变前b的值：", b)

	arrSlice(&b)

	fmt.Println("改变后b的值：", b)
}

func num(i *int) {
	*i += 10
}

func arrSlice(num *[]int) {

	arr := *num

	for k := range arr {
		arr[k] *= 2
	}
}
