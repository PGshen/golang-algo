package datastruct

type bSearchTreeNode struct {
	Value any
	Left  *bSearchTreeNode
	Right *bSearchTreeNode
}

type bSearchTree struct {
	Root *bSearchTreeNode
}

func newBSearchTreeNode(value any) *bSearchTreeNode {
	return &bSearchTreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func newBSearchTree() *bSearchTree {
	return &bSearchTree{
		Root: nil,
	}
}

// 搜索
func (b *bSearchTree) search(num int) *bSearchTreeNode {
	node := b.Root
	for node != nil {
		if node.Value.(int) < num {
			node = node.Right
		} else if node.Value.(int) > num {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

func (b *bSearchTree) insert(num int) {
	cur := b.Root
	if cur == nil {
		b.Root = newBSearchTreeNode(num)
		return
	}
	// 找到需要插入的位置
	var pre *bSearchTreeNode = nil
	for cur != nil {
		// 如果存在相同的值，那么不插入
		if cur.Value == num {
			return
		}
		pre = cur
		if cur.Value.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	// 插入节点
	node := newBSearchTreeNode(num)
	if pre.Value.(int) < num {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

func (b *bSearchTree) remove(num int) {
	cur := b.Root
	if cur == nil {
		return
	}

	// 找到要删除节点的前置节点
	var pre *bSearchTreeNode = nil
	for cur != nil {
		if cur.Value == num {
			break
		}
		pre = cur
		if cur.Value.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	// 要删除的节点不存在
	if cur == nil {
		return
	}

	// 分情况删除
	// 1. 要删除的节点没有子节点或者只有一个子节点
	if cur.Left == nil || cur.Right == nil {
		var child *bSearchTreeNode = nil
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		if cur == nil {
			// 是根节点
			b.Root = child
		} else {
			// 左还是右
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		}
	} else {
		// 2. 要删除的节点有2个子节点
		// 待删除节点的下一个节点
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除tmp
		b.remove(tmp.Value.(int))
		// 用tmp覆盖cur
		cur.Value = tmp.Value
	}
}

// 中序遍历
func (b *bSearchTree) inOrder() []any {
	res := []any{}
	return b.Root.dfs(res)
}

func (b *bSearchTreeNode) dfs(list []any) []any {
	if b == nil {
		return list
	}
	list = b.Left.dfs(list)
	list = append(list, b.Value)
	list = b.Right.dfs(list)
	return list
}
