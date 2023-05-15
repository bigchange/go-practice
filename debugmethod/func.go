/*
 * @Author: Jerry You
 * @CreatedDate: 2023/4/21
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/4/21 11:44
 */
package main

import (
	"net/http"
	"sync"
)

var cli = http.DefaultClient

func fn(w *sync.WaitGroup) {
	defer w.Done()
	// do some things
	i := 0
	for ; i < 10000; {
		if i > 100 {
			continue
		}
		cli.Get("www.baidu.com")
		i++
	}

}
