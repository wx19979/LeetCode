// 两个数组的交集二
package main

import "fmt"

// //用于查找两个数组的交集函数,利用哈希表
func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}                    //创建一个存储结果的数集
	len1, len2 := len(nums1), len(nums2) //获取两数组的长度
	longlen, shortlen := len1, len2      //先赋值长度(默认第一个数组长)
	longnum, shortnum := nums1, nums2    //再赋值数组
	if len1 < len2 {                     //如果长度是后面的常就反过来赋值
		longlen, shortlen = len2, len1
		longnum, shortnum = nums2, nums1
	}
	myhashtable := make(map[int]int) //创建哈希表
	for i := 0; i < shortlen; i++ {
		myhashtable[shortnum[i]] += 1 //对短的数组进行哈希计数
	}
	for i := 0; i < longlen; i++ {
		for key, value := range myhashtable { //用长的数组和哈希表中的内容逐一去做对比
			if key == longnum[i] && value > 0 { //如果出现相等的情况并且哈希表的值是大于零的情况
				result = append(result, key) //直接写入结果数组
				myhashtable[key]--           //修改哈希数组的值
			}
		}
	}
	return result //返回结果集
}

func main() {
	nums1, nums2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4} //创建用于测试的数组
	result := intersect(nums1, nums2)                    //调用测试
	fmt.Println(result)                                  //输出结果
}
