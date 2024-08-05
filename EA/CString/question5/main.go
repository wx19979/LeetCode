package main

import (
	"fmt"
	"strings"
)

// 判断回文字符串
func isPalindrome(s string) bool {
	s1 := ""                      //创建一个用于存储去除无意义的字符串
	for i := 0; i < len(s); i++ { //逐个遍历字符串
		if isalnum(s[i]) { //如果当前的字符是有意义的字符
			s1 += string(s[i]) //将当前字符加入字符串中
		}
	}
	//然后统一转换成统一为小写
	s2 := strings.ToLower(s1)
	//获取到处理完的字符串的长度
	n := len(s2)               //获取到处理好的字符串长度
	for i := 0; i < n/2; i++ { //利用双指针方式进行判断是否是回文字符串
		if s2[i] != s2[n-i-1] { //头指针尾指针出现不相等的情况
			return false //直接返回错误
		}
	}
	return true //整体遍历完之后没错直接返回是
}

// 用于判断当前的内容是否是一个合法的字符
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') //返回当前的字符是否是在有意义字符的范围之中
}
func main() {
	s := "A man, a plan, a canal: Panama"
	isP := isPalindrome(s)
	fmt.Println(isP)
}
