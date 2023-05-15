/*
 * @Author: Jerry You
 * @CreatedDate: 2023/4/19
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/4/19 18:07
 */

package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {

	// 设置GMP中P的数量（Processor数量）
	// P 中维护了该P下需要调度的goroutine队列
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, _ := os.Create("output/trace.output")
	defer f.Close()

	_ = trace.Start(f)
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fn(&wg)
	}
	wg.Wait()
}
