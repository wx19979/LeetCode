package main

import "fmt"

//找出0到n中没有出现过的数字
func missingNumber(nums []int) int {
	//创建一个用于存储结果的变量
	res := -1
	//首先获取到数组的长度
	n_len := len(nums)
	//根据这个长度创建一个哈希表
	m_nums := make([]bool, n_len+1)
	//初始化为false
	for i := 0; i < n_len+1; i++ {
		m_nums[i] = false
	}
	//开始遍历数组
	for i := 0; i < n_len; i++ {
		m_nums[nums[i]] = true //修改哈希表的值
	}
	//重新遍历哈希表
	for j := 0; j < n_len+1; j++ {
		if !m_nums[j] {
			res = j
			break
		}
	}
	return res
}
func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	num := missingNumber(nums)
	fmt.Println(num)
}
