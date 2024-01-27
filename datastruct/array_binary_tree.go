package datastruct

type arrayBinaryTree struct {
	tree []any
}

func newArrayBinaryTree(arr []any) *arrayBinaryTree {
	return &arrayBinaryTree{
		tree: arr,
	}
}

func (a *arrayBinaryTree) size() int {
	return len(a.tree)
}

func (a *arrayBinaryTree) val(i int) any {
	if i < 0 || i >= a.size() {
		return nil
	}
	return a.tree[i]
}

func (a *arrayBinaryTree) left(i int) int {
	return 2*i + 1
}

func (a *arrayBinaryTree) right(i int) int {
	return 2*i + 2
}

func (a *arrayBinaryTree) parent(i int) int {
	return (i - 1) / 2
}

func (a *arrayBinaryTree) LevelOrder() []any {
	var res []any
	for i := 0; i < a.size(); i++ {
		if a.val(i) != nil {
			res = append(res, a.val(i))
		}
	}
	return res
}

func (a *arrayBinaryTree) dfs(i int, order string, res *[]any) {
	if a.val(i) == nil {
		return
	}
	if order == "pre" {
		*res = append(*res, a.val(i))
	}
	a.dfs(a.left(i), order, res)
	if order == "in" {
		*res = append(*res, a.val(i))
	}
	a.dfs(a.right(i), order, res)
	if order == "post" {
		*res = append(*res, a.val(i))
	}
}

func (a *arrayBinaryTree) PreOrder() []any {
	var res []any
	a.dfs(0, "pre", &res)
	return res
}

func (a *arrayBinaryTree) InOrder() []any {
	var res []any
	a.dfs(0, "in", &res)
	return res
}

func (a *arrayBinaryTree) PostOrder() []any {
	var res []any
	a.dfs(0, "post", &res)
	return res
}
