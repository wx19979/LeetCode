// 给你一个非严格递增排列的数组nums ,请你原地删除重复出现的元素,使每个元素只出现一次,返回删除后数组的新长度。
// 元素的相对顺序应该保持一致。然后返回nums中唯一元素的个数。
package main

import "fmt"

// 双指针删除重复元素
func removeDuplicates(nums []int) int {
	n := len(nums) //先获取数组长度
	if n == 0 {    //如果长度为0直接返回
		return 0
	}
	slow := 1                         //定义慢指针
	for fast := 1; fast < n; fast++ { //快指针进行循环
		if nums[fast] != nums[fast-1] { //如果快指针和其前面的内容不重复
			nums[slow] = nums[fast] //说明当前的值不是前面重复的值直接赋值给慢指针
			slow++                  //慢指针后移
		}
	}
	return slow //最后返回的慢指针位置就是长度
}

func main() {
	nums := [10]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums[:])
	fmt.Println(nums)
}
