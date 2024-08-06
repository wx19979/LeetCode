package main

import "fmt"

//在haystack变量中找出needle多匹配第一个字符串的位置(暴力破解)
func strStr(haystack string, needle string) int {
	//首先获取两个字符串的长度
	n, m := len(haystack), len(needle)
	//创建跳跃点
outer:
	for i := 0; i+m <= n; i++ { //文本串循环操作
		for j := range needle { //模式串循环操作
			if haystack[i+j] != needle[j] { //如果当前出现不匹配的情况
				continue outer //直接进行外部循环
			}
		}
		return i //最终如果要是能完成循环直接返回对应位置的值
	}
	return -1 //没有的情况直接返回-1
}

//寻找next[j]的函数
func getNext(next []int, s string) {
	//首先初始化j和i
	j := -1
	next[0] = j
	//进行循环操作进行匹配计算
	for i := 1; i < len(s); i++ { //i需要从1开始进行比较
		//如果出现了不匹配的情况
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
}

//利用next[j]数组实现匹配操作
func strStr1(haystack string, needle string) int {
	//首先获取到两个数组的长度
	n, m := len(haystack), len(needle)
	//判断模式串是空串的情况
	if m == 0 {
		return 0
	}
	//开始创建next[j]数组
	nextj := make([]int, m)
	// 初始化变量用于记录匹配情况,j个数也表示匹配的个数
	j := -1
	//计算next[j]数组
	getNext(nextj, needle)
	//进行匹配操作
	for i := 0; i < n; i++ {
		//如果出现不匹配的状态
		for j >= 0 && haystack[i] != needle[j+1] {
			j = nextj[j] //找到匹配的位置
		}
		if haystack[i] == needle[j+1] { //匹配的情况
			j++ //j继续向上加
		}
		if j == m-1 { //如果当前j匹配了模式串长度说明全部匹配
			return i - m + 1 //直接返回匹配初始的位置
		}
	}
	return -1 //没配上就直接返回-1
}
func main() {
	haystack, needle := "leetcode", "leeto"
	i := strStr(haystack, needle)
	fmt.Println(i)
	haystack1, needle1 := "leetcode", "leeto"
	j := strStr1(haystack1, needle1)
	fmt.Println(j)
}
