package embedded

import (
	"time"
)

type Pollfunc func() interface{}

func Poll(f Pollfunc, period time.Duration, sink chan interface{}) chan bool {
	kill := make(chan bool)
	go func(kill chan bool) {
		for {
			select {
			case <-kill:
				return
			default:
				if period > 0 {
					time.Sleep(period)
				}
				sink <- f()
			}
		}
	}(kill)
	return kill
}
