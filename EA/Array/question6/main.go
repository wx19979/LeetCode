// 两个数组的交集二
package main

import (
	"fmt"
	"sort"
)

// 用于查找两个数组的交集函数,利用哈希表
func intersect1(nums1 []int, nums2 []int) []int {
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

// 利用双指针排序方式查找并集
func intersect2(nums1 []int, nums2 []int) []int {
	result := []int{} //创建一个存储结果的数集
	//首先先将两个数组进行排序
	sort.Ints(nums1)
	sort.Ints(nums2)
	for pN1, pN2 := 0, 0; pN1 < len(nums1) && pN2 < len(nums2); { //创建指针指向两个数组的头部并进行循环
		if nums1[pN1] == nums2[pN2] { //如果出现相等的情况两指针同时向后挪
			result = append(result, nums1[pN1]) //相等的情况直接加入结果集
			pN1++                               //两个指针分别后移
			pN2++
		} else if nums1[pN1] < nums2[pN2] { //如果第一个数组值小于第二个数组值
			pN1++ //移动第一个数组的指针
		} else { //否则移动第二个数组的指针
			pN2++
		}
	}
	return result //最后返回结果集
}

func main() {
	nums1, nums2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4} //创建用于测试的数组
	result1 := intersect1(nums1, nums2)                  //调用测试
	fmt.Println(result1)                                 //输出结果
	nums3, nums4 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4} //创建用于测试的数组
	result2 := intersect2(nums3, nums4)                  //调用测试
	fmt.Println(result2)                                 //输出结果
}
