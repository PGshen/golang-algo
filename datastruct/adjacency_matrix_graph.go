package datastruct

import "fmt"

type adjMatGraph struct {
	// 顶点列表
	vertices []int
	// 矩阵
	adjMat [][]int
}

func newAdjMatGraph(vertices []int, edges [][]int) *adjMatGraph {
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	// 初始化图
	g := &adjMatGraph{
		vertices: vertices,
		adjMat:   adjMat,
	}
	// 添加边
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g
}

func (g *adjMatGraph) size() int {
	return len(g.vertices)
}

// 添加边
func (g *adjMatGraph) addEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index out of bounds exception")
	}
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

// 删除边
func (g *adjMatGraph) removeEdge(i, j int) {
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index out of bounds exception")
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

// 添加顶点
func (g *adjMatGraph) addVertex(val int) {
	n := g.size()
	g.vertices = append(g.vertices, val)
	// 添加一行
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	// 添加一列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

// 删除顶点
func (g *adjMatGraph) removeVertex(index int) {
	if index >= g.size() || index < 0 {
		fmt.Errorf("%s", "Index out of bounds exception")
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

// 广度优先遍历
func (g *adjMatGraph) bfs() []int {
	res := make([]int, 0)
	startIndex := -1
	for i := range g.vertices {
		for _, val := range g.adjMat[i] {
			if val == 1 { // 找到第一个点
				startIndex = i
				break
			}
		}
		if startIndex != -1 {
			break
		}
	}
	if startIndex == -1 {
		return res
	}
	// 记录是否已经访问过
	visited := make(map[int]struct{})
	visited[startIndex] = struct{}{}

	// 使用切片作为队列
	queue := make([]int, 0)
	queue = append(queue, startIndex)
	for len(queue) > 0 {
		vet := queue[0]
		queue = queue[1:]
		res = append(res, g.vertices[vet]) // 节点值

		for index, adjVet := range g.adjMat[vet] {
			if adjVet == 0 { // 0 说明这个边不存在
				continue
			}
			_, isVisited := visited[index]
			if !isVisited {
				queue = append(queue, index)
				visited[index] = struct{}{}
			}
		}
	}
	return res
}

// 深度优先遍历
func (g *adjMatGraph) dfs() []int {
	res := make([]int, 0)
	startIndex := -1
	for i := range g.vertices {
		for _, val := range g.adjMat[i] {
			if val == 1 { // 找到第一个点
				startIndex = i
				break
			}
		}
		if startIndex != -1 {
			break
		}
	}
	if startIndex == -1 {
		return res
	}
	visited := make(map[int]struct{})
	g.dfsHelper(visited, &res, startIndex)
	return res
}

func (g *adjMatGraph) dfsHelper(visited map[int]struct{}, res *[]int, idx int) {
	*res = append(*res, g.vertices[idx])
	visited[idx] = struct{}{}

	for index, adjVet := range g.adjMat[idx] {
		if adjVet == 0 { // 0 表示没有这条边
			continue
		}
		_, isVisited := visited[index]
		if !isVisited {
			g.dfsHelper(visited, res, index)
		}
	}
}

func (g *adjMatGraph) print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t邻接矩阵 = \n")
	for i := range g.adjMat {
		fmt.Printf("\t\t\t%v\n", g.adjMat[i])
	}
}
