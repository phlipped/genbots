package main

import (
	"github.com/phlipped/genbots/arena"
)

// Run until interrupted 
func RunLoop(stop <-chan struct{}, a arena.Arena) {
	for {
		select {
		case <-stop:
			break
		default:
			a.Step()
		}
	}
}

func main() {

}
