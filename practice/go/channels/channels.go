package channels

import (
	"fmt"
	"sync"
	"time"

)

// Runner - runner type
type Runner struct {
	stop chan bool
	wg   *sync.WaitGroup
}

// NewRunner - returns a new runner
func NewRunner(stop chan bool, wg *sync.WaitGroup) *Runner {
	return &Runner{stop, wg}
}

// Run - runs the stuff...
func (r *Runner) Run() {
	for {
		fmt.Println("going to start...")
		r.wg.Add(1)
		s := <-r.stop

		if s {
			r.wg.Done()
			fmt.Println("going to stop...")
			return
		}

		time.Sleep(5000)
		fmt.Println("running...")
	}
}
