package main

import "fmt"

//dp[i]=max(nums[i]+f[i-1],nums[i])
//找出最大子序和的函数
func maxSubArray1(nums []int) int {
	//创建一个用于存储最大值的变量
	max := nums[0]
	//从第二个数字开始向后遍历
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+nums[i-1] { //如果出现结尾数字加上前面一个数字更大
			nums[i] += nums[i-1] //那么就直接加上去
		}
		if nums[i] > max { //如果当前加的值比最大值还要大
			max = nums[i] //那么更新最大值
		}
	}
	return max
}

//找出最大子序和的函数(状态转移更明显的写法)
func maxSubArray2(nums []int) int {
	len_n := len(nums)       //获取数组的长度
	dp := make([]int, len_n) //根据长度创建dp数组
	//设置好dp的初始值
	dp[0] = nums[0]
	//创建一个存储最大值的变量
	max_n := dp[0]
	//开始循环进行操作
	for i := 1; i < len_n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if max_n < dp[i] {
			max_n = dp[i]
		}
	}
	return max_n
}

//创建用于比较大小的函数
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxN1 := maxSubArray1(nums1)
	fmt.Println(maxN1)
	nums2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxN2 := maxSubArray2(nums2)
	fmt.Println(maxN2)
}
