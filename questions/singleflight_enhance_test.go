package questions

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"sync/atomic"
)

func TestWriteGroup_Do(t *testing.T) {
	type args struct {
		key string
		fn  func() error
	}
	writeGroup := NewWriteGroup()
	var result atomic.Value
	events := make(chan int, 100)

	go func() {
		for i := 0; i < 100; i++ {
			events <- i
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for val := range events {
				getEvent := val
				writeGroup.Do("recv_msg", func() error {
					result.Store(getEvent)
					return nil
				})
			}
		}()
	}

	fmt.Println("gorpoutines:", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
	close(events)
	fmt.Printf("result:%v\n", result.Load())
}

func BenchmarkWriteGroup(b *testing.B) {
	writeGroup := &WriteGroup{}
	var count atomic.Int32
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			writeGroup.Do("register_event_msg", func() error {
				count.Add(1)
				d := time.Duration(rand.Intn(10))
				time.Sleep(d * time.Millisecond)
				return nil
			})
		}
	})
	b.Logf("count:%v", count.Load())
}
