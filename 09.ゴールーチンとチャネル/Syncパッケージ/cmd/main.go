package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// lock()
	// rwlock()
	wg()
}

func lock() {
	var m sync.Mutex
	m.Lock()
	go func() {
		time.Sleep(3 * time.Second)
		m.Unlock()
		fmt.Println("unlock1")
	}()
	// var m2 sync.Mutex // もちろんこれだと意味ない
	// m2.Lock()
	m.Lock()
	fmt.Println("この手前でブロック")
	m.Unlock()
	fmt.Println("unlock2")
}

func rwlock() {
	// RLock同士はブロックせず、Lockのみがブロックされる
	var m sync.RWMutex
	m.RLock()
	go func() {
		time.Sleep(3 * time.Second)
		m.RUnlock()
		fmt.Println("unlock1")
	}()
	m.RLock()
	m.RUnlock()
	fmt.Println("unlock2")
}

func wg() {
	var once sync.Once // 複数のGoルーチンから１回だけ呼ばれる関数を定義できる
	var wg sync.WaitGroup
	ps := []string{"tom", "jhon", "yam"}
	for _, p := range ps {
		go func(p string) {
			wg.Add(1)
			time.Sleep(2 * time.Second)
			fmt.Printf("my name is %s\n", p)
			once.Do(func() {
				fmt.Printf("i'm first man! (%s)\n", p)
			})
			defer wg.Done()
		}(p)
		time.Sleep(2 * time.Second)
	}
	wg.Wait()
	fmt.Println("done!")
}
