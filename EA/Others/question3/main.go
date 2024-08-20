package main

import "fmt"

// 颠倒二进制位
func reverseBits(num uint32) (rev uint32) {
	//循环32次进行翻转操作
	for i := 0; i < 32 && num > 0; i++ {
		pn := 1 & num         //获取到最后一位的数字
		rev |= pn << (31 - i) //将最后第i位数字放置到正数第i位
		num >>= 1             //然后将当前数据向右移动保证每次循环都是最后一位
	}
	return //最终返回结果
}
func main() {
	i := reverseBits(43261596)
	fmt.Println(i)
}
