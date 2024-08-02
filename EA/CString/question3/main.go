package main

import "fmt"

func firstUniqChar(s string) int {
	//首先创建一个用于存储个数的哈希表
	cnt := [26]int{}
	//开始逐个的遍历并进行统计
	for _, ch := range s {
		cnt[ch-'a']++
	}
	//再重新遍历一遍找到第一个是只有一个字母的位置
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1 //如果没有的情况下直接返回-1
}

func main() {
	s := "loveleetcode"
	i := firstUniqChar(s)
	fmt.Println(i)
}
