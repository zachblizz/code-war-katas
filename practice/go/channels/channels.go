package channels

import (
	"fmt"
	"time"

)

// Runner - runner type
type Runner struct {
	Stop chan bool
}

// NewRunner - returns a new runner
func NewRunner(stop chan bool) *Runner {
	return &Runner{stop}
}

// Run - runs the stuff...
func (r *Runner) Run() {
	fmt.Println("going to start...")
	for {
		s := <-r.Stop

		if s {
			fmt.Println("going to stop...")
			return
		}

		time.Sleep(5000)
		fmt.Println("running...")
	}
}
