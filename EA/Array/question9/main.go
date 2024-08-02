package main

import "fmt"

//求两数和的数组函数
func twoSum(nums []int, target int) []int {
	// 创建一个用于存储结果的数组
	num := make([]int, 2)
	//依旧是创建双指针
	slow, n, fast := 0, len(nums), 1
	for slow <= n-2 {
		if n-1 < fast && slow != n-2 { //如果快指针到头了,直接定位到慢指针下一个位置
			slow++
			fast = slow + 1
		}
		if nums[slow]+nums[fast] == target { //如果当前的值符合条件直接返回结果
			num[0], num[1] = slow, fast
			break
		}
		fast++ //更新快指针
	}
	return num //最后返回结果
}

func main() {
	nums, target := []int{3, 2, 3}, 6
	tn := twoSum(nums, target)
	fmt.Println(tn)
}
