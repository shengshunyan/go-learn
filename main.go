package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
}

var wg sync.WaitGroup

func cpuMonitor(ctx context.Context) {
	fmt.Println("get traceId: ", ctx.Value("traceId"))

	for {
		select {
		case <-ctx.Done():
			fmt.Println("done!")
			wg.Done()
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("cpu info ...")
		}
	}
}

func main() {
	//var ch = make(chan bool)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	valueCtx := context.WithValue(ctx, "traceId", "4132")

	wg.Add(1)
	go cpuMonitor(valueCtx)

	wg.Wait()

}
