package assert

import "time"

// LinearTimer provides a simple linear backoff timer
type LinearTimer time.Duration

// Tick returns the next duration in the linear backoff sequence
func (lt LinearTimer) Tick() time.Duration {
	return time.Duration(lt)
}
