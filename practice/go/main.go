package main

import (
	"fmt"
	"time"
	"sync"

	"github.com/zach-blizz/code-war-katas/practice/go/channels"

)

func main() {
	stop := make(chan bool, 2)
	wg := new(sync.WaitGroup)
	stop <- false

	runner := channels.NewRunner(stop, wg)
	fmt.Println("going to call run")
	go runner.Run()

	time.Sleep(50000)
	stop <- true

	wg.Wait()
	close(stop)
}