package main

import (
	"fmt"
	"sort"
)

// 找出数组中唯一的单个数字的函数
func singleNumber(nums []int) int {
	sort.Ints(nums) //首先给数组排序
	n := len(nums)  //获取数组的长度
	if n == 1 {     //如果数组只有一个元素直接返回该元素
		return nums[0]
	}
	for i := 1; i < n; i++ { //循环找到唯一个一个不重复的元素
		if nums[i-1] != nums[i] { //如果出现相邻两个元素不等的情况
			if i == 1 { //如果当前元素是数组中第二个证明第一个就是要找的元素直接返回即可
				return nums[i-1]
			} else if i == n-1 { //如果当前元素是数组最后一个那么直接返回数组的最后一个元素
				return nums[i]
			} else { //其他无特殊情况
				i++                       //直接指针向后移动
				if nums[i-1] != nums[i] { //如果出现不等的情况
					return nums[i-1] //直接返回当前元素的前一个元素
				}
			}
		}
	}
	return -100000000 //如果找不到直接返回一个无意义的数字
}

// 利用异或运算求出数组
func main() {
	nums := []int{1}
	num := singleNumber(nums)
	fmt.Println(num)
}
