package main

import "fmt"

//用于获取最大利润的函数
func maxProfit(prices []int) int {
	//转移方程
	//dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
	//dp[i][1]=max(dp[i-1][1],-prices[i])
	//判断是否是空的
	//获取长度
	pr_L := len(prices)
	if pr_L == 0 {
		return 0
	}
	//创建用于存储状态转移的数组
	dp := make([][]int, pr_L)
	//再循环创建每一个数组的状态有两种
	for i := 0; i < pr_L; i++ {
		dp[i] = make([]int, 2)
	}
	//初始化dp的前两个元素
	dp[0][0] = 0          //第一天什么都没买利润是0
	dp[0][1] = -prices[0] //第一个如果买了那利润是负的价格
	//循环计算利润
	for i := 1; i < pr_L; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) //第i天手里没有股票最大利润是第i天手头本身没有股票和第i天手里有股票但是卖掉了里面最大的值
		dp[i][1] = max(dp[i-1][1], -prices[i])           //第i天里手里有股票最大利润是第i天里面手里本身就有股票,和第i天里手里没有股票但是买了股票中的最大值
	}
	//最终返回最大的情况,肯定是手里面没有股票的情况利润最大
	return dp[pr_L-1][0]
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
	prices := []int{7, 6, 4, 3, 1}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
