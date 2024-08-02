package main

import "fmt"

//方法一使用额外数组进行翻转操作
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums)) //创建一个用于存储最终结果的数组
	for i, v := range nums {          //逐个遍历
		newNums[(i+k)%len(nums)] = v //逐个赋值
	}
	copy(nums, newNums) //修改完之后拷贝过去
}

//方法二利用环状替代法进行翻转操作

//求a和b的最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//具体的翻转函数
func rotate2(nums []int, k int) {
	n := len(nums) //获取到数组的长度
	k %= n         //计算k值
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start          //初始化当前的数据,分别是前向的数据,以及当前指针所在的下标
		for ok := true; ok; ok = cur != start { //不断循环的更改直到遇到了最初的起点才结束
			next := (cur + k) % n                        //计算当前位置的下一个应该跳转的位置
			nums[next], pre, cur = pre, nums[next], next //交换位置,并且更新前向指针和下一跳的指针
		}
	}
}

// 方法三数组翻转
// 反转函数
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ { //按照一半一半的进行交换操作
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

//具体根据给出的k进行翻转
func rotate3(nums []int, k int) {
	k %= len(nums)    //计算多少个元素会被翻转到数组头部 k mod n个元素
	reverse(nums)     //整体翻转
	reverse(nums[:k]) //根据计算的k值进行部分翻转
	reverse(nums[k:])
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7} //创建一个数组
	fmt.Println("former:", nums1)
	rotate1(nums1, 3) //调用翻转函数
	fmt.Println("rotate:", nums1)
	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("former:", nums2)
	rotate2(nums2, 3)
	fmt.Println("rotate:", nums2)
	nums3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("former:", nums3)
	rotate3(nums3, 3)
	fmt.Println("former:", nums3)
}
