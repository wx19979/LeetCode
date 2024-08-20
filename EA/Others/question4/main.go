package main

import "fmt"

// 生成杨辉三角的函数
func generate(numRows int) [][]int {
	//创建用于存储最终结果的二维数组
	res := make([][]int, numRows)
	//开始循环操作生成杨辉三角
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1) //这里是第几行存个数比行数多一,所以申请空间要加一
		res[i][0] = 1             //杨辉三角每行第一个元素都是1
		res[i][i] = 1             //杨辉三角每一行最后一个元素也为1
		//开始循环进行状态转移方程
		for j := 1; j < i; j++ { //注意这里计算的值是从第二个开始算起所以j要从1开始
			res[i][j] = res[i-1][j] + res[i-1][j-1] //每一行中间的元素都是上一行中正对着的和后一个元素的和
		}
	}
	return res //最终返回结果
}
func main() {
	res := generate(10)
	fmt.Println(res)
}
