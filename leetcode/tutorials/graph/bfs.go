package graph

// 图的: 广度优先搜索 bfs实现

// 广度优先遍历是一种由近及远的遍历方式，从距离最近的顶点开始访问，并一层层向外扩张

/* 广度优先遍历 BFS */
// 使用邻接表来表示图，以便获取指定顶点的所有邻接顶点
func graphBFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 顶点遍历序列
	res := make([]Vertex, 0)
	// 哈希表，用于记录已被访问过的顶点
	visited := make(map[Vertex]struct{})
	visited[startVet] = struct{}{}
	// 队列用于实现 BFS, 使用切片模拟队列
	queue := make([]Vertex, 0)
	queue = append(queue, startVet)
	// 以顶点 vet 为起点，循环直至访问完所有顶点
	for len(queue) > 0 {
		// 队首顶点出队
		vet := queue[0]
		queue = queue[1:]
		// 记录访问顶点
		res = append(res, vet)
		// 遍历该顶点的所有邻接顶点
		for _, adjVet := range g.adjList[vet] {
			_, isExist := visited[adjVet]
			// 只入队未访问的顶点
			if !isExist {
				queue = append(queue, adjVet)
				visited[adjVet] = struct{}{}
			}
		}
	}
	// 返回顶点遍历序列
	return res
}
