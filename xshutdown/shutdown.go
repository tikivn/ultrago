package xshutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/tikivn/ultrago/xlogaff"
)

func NewCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		logger := xlogaff.GetNewLogger()
		logger.Println("GracefulShutdown")
		cancel()
	}()
	return ctx
}
