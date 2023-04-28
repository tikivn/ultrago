package u_retry

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func BackoffRetryWithContext(ctx context.Context, operation func() error) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 1 * time.Hour
	bCtx := backoff.WithContext(bo, ctx)

	err := backoff.Retry(operation, bCtx)
	if err != nil {
		return err
	}
	return nil
}
