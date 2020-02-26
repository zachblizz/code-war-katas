package main

import (
	"fmt"
	"time"

	"github.com/zach-blizz/code-war-katas/practice/go/channels"

)

func main() {
	stop := make(chan bool, 2)
	stop <- false

	runner := channels.NewRunner(stop)
	fmt.Println("going to call run")
	go runner.Run()

	time.Sleep(50000)
	stop <- true
	close(stop)
}
