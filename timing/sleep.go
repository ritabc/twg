package timing

import (
	"errors"
	"time"
)

var (
	ErrExceededMaxTries = errors.New("timing: exceeded max tries")
)

func PollUntil(fn func() bool, maxTries int) error {
	for i := 0; i < maxTries; i++ {
		time.Sleep(1 * time.Second)
		if fn() {
			return nil
		}
	}
	return ErrExceededMaxTries
}

type Poller struct {
}

// func (p *Poller) Until(fn func() bool, maxTries int) error {

// }
