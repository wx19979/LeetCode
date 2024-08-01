package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	//首先创建用于记录的三个哈希表
	rows, columns, subboxs := [9][9]int{}, [9][9]int{}, [3][3][9]int{}
	//开始遍历二维数组
	for i, row := range board {
		for j, c := range row {
			if c == '.' { //如果当前是'.'字符表示当前位置是空的直接跳过循环
				continue
			}
			index := c - '1'                                                                 //因为数组是从0开始的因此减掉一个1也是从0开始
			rows[i][index]++                                                                 //记录行哈希表
			columns[j][index]++                                                              //记录列哈希表
			subboxs[i/3][j/3][index]++                                                       //记录3*3小块的哈希表
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxs[i/3][j/3][index] > 1 { //逐一对比条件
				return false
			}
		}
	}
	return true
}

func main() {
	//创建测试的案例
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isValid := isValidSudoku(board) //调用测试函数
	fmt.Println(isValid)            //输出测试的结果
}
