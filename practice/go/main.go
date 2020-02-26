package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/zach-blizz/code-war-katas/practice/go/channels"

)

func main() {
	stop := make(chan bool, 2)
	wg := new(sync.WaitGroup)
	stop <- false

	runner := channels.NewRunner(stop, wg)
	fmt.Println("going to call run")
	wg.Add(1)
	go runner.Run()

	go func() {
		time.Sleep(50 * time.Second)
		fmt.Println("going to stop now....")
		stop <- true
	}()

	wg.Wait()
	close(stop)
}
