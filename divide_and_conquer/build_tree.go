package divideandconquer

type TreeNode struct {
	Value any
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(value any) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

// 根据前序、中序遍历构建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := dfsBuildTree(preorder, inorder, 0, 0, len(inorder)-1)
	return root
}

// 根据后序、中序遍历构建二叉树
func buildTree2(postorder, inorder []int) *TreeNode {
	root := dfsBuildTree2(postorder, inorder, len(inorder)-1, 0, len(inorder)-1)
	return root
}

/**
 * @param preorder []int 前序遍历
 * @param inorder []int 中序遍历
 * @param i int 当前树根节点在preorder的索引
 * @param l int 当前树在inorder的左边界
 * @param r int 当前树在inorder的右边界
 */
func dfsBuildTree(preorder []int, inorder []int, i, l, r int) *TreeNode {
	if r-l < 0 {
		return nil
	}
	root := newTreeNode(preorder[i])
	// 找到当前节点在inorder的索引
	m := l
	for m <= r {
		if inorder[m] == preorder[i] {
			break
		}
		m++
	}
	root.Left = dfsBuildTree(preorder, inorder, i+1, l, m-1)
	root.Right = dfsBuildTree(preorder, inorder, i+1+m-l, m+1, r)
	return root
}

// 根据后序、中序遍历构建二叉树
func dfsBuildTree2(postorder, inorder []int, i, l, r int) *TreeNode {
	if r-l < 0 {
		return nil
	}
	root := newTreeNode(postorder[i])
	m := r
	for m >= l {
		if inorder[m] == postorder[i] {
			break
		}
		m--
	}
	root.Right = dfsBuildTree2(postorder, inorder, i-1, m+1, r)
	root.Left = dfsBuildTree2(postorder, inorder, i-1-(r-m), l, m-1)
	return root
}
