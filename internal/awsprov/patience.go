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
