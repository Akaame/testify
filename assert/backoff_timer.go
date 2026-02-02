package assert

import "time"

type BackoffTimer interface {
	Tick() time.Duration
}
