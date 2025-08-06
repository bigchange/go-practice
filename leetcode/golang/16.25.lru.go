package leetcode

import (
	"sort"
	"time"
)

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache struct {
	cap int
	m   map[int]ValueInfo
}
type ValueInfo struct {
	Counter       int
	LastVisitTime int64
	Value         int
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{cap: capacity, m: make(map[int]ValueInfo)}
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.m[key]
	if !ok {
		return -1
	}
	// todo()
	l.update(key, v)
	return v.Value
}

type ItemVaue struct {
	Key int
	V   ValueInfo
}

func (l *LRUCache) evcitItem() {
	var sortItems []ItemVaue
	for key, v := range l.m {
		sortItems = append(sortItems, ItemVaue{Key: key, V: v})

	}
	sort.Slice(sortItems, func(i, j int) bool {
		if sortItems[i].V.LastVisitTime == sortItems[j].V.LastVisitTime {
			return sortItems[i].V.Counter <= sortItems[j].V.Counter
		}
		return sortItems[i].V.LastVisitTime <= sortItems[j].V.LastVisitTime
	})
	evictKey := sortItems[0]
	delete(l.m, evictKey.Key)
}

func (l *LRUCache) update(key int, value ValueInfo) {
	value.Counter++
	newValue := ValueInfo{
		Counter:       value.Counter,
		LastVisitTime: time.Now().UnixNano(),
		Value:         value.Value,
	}
	l.m[key] = newValue
}

func (l *LRUCache) Put(key int, value int) {
	size := len(l.m)
	v, ok := l.m[key]
	if !ok {
		newValue := ValueInfo{
			Counter:       1,
			LastVisitTime: time.Now().UnixNano(),
			Value:         value,
		}
		if l.cap == size {
			l.evcitItem()
			l.m[key] = newValue
		} else {
			l.m[key] = newValue
		}
		return
	}
	v.Counter++
	v.LastVisitTime = time.Now().UnixNano()
	// put 相同key，可能是不同的value，需要更新对应value
	v.Value = value
	l.m[key] = v
}
