package main

import "fmt"

//判断两个单词是否是有效的字母异位词函数
func isAnagram(s string, t string) bool {
	m, n := len(s), len(t) //获取两个字符串的长度
	if m != n {            //如果两个字符串都不相等直接返回错误
		return false
	}
	s_t, t_t := [26]int{}, [26]int{} //创建用于存储结果的哈希表
	for i := 0; i < n; i++ {         //统计哈希表
		s_t[s[i]-'a']++
		t_t[t[i]-'a']++
	}
	//重新遍历两个哈希表进行比对
	for i := 0; i < 26; i++ {
		if s_t[i] != t_t[i] {
			return false
		}
	}
	return true
}
func main() {
	s, t := "rat", "car"
	isA := isAnagram(s, t)
	fmt.Println(isA)
}
