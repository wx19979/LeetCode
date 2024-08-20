package main

import (
	"fmt"
	"math/bits"
)

// 计算海明码距离的函数(直接利用内置函数进行求解)
func hammingDistance1(x int, y int) int {
	return bits.OnesCount(uint(x ^ y)) //bits.OnesCount返回的是1的个数,将x和y异或的结果转换为无符号整型
}

// 计算海明码距离的函数(移位实现计数)
func hammingDistance2(x int, y int) (ans int) {
	t := x ^ y
	fmt.Println(t)
	//创建一个用于存储结果的变量
	for s := x ^ y; s > 0; s >>= 1 { //s是x和y的异或的结果,将s不断向右移动直到s等于0
		ans += s & 1 //判断s最右侧那位是否是1如果是1就累加操作
	}
	return
}
func main() {
	res1 := hammingDistance1(1, 4)
	fmt.Println(res1)
	res2 := hammingDistance2(1, 4)
	fmt.Println(res2)
}
