package main

import (
	"fmt"
	"sort"
)

// 找出数组中是否存在重复元素的函数
func containsDuplicate1(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		fast := i + 1
		for j := fast; j < n; j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums) //首先将数组排序
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
func main() {
	nums := []int{1, 2, 3}
	isD1 := containsDuplicate1(nums)
	fmt.Println(isD1)
	isD2 := containsDuplicate2(nums)
	fmt.Println(isD2)
}
