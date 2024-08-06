package main

import (
	"fmt"
	"math"
	"strings"
)

// 清理函数
func clean(s string) (sign int, abs string) {
	//首先去除首尾的空格
	s = strings.TrimSpace(s)
	//判断s是否是空串
	if s == "" {
		return
	}
	//判断s的第一个有效字符分情况讨论
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': //如果第一个字符是有效数字
		sign, abs = 1, s //直接放入后面即可
	case '+': //如果是正号
		sign, abs = 1, s[1:] //直接从后面进行截取,并将符号为赋值为正
	case '-': //如果是负号
		sign, abs = -1, s[1:] //直接从后面进行截取,并将符号为赋值为负
	default: //其他无意义的情况
		abs = "" //直接赋值为空串
		return
	}
	//判断后面的具体情况
	for i, b := range abs {
		if b < '0' || '9' < b { //如果后面的字符串出现不是数字的情况
			abs = abs[:i] //直接截取字符串
			break         //截断循环
		}
	}
	return //最终返回
}

// 进行组合操作的函数
func convert(sign int, absStr string) int {
	//首先创建一个用于存储最终结果的变量
	absNum := 0
	for _, b := range absStr { //遍历后面的字符进行操作
		absNum = absNum*10 + int(b-'0') //对位数进行累加操作
		switch {                        //判断溢出的情况
		case sign == 1 && absNum*sign > math.MaxInt32: //判断正溢出情况
			return math.MaxInt32 //返回范围内最大值
		case sign == -1 && absNum*sign < math.MinInt32: //判断负溢出情况
			return math.MinInt32 //返回范围内最小值
		}
	}
	return sign * absNum //返回最终的结果
}

func myAtoi(str string) int {
	return convert(clean(str)) //直接调用整理和结合函数
}

func main() {
	s := "   -042"
	i := myAtoi(s)
	fmt.Println(i)
}
