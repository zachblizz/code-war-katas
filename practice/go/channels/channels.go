package channels

import (
	"fmt"
	"sync"
	"time"

)

// Runner - runner type
type Runner struct {
	stop <-chan bool
	wg   *sync.WaitGroup
}

// NewRunner - returns a new runner
func NewRunner(stop <-chan bool, wg *sync.WaitGroup) *Runner {
	return &Runner{stop, wg}
}

// Run - runs the stuff...
func (r *Runner) Run() {
	fmt.Println("going to start...")

	for {
		select {
		case s := <-r.stop:
			fmt.Println(s)
			if s {
				r.wg.Done()
				fmt.Println("done...")
				return
			}
		default:
			time.Sleep(5 * time.Second)
			fmt.Println("running...")
		}
	}
}
