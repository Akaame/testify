package assert

import (
	"math"
	"time"
)

// ExponentialTimer provides an exponential backoff timer
type ExponentialTimer struct {
	Base        time.Duration
	Factor      float64
	currentTick int
}

// NewExponentialTimer creates a new ExponentialTimer with the given base duration and factor
func NewExponentialTimer(base time.Duration, factor float64) *ExponentialTimer {
	return &ExponentialTimer{
		Base:        base,
		Factor:      factor,
		currentTick: 0,
	}
}

// Tick returns the next duration in the exponential backoff sequence
func (et ExponentialTimer) Tick() time.Duration {
	return time.Duration(float64(et.Base) * math.Pow(et.Factor, float64(et.currentTick)))
}
