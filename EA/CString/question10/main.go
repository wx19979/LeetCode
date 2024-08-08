package main

import (
	"fmt"
)

// 求解字符串数组的最长前缀(横向比较法)
func longestCommonPrefix1(strs []string) string {
	//获取到字符串数组的长度
	count := len(strs)
	//判断当前字符串数组是否是空数组
	if count == 0 {
		return "" //空的直接返回空串
	}
	//创建一个用于存储最长前缀的字符串变量
	prefix := strs[0]
	//开始循环的横向比较操作
	for i := 1; i < count; i++ { //注意这里要从第二个开始进行比较操作
		prefix = lcp(prefix, strs[i]) //将新获得的最长前缀进行赋值
		if len(prefix) == 0 {         //如果当前前缀长度是0
			break //可以不用找了直接停止循环操作
		}
	}
	return prefix //循环完之后直接返回结果
}

// 用于求解两个字符串之间的最长前缀的函数
func lcp(str1, str2 string) string {
	//首先设置用于进行对比遍历操作的指针,并且获取到遍历的最大长度
	index, count := 0, min(len(str1), len(str2))
	for index < count && str1[index] == str2[index] { //循环当两个字符串都是相等的情况就挪动指针
		index++ //挪动指针
	}
	return str1[:index] //返回最终的最长前缀
}

// 用于求解两数中最小的数字的函数
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 求解字符串数组的最长前缀(纵向比较法)
func longestCommonPrefix2(strs []string) string {
	//首先获取到字符串数组的长度
	count := len(strs)
	//判断是否是空数组情况
	if count == 0 {
		return "" //空数组直接返回空串
	}
	//按照str[0]开始进行纵向对比
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < count; j++ {
			//不断循环直到发现纵向对比到某一个字符串的最长的位置,或者出现纵向比较不同的情况结束循环
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i] //直接返回最终的匹配结果
			}
		}
	}
	//如果整个循环能跑完说明当前strs[0]就是最长的前缀直接返回strs[0]
	return strs[0]
}

// 求解字符串数组的最长前缀(分治法)
func longestCommonPrefix3(strs []string) string {
	//首先获取到字符串数组的长度
	count := len(strs)
	//判断是空数组的情况
	if count == 0 {
		return "" //空数组直接返回空
	}
	//定义一个分治法的函数对象方便后面进行递归操作
	var lcp func(int, int) string
	lcp = func(start, end int) string { //具体定义该函数
		if start == end { //当比较的数组相遇的时候直接返回一个即可
			return strs[start]
		}
		//首先计算出start和end的中间值
		mid := (start + end) / 2
		//开始递归调用分治法计算两侧的最长前缀
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		//求出左右两侧的最长前缀的最短长度
		minLength := min(len(lcpLeft), len(lcpRight))
		//开始对左右两侧进行比较求出左右两侧的最长前缀
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] { //当出现不相等的情况说明已经匹配到最长的前缀了
				return lcpLeft[:i] //直接返回当前的结果
			}
		}
		//如果一直都是匹配的说明他们前缀都是相等随便返回一个即可
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1) //递归调用分治法函数
}

// 求解字符串数组的最长前缀(二分查找法)
func longestCommonPrefix4(strs []string) string {
	//首先判断数组是否是空数组
	count := len(strs)
	if count == 0 { //如果是空数组返回空串
		return ""
	}
	//定义一个用于判断当前的长度前缀是否是公共的前缀的函数
	isCommonProfix := func(mlength int) bool {
		for i := 1; i < count; i++ { //注意这里要从1开始因为都是在和str[0]做比较
			if strs[i][:mlength] != strs[0][:mlength] { //逐个按照这个长度截取字符串和str[0]截取的字符串进行比较
				return false //如果出现不同直接返回错误
			}
		}
		return true //循环完都没有不同说明就是公共的前缀字符串
	}
	//初始化当前的长度
	minLength := len(strs[0])
	//然后逐个对比找到最短的长度
	for i := 1; i < count; i++ {
		len_i := len(strs[i])  //获取当前的长度
		if len_i < minLength { //如果小于最小的长度
			minLength = len_i //直接进行赋值
		}
	}
	low, high := 0, minLength //初始化二分法的两个指针
	for low < high {          //开始进行二分法的循环
		midlen := (high-low+1)/2 + low //以后计算中值直接用这个防止溢出现象
		if isCommonProfix(midlen) {    //如果当前的截取的字符串是公共前缀
			low = midlen //说明公共前缀的长度大于midlen值那么向上二分查找
		} else { //否则说明公共前缀的长度小于midlen的值那么向下二分查找
			high = midlen - 1
		}
	}
	return strs[0][:low] //最终返回结果
}
func main() {
	strs1 := []string{"flower", "flow", "flight"}
	str1 := longestCommonPrefix1(strs1)
	fmt.Println(str1)
	strs2 := []string{"flower", "flow", "flight"}
	str2 := longestCommonPrefix2(strs2)
	fmt.Println(str2)
	strs3 := []string{"flower", "flow", "flight"}
	str3 := longestCommonPrefix3(strs3)
	fmt.Println(str3)
	strs4 := []string{"flower", "flow", "flight"}
	str4 := longestCommonPrefix4(strs4)
	fmt.Println(str4)
}
