// 给你一个整数数组prices,其中prices[i]表示某支股票第i天的价格。在每一天,你可以决定是否购买和/或出售股票。
// 你在任何时候最多只能持有一股股票。你也可以先购买,然后在同一天出售。返回你能获得的最大利润。
package main

import "fmt"

//获取最大值的函数
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//利用动态规划解出买卖股票的最佳时机
func maxProfit1(prices []int) int {
	n := len(prices)        //获取数组的长度
	dp := make([][2]int, n) //根据数组长度创建一个二维数组
	dp[0][1] = -prices[0]   //计算出如果一开始就买股票所获得的利润
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) //状态转移方程(当前是手头没有股票的状态所获得的最大的利润)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) //状态转移方程(当前是手头有股票的状态所获得的最大的利润)
	}
	return dp[n-1][0] //返回最后的结果
}

//利用动态规划解出买卖股票最佳时机(更加简单的方式)
func maxProfit2(prices []int) int {
	n := len(prices)          //获取长度
	dp0, dp1 := 0, -prices[0] //直接计算状态1和状态2的值
	for i := 1; i < n; i++ {  //循环计算转移方程的结果
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0 //返回结果
}

//利用贪心算法计算
func maxProfit3(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ { //将其划归为相邻的柱体那么想要有最大的获利,必须保持每次都要局部获利,最终就是最获利的情况
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}      //定义一个价格数组
	maxprices1 := maxProfit1(prices)       //调用获取最大利润的函数
	fmt.Println("maxProfit1:", maxprices1) //输出最终的结果
	maxprices2 := maxProfit2(prices)
	fmt.Println("maxProfit2:", maxprices2)
	maxprices3 := maxProfit3(prices)
	fmt.Println("maxProfit3:", maxprices3)
}
