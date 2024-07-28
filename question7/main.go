package main

import "fmt"

//数组进行加一的操作，直接找9,根据9的个数不同分别做讨论
func plusOne(digits []int) []int {
	//首先获取到数组的长度
	n := len(digits)
	for i := n - 1; i >= 0; i-- { //从后向前进行遍历
		if digits[i] != 9 { //找到第一个不等于9的元素
			digits[i]++                  //将当前非9的元素加一
			for j := i + 1; j < n; j++ { //向后遍历如果出现那种199情况直接后面都变成0
				digits[j] = 0 //后面都为9的情况直接进位后面变零
			}
			return digits
		}
	}
	//如果出现所有数字都是9的情况直接后面都变零前面变1
	digits = make([]int, n+1)
	digits[0] = 1 //首位赋值为1
	return digits //直接返回值
}

func main() {
	digits := []int{9, 9}
	result := plusOne(digits)
	fmt.Println(result)
}
