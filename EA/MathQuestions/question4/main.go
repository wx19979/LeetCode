package main

import "fmt"

//罗马数字转整数的函数
func romanToInt(s string) int {
	len_s := len(s)
	//创建用于存储数值的变量
	res := 0
	for i := len_s - 1; i >= 0; i-- { //需要将数组从右向左依次扫描
		switch s[i] { //判断第i个位置的数据
		case 'I': //代表是1
			res = 1                         //看到I说明已经有一个I了所以初始化为1
			for i-1 >= 0 && s[i-1] == 'I' { //循环向后找是否带有I,并且移动的位置要合法
				res += 1 //累加操作
				i--      //指针向后移动
			}
			if i-1 >= 0 && s[i-1] == 'V' { //如果遇到5
				res += 5 //这种情况就直接加上5即可
				i--      //指针继续向后移动
			}
		case 'V': //代表是5
			if i-1 >= 0 && s[i-1] == 'I' { //如果遇到左面有1那么需要剪掉1,变为4
				res += 4 //进行累加操作
				i--      //指针向后移动
			} else { //否则什么都没有直接就加上原有的数值
				res += 5
			}
		case 'X': //代表是10
			if i-1 >= 0 && s[i-1] == 'I' { //以下均同理
				res += 9
				i--
			} else {
				res += 10
			}
		case 'L': //代表是50
			if i-1 >= 0 && s[i-1] == 'X' {
				res += 40
				i--
			} else {
				res += 50
			}
		case 'C': //代表是100
			if i-1 >= 0 && s[i-1] == 'X' {
				res += 90
				i--
			} else {
				res += 100
			}
		case 'D': //代表是500
			if i-1 >= 0 && s[i-1] == 'C' {
				res += 400
				i--
			} else {
				res += 500
			}
		case 'M': //代表是1000
			if i-1 >= 0 && s[i-1] == 'C' {
				res += 900
				i--
			} else {
				res += 1000
			}
		}
	}
	return res //最终累加结束返回值即可
}
func main() {
	s := "MCMXCIV"
	num := romanToInt(s)
	fmt.Println(num)
}
