package main

import (
	"fmt"
)

// 1. 题目 ：编写一个Go程序，定义一个函数，
// 该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值。
func addTen(i *int) {
	*i += 10
}

// 2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func multSlices(i *[]int) {
	slice := *i
	for idx := range slice {
		slice[idx] *= 2
	}
}

func main() {
	i := 5
	fmt.Println("调用addTen之前i的值为：", i)
	addTen(&i)
	fmt.Println("调用addTen之后i的值为：", i)

	slice := []int{1, 2, 3}
	fmt.Println("调用multSlices之前i的值为：", slice)
	multSlices(&slice) // 传递的是地址 修改的是值 & *
	fmt.Println("调用multSlices之后i的值为：", slice)

}
