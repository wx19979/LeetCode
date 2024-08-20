package main

import "fmt"

//求位一的个数即将该数转为二进制数之后上面是1的个数
func hammingWeight(n int) int {
	res := 0     //用于记录结果
	for n != 0 { //当n不为0的时候开始循环计算
		if n%2 == 1 { //如果和2取余操作为1
			res++ //那么就记录一下
		}
		n = n / 2 //然后n不停的除2
	}
	return res //最终循环结束就可以统计出一共取余操作了多少个1
}
func main() {
	res := hammingWeight(2147483645)
	fmt.Println(res)
}
