package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用于查看是否是对称的树的函数,利用递归的方式进行检查
func isSymmetric1(root *TreeNode) bool {
	return check(root, root)
}

// 用于具体的检查操作
func check(p, q *TreeNode) bool {
	if p == nil && q == nil { //如果当前都是空的说明是对称的
		return true
	}
	if p == nil || q == nil { //只有一个有节点说明不对称
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left) //递归判断两个值是否是相等
}

// 用于查看是否是对称的树的函数,利用迭代法和队列进行求解
func isSymmetric2(root *TreeNode) bool {
	//创建用于存储对比左右子树的节点变量
	u, v := root, root
	//创建用于进行遍历操作的队列
	q := []*TreeNode{}
	q = append(q, u) //将左右都加入队列中
	q = append(q, v)
	for len(q) > 0 { //开始循环遍历
		u, v = q[0], q[1] //从队列中取出两元素
		q = q[2:]         //进行出队操作
		//进行判断操作
		if u == nil && v == nil { //如果两个节点都是空的就继续向下遍历进行对比
			continue
		}
		if u == nil || v == nil { //如果其中有一个是空的说明不对称直接返回错误
			return false
		}
		if u.Val != v.Val { //如果两个值不同说明也是不对称的
			return false
		}
		//将树的节点按照应该对称的方式加入到队列中以使得能够继续进行判断树是否对称
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true //整体比对都没有错误直接返回正确信息
}

func main() {

}
