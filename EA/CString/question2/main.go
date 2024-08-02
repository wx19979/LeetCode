package main

import "fmt"

//反转数字函数
func reverse(x int) int {
	res := 0     //创建一个用于存储最终结果的变量
	for x != 0 { //不断循环计算
		//每次都需要取尾数
		tmp := x % 10
		if res > 214748364 || (res == 214748364 && tmp > 7) { //如果超过最大数值的范围
			return 0
		}
		if res < -214748364 || (res == -214748364 && tmp < -8) { //如果超过最小数值的范围
			return 0
		}
		res = res*10 + tmp //将这一个位数值拼接在最终的结果上
		x /= 10            //x继续循环的除以10
	}
	return res //返回最终的结果
}

func main() {
	x1, x2, x3, x4 := 123, -123, 120, 0 //一共是4个数字
	x1 = reverse(x1)                    //调用反转函数
	x2 = reverse(x2)
	x3 = reverse(x3)
	x4 = reverse(x4)
	fmt.Println(x1, x2, x3, x4) //输出最终结果
}
