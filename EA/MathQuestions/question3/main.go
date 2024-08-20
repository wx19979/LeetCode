package main

import "fmt"

//给一个数字判断该数字是否是3的幂次方
func isPowerOfThree(n int) bool {
	if n == 0 { //注意这里0不是3的幂次方
		return false
	} else if n == 1 { //但是1是3的幂次方因为3的零次幂等于1
		return true
	}
	for n%3 == 0 { //循环判断n是否能被3整除
		n = n / 3   //n不断的除以3
		if n == 1 { //如果除到最后n能够等于1,说明是3的幂次
			return true
		}
	}
	return false //不能被整除就不是3的幂次
}
func main() {
	isT := isPowerOfThree(26)
	fmt.Println(isT)
}
