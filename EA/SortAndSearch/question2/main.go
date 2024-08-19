package main

import (
	"fmt"
	"sort"
)

// 原题用于进行检测版本是否是正确的函数,true表示的是当前是错误的版本,false表示当前是正确的版本
func isBadVersion(version int) bool {
	if version == 8 { //设置错误开始时从第八个版本开始的
		return true
	} else {
		return false
	}
}

// 用于发现版本错误的版次函数(利用二分法)
func firstBadVersion1(n int) int {
	//创建用于搜索的左右指针
	left, right := 1, n
	for left != right { //开始循环进行比较操作
		mid := left + (right-left)/2 //计算中间的值,这种方式是防止溢出
		if isBadVersion(mid) {       //如果当前是错误的版本
			right = mid //说明错误版本在left和mid之间,需要向前寻找
		} else { //否则说明当前版本时正确的版本错误源头在后面
			left = mid + 1
		}
	}
	return left //当左右指针相遇说明找到了错误的版本,返回谁都可以
}

// 直接利用库函数直接搜索
func firstBadVersion2(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
func main() {
	bad1 := firstBadVersion1(10)
	fmt.Println(bad1)
	bad2 := firstBadVersion2(10)
	fmt.Println(bad2)
}
