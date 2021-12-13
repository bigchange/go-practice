package leetcode

import (
	"container/heap"
	"sort"
)

/**
优先队列 + 贪心
思路：
在资本为: w 的情况下
每轮决策，在所有启动资本 capital[i] <= w 的项目中， 优先取 profits最大的项目即可
然后更新 w += max(profits[i])
直到取到指定的 k 个项目数，或者当前资本不足选取任何项目

优先队列： 可以使用最大堆实现

*/

type MaxHeap struct {
	Profits []int
}
type Item struct {
	capital int
	profit int
}

type ItemSorter []Item

func (s ItemSorter) Len() int { return len(s)}

func (s ItemSorter) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}
func (s ItemSorter) Less(i, j int) bool {
	return s[i].capital < s[j].capital
}

func (h *MaxHeap)  Len() int { return len(h.Profits)}

func (h *MaxHeap) Less(i, j int) bool {
	return h.Profits[i] > h.Profits[j]
}

func (h *MaxHeap) Swap(i, j int) {
	h.Profits[i], h.Profits[j] = h.Profits[j], h.Profits[i]
}

func (h *MaxHeap) Pop() interface{} {
	item := h.Profits[h.Len() - 1]
	h.Profits = h.Profits[:h.Len() - 1]
	return item
}
func (h *MaxHeap) Push(x interface{}) {
	h.Profits = append(h.Profits, x.(int))
}
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
  maxHeap := &MaxHeap{}
  n := len(capital)
	heap.Init(maxHeap)
  var items []Item
  for i:= 0; i < len(capital); i++ {
  	items = append(items, Item{capital: capital[i], profit: profits[i]})
	}
	// 每个项目只能加入一次到优先队列中
	// 首先可将项目数据按照启动资金capital升序排列
	sort.Sort(ItemSorter(items))
	i := 0
	// 前提： 需要选择k个项目
	for k > 0 {
		// 依次遍历capital列表，
		// 每轮在跟新w后，继续从剩下的项目中挑选满足：
		// 启动资金 <= w 的项目, 并加入优先队列中
		for i < n && items[i].capital <= w {
			heap.Push(maxHeap, items[i].profit)
			i++
		}
		// i 记录当前位置，第一个启动资本 > w 的项目
		// 提前终止条件： 如果队列为空， 资本w不足选择任何项目，退出
		if maxHeap.Len() == 0 {
			break
		}
		// 再优先队列中取出profit最大的项目，然后更新w
		// w 随着每次决策越来越大
		w += heap.Pop(maxHeap).(int)
		k--
	}
	return w
}