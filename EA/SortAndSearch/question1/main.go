package main

import (
	"fmt"
	"sort"
)

// 用于给两个数组进行组合排序的函数,最终结果放到第一个数组中
func merge1(nums1 []int, m int, nums2 []int, n int) {
	//创建用于进行操作的两个指针
	p1, p2 := 0, 0
	//创建用于存储最终结果的数组
	sorted := []int{}
	//循环进行添加操作
	for {
		if p1 == m { //当第一个数组都放到结果数组中
			sorted = append(sorted, nums2[p2:]...) //直接将第二个数组后面的接到结果数组中
			break                                  //结束循环
		}
		if p2 == n { //否则就将第一个数组剩余元素接到结果数组中
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] { //如果第一个数组的元素小于第二个数组的元素
			sorted = append(sorted, nums1[p1]) //将第一个数组的元素放到数组中
			p1++                               //移动第一个数组的指针
		} else { //否则反向操作
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted) //最终将结果放到第一个数组中
}

// 第二种方式,直接将nums2数组接入到后面,然后再调用排序函数
func merge2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) //将nums2拷贝到nums1后面
	sort.Ints(nums1)       //调用排序函数
}
func main() {
	num1, num2 := []int{0}, []int{1}
	merge1(num1, 0, num2, 1)
	fmt.Println(num1)
	num3, num4 := []int{0}, []int{1}
	merge2(num3, 0, num4, 1)
	fmt.Println(num3)
}
