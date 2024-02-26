package divideandconquer

import "container/list"

func solveHanota(A, B, C *list.List) {
	dfsHanota(A.Len(), A, B, C)
}

// 移动i个圆盘，将上边的i-1个作为一个整体，问题f(i)分解位f(i-1)和f(1)
func dfsHanota(i int, src, buf, tar *list.List) {
	if i == 1 {
		move(src, tar)
		return
	}
	dfsHanota(i-1, src, tar, buf)
	move(src, tar)
	dfsHanota(i-1, buf, src, tar)
}

func move(src, tar *list.List) {
	disk := src.Back()
	tar.PushBack(disk.Value)
	src.Remove(disk)
}
