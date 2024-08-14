package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 对数进行层序遍历的函数
func levelOrder(root *TreeNode) [][]int {
	//首先创建用于存储最终结果的二维数组
	res := [][]int{}
	if root == nil { //判断是否是空树
		return res //空树直接返回空值
	}
	//创建一个以根节点为开头队列
	arr := []*TreeNode{root}
	//开始循环操作
	for len(arr) > 0 {
		//首先重置长度
		size := len(arr)
		//创建一个用于存储当前层的结果的数组
		curRes := []int{}
		//开始遍历队列
		for i := 0; i < size; i++ {
			//首先从队列取出当前遍历的元素
			node := arr[i]
			curRes = append(curRes, node.Val) //将当前遍历的值放到暂存结果的数组
			//如果左右子树都存在将左右子树也放到当前的队列中
			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		//将刚刚都遍历过的元素进行出队处理
		arr = arr[size:]
		//该层便利完将获得的暂存结果放到结果二维数组中
		res = append(res, curRes)
	}
	//最终返回结果
	return res
}

func main() {

}
