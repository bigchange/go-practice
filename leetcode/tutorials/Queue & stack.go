package tutorials

// Queue
// 1. support operations：
// Enqueue
// Dequeue
// Front
// Size
// Empty

// 2. implementations
// List: container/list
// Slice
import (
	"container/list"
	"fmt"
	"sync"
)

// List implementation
type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

// Slice implementation
type customQueueWithSlice struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customQueueWithSlice) Push(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customQueueWithSlice) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *customQueueWithSlice) Front() (string, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (c *customQueueWithSlice) Size() int {
	return len(c.stack)
}

func (c *customQueueWithSlice) Empty() bool {
	return len(c.stack) == 0
}

// Stack
// 1.support operations
// Push
// Pop
// Front
// Size
// Empty
// 2. implementations
// List: container/list
// Slice

// list implementation
type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}


// Slice implementation
type customStackWithSlice struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStackWithSlice) Push(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStackWithSlice) Pop() error {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
		return nil
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *customStackWithSlice) Front() (string, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len-1], nil
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (c *customStackWithSlice) Size() int {
	return len(c.stack)
}

func (c *customStackWithSlice) Empty() bool {
	return len(c.stack) == 0
}
