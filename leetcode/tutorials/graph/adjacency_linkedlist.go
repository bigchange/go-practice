package graph

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 邻接表

// 顶点节点表示
type Vertex struct {
	Val int `json:"val"`
}

/* 基于邻接表实现的无向图类 */
type graphAdjList struct {
	// 邻接表，key: 顶点，value：该顶点的所有邻接顶点
	adjList map[Vertex][]Vertex
}

/* 构造方法 */
// edges：
// 第一维： 表示边的数量
// 第二维： 表示节点之间的关联
func newGraphAdjList(edges [][]Vertex) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	// 添加所有顶点和边
	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

/* 获取顶点数量 */
func (g *graphAdjList) size() int {
	return len(g.adjList)
}

/* 添加边 */
func (g *graphAdjList) addEdge(vet1 Vertex, vet2 Vertex) error {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	// 节点必须已存在，以及边不能是单个节点指向自身
	if !ok1 || !ok2 || vet1 == vet2 {
		return errors.New("edge not allowed")
	}
	// 添加边 vet1 - vet2
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
	return nil
}

/* 删除边 */
func (g *graphAdjList) removeEdge(vet1 Vertex, vet2 Vertex) error {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		return errors.New("edge not allowed")
	}
	// 删除边 vet1 - vet2
	g.DeleteSliceElms(vet1, vet2)
	g.DeleteSliceElms(vet2, vet1)
	return nil
}

/* 添加顶点 */
func (g *graphAdjList) addVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	// 在邻接表中添加一个新链表
	g.adjList[vet] = make([]Vertex, 0)
}

/* 删除顶点 */
func (g *graphAdjList) removeVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if !ok {
		panic("error")
	}
	// 在邻接表中删除顶点 vet 对应的链表
	delete(g.adjList, vet)
	// 遍历其他顶点的链表，删除所有包含 vet 的边
	for v := range g.adjList {
		g.DeleteSliceElms(v, vet)
	}
}

func (g *graphAdjList) DeleteSliceElms(vet1, vet2 Vertex) {
	i := 0
	list := g.adjList[vet1]
	for index, v := range list {
		//  找到需要删除的vet
		if v == vet2 {
			i = index
			break
		}
	}
	g.adjList[vet1] = append(g.adjList[vet1][:i], g.adjList[vet1][i+1:]...)
}

/* 打印邻接表 */
func (g *graphAdjList) print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for _, vet := range v {
			builder.WriteString(strconv.Itoa(vet.Val) + " ")
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}
