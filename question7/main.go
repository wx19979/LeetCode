package main

import "fmt"

//数组进行加一的操作，直接找9,根据9的个数不同分别做讨论
func plusOne(digits []int) []int {
	n := len(digits)              //获取到长度
	for i := n - 1; i >= 0; i-- { //从后向前依次遍历
		if digits[i] != 9 { //找到数组中第一个不为9的元素直接加一
			digits[i]++ //直接加一
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func main() {
	digits := []int{1, 2, 3}
	result := plusOne(digits)
	fmt.Println(result)
}
