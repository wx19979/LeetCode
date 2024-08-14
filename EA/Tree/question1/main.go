package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据数组创建二叉树(左子树是2i+1,右子树是2i+2)
// func createBiTree(list []int) *TreeNode {

// }
// 利用递归的深度优先遍历搜索实现找到树的最大深度
func maxDepth1(root *TreeNode) int {
	//判断在空树的情况直接返回0
	if root == nil {
		return 0
	}
	//递归进行深度优先遍历
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1 //注意这里面没递归一层是需要加1的
}

// 首先创建一个比较大小的函数
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 利用队列的广度优先搜索实现找到树的最大深度
func maxDepth2(root *TreeNode) int {
	//首先判断树是否是空树
	if root == nil {
		return 0
	}
	//创建用于存储广度优先遍历的队列
	queue := []*TreeNode{}
	ans := 0 //创建用于存储最终结果的变量
	//将根节点放入队列中
	queue = append(queue, root)
	for len(queue) > 0 { //开始进行队列循环,队列完全为空循环结束
		sz := len(queue) //用于记录当前入队的元素个数
		for sz > 0 {     //开始对队列进行操作
			//首先取出当前遍历的节点
			node := queue[0]
			//然后对队列进行出队操作
			queue = queue[1:]
			//对计数器进行更新
			sz--
			//判断当前队列是否有子节点
			if node.Left != nil { //当前节点有左子树
				queue = append(queue, node.Left) //将左子树放到队列中
			}
			if node.Right != nil { //当前节点有右子树
				queue = append(queue, node.Right) //将右子树放入队列中
			}
		}
		ans++ //每次完成一层的操作计数器都加1记录层数
	}
	return ans //最终返回层数的结果
}

func main() {

}
