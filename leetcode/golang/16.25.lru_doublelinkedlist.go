package leetcode

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache_2 struct {
	cap int
	m   map[int]ValueInfo
}

func ConstructorLRU_2(capacity int) LRUCache {
	return LRUCache{cap: capacity, m: make(map[int]ValueInfo)}
}

func (l *LRUCache) Get(key int) int {

}

func (l *LRUCache) Put(key int, value int) {

}
