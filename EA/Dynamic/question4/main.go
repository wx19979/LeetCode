package main

import (
	"fmt"
)

// 打家劫舍函数,首先要明确每家有两个状态一种是被偷,
// 一种是没有被偷, 并且被偷的两家不相邻那么应该有两种
// 情况一种是第i家被偷了,那么第i-1家不能被偷。还有一种
// 是第i家没有被偷,那么第i-1家可以被偷也可以不被偷
// 即dp[i][0]=max(dp[i-1][0],dp[i-1][1])
// dp[i][1]=dp[i-1][0]+nums[i]
func rob(nums []int) int {
	//首先获取数组的长度
	len_n := len(nums)
	//根据数组的长度创建相应的状态转移二维数组
	dp := make([][]int, len_n)
	//然后再在该数组上初始化具有两种状态
	for i := 0; i < len_n; i++ {
		dp[i] = make([]int, 2)
	}
	//将刚开始的两种状态进行初始化,即第一家被偷的状态,和第一家没有被偷的状态
	dp[0][0] = 0       //第一家没有被偷
	dp[0][1] = nums[0] //第一家被偷获得的钱
	//开始循环计算状态转移方程
	for i := 1; i < len_n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]) //计算第i家没有被偷的最大情况
		dp[i][1] = dp[i-1][0] + nums[i]        //计算第i家被偷而第i-1家没有被偷获得的钱
	}
	return max(dp[len_n-1][0], dp[len_n-1][1]) //返回两种情况的最大值
}

// 比较大小的函数
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	maxr := rob(nums)
	fmt.Println(maxr)
}
