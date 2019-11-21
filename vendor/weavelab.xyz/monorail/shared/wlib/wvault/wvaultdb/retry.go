package wvaultdb

import (
	"context"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wvault"
)

type (
	isRetryableFunc func(err error) bool
	doFunc          func(ctx context.Context) error
)

func doWithRetries(ctx context.Context, f doFunc, isRetryable isRetryableFunc, maxAttempts int) error {

	var err error
	for i := 0; i < maxAttempts; i++ {
		wvault.Clock.Sleep(delayTime(i))

		err = f(ctx)
		if err == nil {
			return nil
		}

		// based on error, should we retry?
		if isRetryable(err) == false {
			return err
		}

		// check context deadline exceeded
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	return err
}

// poor backoff implementation to avoid a dependency import
var attemptDelays = []time.Duration{
	time.Second * 1,
	time.Second * 5,
	time.Second * 15,
	time.Second * 30,
	time.Second * 60, //  will be reused for attempts indexed past this value
}

func delayTime(i int) time.Duration {
	if i >= len(attemptDelays) {
		return attemptDelays[len(attemptDelays)-1]
	}
	return attemptDelays[i]
}
