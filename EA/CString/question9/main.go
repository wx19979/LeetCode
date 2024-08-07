package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 外观数列(直接循环非递归)
func countAndSay1(n int) string {
	//首先创建一个用于存储前项结果的字符串并初始化为数列的第一项
	prev := "1"
	//然后开始循环操作构建外观数列
	for i := 2; i <= n; i++ {
		//首先创建一个用于存储当前循环下的数列项的字符串
		cur := &strings.Builder{}
		//开始循环构构建操作
		for j, start := 0, 0; j < len(prev); start = j {
			//开始统计prev[start]的个数
			for j < len(prev) && prev[j] == prev[start] {
				j++ //如果当前遍历的内容和prev[start]相同可以记录进去
			}
			cur.WriteString(strconv.Itoa(j - start)) //将统计好的数字写入cur中
			cur.WriteByte(prev[start])               //写入当前prev完成当前循环构建的操作
		}
		prev = cur.String() //整体构建循环结束将cur的外观数列项的值赋值给prev用于下次循环
	}
	return prev //最终返回获取到的值
}

// 外观数列函数(递归法)
func countAndSay2(n int) string {
	//首先设置递归的出口
	if n == 1 {
		return "1"
	}
	//然后递归调用得到它前一项的字符串
	s := countAndSay2(n - 1)
	//然后开始递归操作
	i, res := 0, "" //其中i是计数器,类似于慢指针,res是存储最终结果的字符串变量
	//开始处理该数列的前一项s字符串
	for j, c := range s { //需要逐项遍历字符串s,j属于快指针
		//需要计算当前遍历的位置s[i]前面到底有几个统计完之后加入到res字符串中
		if c != rune(s[i]) { //如果出现当前遍历的字符和前面慢指针所指的内容不一致的情况
			res += strconv.Itoa(j-i) + string(s[i]) //利用快指针减去慢指针计算s[i]在字符串中连续出现的次数,然后再加上s[i]组成外观数列
			i = j                                   //更新慢指针
		}
	}
	//整个循环结束之后还需要再把第n项加入到res中
	res += strconv.Itoa(len(s)-i) + string(s[i])
	return res //最终返回结果
}

func main() {
	s1 := countAndSay1(4)
	fmt.Println(s1)
	s2 := countAndSay2(4)
	fmt.Println(s2)
}
