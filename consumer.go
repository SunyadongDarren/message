package message

import (
	"context"
	"fmt"
	"sync"
)

func Consumer() {
	//step1 初始化工作协程
	workerNum := globalQueue.WorkNums
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go worker(context.Background(), globalQueue.Ch, &wg)
	}
	wg.Wait()

}

func worker(ctx context.Context, ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}
