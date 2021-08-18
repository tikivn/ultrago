package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tikivn/tially/pkg/util/logaff"
)

func NewCtx() context.Context {
	logger := logaff.GetNewLogger()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		logger.Println("Graceful shutdown")
		time.Sleep(1 * time.Second) // just to be safe that ingress delivery request at the same time system kill
		cancel()
	}()
	return ctx
}
