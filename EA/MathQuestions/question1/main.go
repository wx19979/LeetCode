package main

import (
	"fmt"
	"strconv"
)

// 如果i是3的倍数返回"Fizz",如果i是5的倍数返回"Buzz",如果i是3和5的倍数返回"FizzBuzz",否则就直接是i的字符串
func fizzBuzz(n int) []string {
	//创建用于存储结果字符串的变量
	res := []string{}
	//开始从1循环计算结果并放到结果字符串中
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 { //如果是3和5的倍数
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 { //如果只是3的倍数
			res = append(res, "Fizz")
		} else if i%5 == 0 { //如果只是5的倍数
			res = append(res, "Buzz")
		} else { //其他不满足的情况直接转为字符串
			res = append(res, strconv.Itoa(i))
		}
	}
	return res //最终返回结果
}
func main() {
	res := fizzBuzz(15)
	fmt.Println(res)
}
