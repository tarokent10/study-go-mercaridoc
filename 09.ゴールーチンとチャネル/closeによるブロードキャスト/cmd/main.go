package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fire := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1) //ちなみにこのADDをgoroutineの中に入れるのはwg.wait()時点ではカウント0で待たずに終わるのでNG
		i := i    //この代入がないとＧｏルーチンはみなループ内で最後の値にアクセスする
		go func() {
			<-fire
			fmt.Printf("counter-%d\n", i)
			wg.Done()
		}()
	}
	time.Sleep(time.Second * 3)
	close(fire) // 1:nで親から子にブロードキャスト
	wg.Wait()

}
