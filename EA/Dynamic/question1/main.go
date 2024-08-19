package main

import "fmt"

//爬楼梯函数 dp[n]=dp[n-1]+dp[n-2](利用数组存储状态)
func climbStairs1(n int) int {
	//判断如果时1或者2直接返回结果
	if n == 1 || n == 2 {
		return n
	}
	//创建一个用于存储结果的数组
	dp := make([]int, n)
	dp[0] = 1 //初始化前面两个
	dp[1] = 2
	for i := 2; i < n; i++ { //循环计算后面的值并将结果存储到数组中
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1] //最终返回结果
}

//减少空间的利用(滚动数组)
func climbStairs2(n int) int {
	//如果是前两项直接返回结果
	if n == 1 || n == 2 {
		return n
	}
	prePre := 1 //直接记录其中的两个状态来实现,大大的节省了空间
	pre := 2
	for i := 2; i < n; i++ { //循环实现状态转移方程
		pre += prePre
		prePre = pre - prePre
	}
	return pre
}

//记忆化递归计算
func climbStairs3(n int) int {
	memo := make([]int, n+1)
	return climbStairsMemo(n, memo)
}
func climbStairsMemo(n int, memo []int) int {
	//判断当前的memo数组是否是已经被赋值,如果是直接从数组中取出结果
	if memo[n] > 0 {
		return memo[n] //直接取出来返回结果
	}
	if n == 1 || n == 2 { //如果为1或者2直接等于其本身并且进行赋值即可
		memo[n] = n //记录当前的状态
	} else { //其他情况就递归利用状态转移方程
		memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo) //状态转移方程并记录当前的状态
	}
	//最后共返回结果
	return memo[n]
}
func main() {
	dp1 := climbStairs1(45)
	fmt.Println(dp1)
	dp2 := climbStairs2(45)
	fmt.Println(dp2)
	dp3 := climbStairs3(45)
	fmt.Println(dp3)
}
