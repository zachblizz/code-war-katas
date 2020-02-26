package channels

import (
	"fmt"
	"time"

)

// IRunner - runner interface...
type IRunner interface {
	run()
}

// Runner - runner type
type Runner struct {
	Stop chan bool
}

// NewRunner - returns a new runner
func NewRunner(stop &chan bool) *Runner {
	return &Runner{stop}
}

func (r *Runner) run() {
	for true {
		select {
		case s := <-r.Stop:
			fmt.Println("stopping....")
			return
		default:
			fmt.Println("continuing...")
		}

		time.Sleep(5000)
		fmt.Println("running...")
	}
}
