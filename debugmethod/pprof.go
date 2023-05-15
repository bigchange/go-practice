/*
 * @Author: Jerry You
 * @CreatedDate: 2023/4/19
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/4/19 18:07
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
)

func main() {

	// 设置GMP中P的数量（Processor数量）
	// P 中维护了该P下需要调度的goroutine队列
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 开启pprof功能
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	http.ListenAndServe(":8081", nil)
	wait := make(chan os.Signal)
	signal.Notify(wait)
	select {
	case s := <-wait:
		fmt.Printf("terminated by signal:%v", s.String())
	}
}
