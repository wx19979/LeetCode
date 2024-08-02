package main

import "fmt"

//将零都挪到后面的函数,还是用双指针来解决该问题
func moveZeroes(nums []int) {
	n := len(nums)                    //获取数组长度
	slow := 0                         //初始化慢指针
	for fast := 0; fast < n; fast++ { //快指针遍历
		if nums[fast] != 0 { //如果当前是非零元素
			nums[slow] = nums[fast] //直接将快指针位置元素放置到慢指针位置
			slow++                  //慢指针向后移动

		}
	}
	for j := slow; j < n; j++ { //在慢指针后面都应该被赋值为0
		nums[j] = 0 //直接逐个赋值为0
	}
}
func main() {
	nums := []int{0, 2, 2, 0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
