package questions

// 基于常规singleflight的增强
// 对于高并发写事件的合并处理

// 常规的singleflight是基于高并发的读事件处理

import (
	"sync"

	"golang.org/x/sync/singleflight"
)

type WriteGroup struct {
	mu    sync.Mutex
	wgs   map[string]*sync.WaitGroup
	group singleflight.Group
}

func NewWriteGroup() *WriteGroup {
	return &WriteGroup{}
}

func (g *WriteGroup) Do(key string, fn func() error) error {
	g.mu.Lock()
	if g.wgs == nil {
		g.wgs = make(map[string]*sync.WaitGroup)
	}
	wg, ok := g.wgs[key]
	if !ok {
		wg = &sync.WaitGroup{}
		wg.Add(1)
		g.wgs[key] = wg
	}
	g.mu.Unlock()

	if !ok {
		err := fn()

		g.mu.Lock()
		wg.Done()
		delete(g.wgs, key)
		g.mu.Unlock()
		return err
	}

	wg.Wait()
	_, err, _ := g.group.Do(key, func() (interface{}, error) {
		return nil, fn()
	})
	return err
}
