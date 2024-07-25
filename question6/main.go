// 两个数组的交集二
package main

import "fmt"

//用于根据角标删除数组元素的函数
func del_num(num []int, index int) {
	num1 := num[index:]
	num = num[:index-1]
	num = append(num, num1...)
}

// //用于查找两个数组的交集函数
// func intersect(nums1 []int, nums2 []int) []int {
// 	//首先获取到两个数组的长度,初始化两个数组的指针
// 	n1, n2, p1, p2, long, short := len(nums1), len(nums2), 0, 0, 0, 0
// 	var num3 []int
// 	if n1 > n2 { //先判断哪个短然后就用哪个作为结束长度
// 		long = n1
// 		short = n2
// 	} else {
// 		long = n1
// 		short = n2
// 	}
// 	for p2 < short { //利用short进行对比

// 	}
// }
func main() {
	nums := []int{1, 2, 3, 4}
	del_num(nums, 2)
	fmt.Println(nums)
}
