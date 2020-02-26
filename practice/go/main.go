package main

import "github.com/zach-blizz/code-war-katas/practice/go/channels"

func main() {
	stop := make(chan bool)
	stop <- false

	runner := NewRunner(&stop)

}
