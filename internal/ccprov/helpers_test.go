package ccprov

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// testConfig is an SDK configuration pointed at a fake, with the standard retryer making `attempts` attempts and no
// backoff: the SDK's retry decisions are real, only the waiting is removed.
func testConfig(endpoint string, attempts int) aws.Config {
	return aws.Config{
		Region:       "us-east-1",
		BaseEndpoint: aws.String(endpoint),
		Credentials:  credentials.NewStaticCredentialsProvider("AKIDTEST", "secret", ""),
		Retryer: func() aws.Retryer {
			return retry.NewStandard(func(o *retry.StandardOptions) {
				o.MaxAttempts = attempts
				o.Backoff = retry.BackoffDelayerFunc(func(int, error) (time.Duration, error) { return 0, nil })
			})
		},
	}
}

func noSleep(context.Context, time.Duration) error { return nil }

var (
	instantPatience = patience{attempts: notFoundPatience.attempts, base: time.Millisecond, max: time.Millisecond, sleep: noSleep}
	instantPacing   = pacing{first: time.Millisecond, max: time.Millisecond, sleep: noSleep, now: time.Now}
)
