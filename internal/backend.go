package internal

import (
	"fmt"
	"sync"
)

// 变量
var (
	__backendOnce sync.Once
	__backend     *Backend
)

type Backend struct {
	pool  chan func() // 一个管道的池子
	debug bool        // 是否启动调试
}

// 实例化对象
// 不管实例多少次，对象就创建一次。
func NewBackend(debug ...bool) *Backend {
	__backendOnce.Do(func() {
		var bug bool
		if len(debug) > 0 {
			bug = debug[0]
		}
		__backend = &Backend{
			debug: bug,
		}
		__backend.pool = make(chan func(), 10) // 初使 pool
		// 开启一个 goroutine
		go __backend.monitor()
	})
	return __backend
}

// 添加服务
func (b *Backend) Add(fn func()) {
	if b.debug {
		fmt.Println("add +1")
	}
	b.pool <- fn
}

// 监听服务
func (b *Backend) monitor() {
	for {
		select {
		case fn, ok := <-b.pool:
			if b.debug {
				fmt.Printf("ok:%v", ok)
			}
			if ok {
				fn()
			}
		}
	}
}
