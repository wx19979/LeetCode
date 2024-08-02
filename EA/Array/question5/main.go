package main

import (
	"fmt"
	"sort"
)

// 找出数组中唯一的单个数字的函数
func singleNumber1(nums []int) int {
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

// 利用异或运算找出唯一的一个不重复元素
// 任何数和 0 做异或运算，结果仍然是原来的数，即 a⊕0=a。
// 任何数和其自身做异或运算，结果是 0，即 a⊕a=0。
// 异或运算满足交换律和结合律，即 a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。
func singleNumber2(nums []int) int {
	single := 0                //初始化用于存储结果的变量
	for _, num := range nums { //逐个的遍历
		single ^= num //通过异或操作筛选出唯一不重复的元素
	}
	return single //返回元素
}

// 利用异或运算求出数组
func main() {
	nums1 := []int{1}
	num1 := singleNumber1(nums1)
	fmt.Println(num1)
	nums2 := []int{2, 2, 1}
	num2 := singleNumber2(nums2)
	fmt.Println(num2)
}
