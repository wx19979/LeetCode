package main

import "fmt"

//反转字符串函数(直接交换位置)
func reverseString1(s []byte) {
	//首先获取字节数组长度
	n := len(s)
	for i := 0; i < n/2; i++ { //对前一个和后一个进行调换
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

//利用双指针解决反转问题
func reverseString2(s []byte) {
	left, right := 0, len(s)-1 //初始化左右双指针
	for left < right {         //循环进行翻转操作
		s[left], s[right] = s[right], s[left] //交换两指针所指位置的内容
		left++                                //更新左指针
		right--                               //更新右指针
	}
}

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'} //创建一个字节数组
	reverseString1(s1)                    //调用翻转函数
	fmt.Println(s1)                       //输出反转的结果
	s2 := []byte{'h', 'e', 'l', 'l', 'o'} //创建一个字节数组
	reverseString2(s2)                    //调用翻转函数
	fmt.Println(s2)                       //输出反转的结果
}
