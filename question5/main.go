package main

import (
	"fmt"
	"sort"
)

// 找出数组中唯一的单个数字的函数
func singleNumber(nums []int) int {
	sort.Ints(nums) //首先给数组排序
	n := len(nums)  //获取数组的长度
	if n == 1 {
		return nums[0]
	}
	for i := 1; i < n; i++ {
		if nums[i-1] != nums[i] {
			if i == 1 {
				return nums[i-1]
			} else if i == n-1 {
				return nums[i]
			} else {
				i++
				if nums[i-1] != nums[i] {
					return nums[i-1]
				}
			}
		}
	}
	return -100000
}

// 利用异或运算求出数组
func main() {
	nums := []int{1}
	num := singleNumber(nums)
	fmt.Println(num)
}
