package main

import (
	"fmt"
	"math"
	"strings"
)

// 清理函数
func clean(s string) (sign int, abs string) {
	//先去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" { //如果当前的字符串是空串直接返回
		return
	}
	//判断第一个字符
	switch s[0] {
	//有效的
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	//有效的,正号
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
		//无效的当空字符处理,并且直接返回
	default:
		abs = ""
		return
	}
	for i, b := range abs { //继续判断剩下的字符是否在范围内
		if b < '0' || '9' < 0 { //如果剩下的内容不是数字
			abs = abs[:i] //直接把前面是有效的数字进行剪切
			break         //剪切完直接退出循环
		}
	}
	return //最终什么都不满足的情况直接返回
}

// 进行组合操作的函数
func convert(sign int, absStr string) int {
	absNum := 0                // 首先初始化待返回数值
	for _, b := range absStr { //遍历后面的数值
		absNum = absNum*10 + int(b-'0') //每一位可以用来进行进位
		switch {                        //根据情况进行区分
		case sign == 1 && absNum > math.MaxInt32: //如果当前数值正向溢出
			return math.MaxInt32 //直接返回最大值
		case sign == -1 && absNum*sign < math.MinInt32: //如果当前数值负向溢出
			return math.MinInt32 //直接返回最小值
		}
	}
	return sign * absNum //最终返回最终的值
}

func myAtoi(str string) int {
	return convert(clean(str)) //直接返回最终的结果
}

func main() {
	s := "   -042"
	i := myAtoi(s)
	fmt.Println(i)
}
