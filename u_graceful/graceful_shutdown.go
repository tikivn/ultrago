package u_graceful

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tikivn/ultrago/u_logger"
)

func NewCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		u_logger.NewLogger().Println("Graceful shutdown")
		time.Sleep(1 * time.Second) // just to be safe that ingress delivery request at the same time system kill
		cancel()
	}()
	return ctx
}

func BlockListen(ctx context.Context, fn func() error) error {
	lisErr := make(chan error, 1)
	go func() {
		if e := fn(); e != nil {
			lisErr <- e
		} else {
			close(lisErr)
		}
	}()
	for {
		select {
		case err, _ := <-lisErr:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
