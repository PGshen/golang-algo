package datastruct

import "container/list"

type BTreeNode struct {
	Value any
	Left  *BTreeNode
	Right *BTreeNode
}

func newBTreeNode(value any) *BTreeNode {
	return &BTreeNode{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

// 层序遍历
func (b *BTreeNode) levelOrder() []any {
	root := b
	queue := list.New()
	queue.PushBack(root)
	ret := make([]any, 0)
	for queue.Len() > 0 {
		// 队首出队
		node := queue.Remove(queue.Front()).(*BTreeNode)
		ret = append(ret, node.Value)
		// 左子节点 入队
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		// 右子节点 入队
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return ret
}

func levelOrder(root *BTreeNode) [][]int {
	res := make([][]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		levelNodes := []BTreeNode{}
		for queue.Len() > 0 {
			node := queue.Remove(queue.Front()).(*BTreeNode)
			levelNodes = append(levelNodes, *node)
		}
		levelRes := []int{}
		for _, node := range levelNodes {
			levelRes = append(levelRes, node.Value.(int))
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, levelRes)
	}
	return res
}

// 前序遍历
func (b *BTreeNode) preOrder(list []any) []any {
	if b == nil {
		return list
	}
	list = append(list, b.Value)
	list = b.Left.preOrder(list)
	list = b.Right.preOrder(list)
	return list
}

// 中序遍历
func (b *BTreeNode) inOrder(list []any) []any {
	if b == nil {
		return list
	}
	list = b.Left.inOrder(list)
	list = append(list, b.Value)
	list = b.Right.inOrder(list)
	return list
}

// 后序遍历
func (b *BTreeNode) postOrder(list []any) []any {
	if b == nil {
		return list
	}
	list = b.Left.postOrder(list)
	list = b.Right.postOrder(list)
	list = append(list, b.Value)
	return list
}
