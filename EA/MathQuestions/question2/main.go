package main

import "fmt"

//利用枚举算法进行计算(该算法超时)
func isPrime(x int) bool {
	//循环枚举判断是否是质数
	for i := 2; i*i <= x; i++ { //从2到根号x逐个去列举
		if x%i == 0 { //如果有能被整除的数说明其不是质数
			return false
		}
	}
	return true //循环结束都没找到其他因数说明是质数
}

// 计算小于n的所有质数
func countPrimes1(n int) int {
	//创建一个用于存储结果的变量
	cnt := 0
	for i := 2; i < n; i++ { //从2到n逐个判断是否是质数
		if isPrime(i) { //如果是进行累加
			cnt++
		}
	}
	return cnt //最终返回结果
}

// 埃氏筛法找到质数的个数
func countPrimes2(n int) (cnt int) {
	//创建用于存储是否是质数的数组
	isPrimes := make([]bool, n)
	//循环赋值为质数
	for i := range isPrimes {
		isPrimes[i] = true
	}
	//开始循环记录质数个数以及标记非质数
	for i := 2; i < n; i++ {
		if isPrimes[i] { //如果当前标记的是质数
			cnt++ //质数个数加一
			//然后将质数的倍数都标记为非质数
			for j := 2 * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}
	return
}
func main() {
	res1 := countPrimes1(10)
	fmt.Println(res1)
	res2 := countPrimes2(10)
	fmt.Println(res2)
}
