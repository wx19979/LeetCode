package main

import "fmt"

//方法一利用辅助数组,进行变换之后可以直接覆盖原始数组
func rotate1(matrix [][]int) {
	//首先获取到数组的行列长度
	m := len(matrix)
	//根据当前获取的长度创建一个二维数组
	copy_m := make([][]int, m)
	for i := 0; i < m; i++ {
		copy_m[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ { //逐层循环遍历二维数组
			copy_m[j][m-i-1] = matrix[i][j] //依次的将位置赋值到位
		}
	}
	copy(matrix, copy_m) //将变换后的图形进行复制即可
}

//方法二利用水平翻转和对角线对调来进行旋转
func rotate2(matrix [][]int) {
	n := len(matrix) //获取数组的长度
	//先进行水平翻转
	for i := 0; i < n/2; i++ { //可以单独对行进行操作即可
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	//再进行对角线的交换
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ { //对角线需要进行两层的操作
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotate3(matrix [][]int) {
	n := len(matrix) //获取长度
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ { //一半一半的交换,进行原地交换操作
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func main() {
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate1(matrix1)
	fmt.Println(matrix1)
	matrix2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate2(matrix2)
	fmt.Println(matrix2)
	matrix3 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate3(matrix3)
	fmt.Println(matrix3)
}
