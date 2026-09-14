package awsprov

import (
	"context"
	"time"
)

// patience is how long a Read waits for a resource EC2 may not have propagated yet.
type patience struct {
	attempts int
	base     time.Duration
	max      time.Duration
	sleep    func(context.Context, time.Duration) error
}

// notFoundPatience: 5 attempts, 250ms doubling, capped at 2s — at most 3.75s of waiting.
var notFoundPatience = patience{attempts: 5, base: 250 * time.Millisecond, max: 2 * time.Second, sleep: sleepCtx}

// once reads a single time: Import, where the ID came from discovery and so is already visible.
var once = patience{attempts: 1}

// wait calls try until it reports found, returns an error, or the attempts run out — the last answering
// (false, nil). An error from try stops at once: only absence is worth waiting out.
func (pt patience) wait(ctx context.Context, try func() (bool, error)) (bool, error) {
	delay := pt.base
	for attempt := 1; ; attempt++ {
		found, err := try()
		if err != nil || found {
			return found, err
		}
		if attempt >= pt.attempts {
			return false, nil
		}
		if err := pt.sleep(ctx, delay); err != nil {
			return false, err
		}
		delay = min(delay*2, pt.max)
	}
}

func sleepCtx(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
