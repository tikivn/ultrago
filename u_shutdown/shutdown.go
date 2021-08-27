package u_shutdown

import (
	"context"
	logaff "github.com/tikivn/ultrago/u_logaff"
	"os"
	"os/signal"
	"syscall"
)

func NewCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		logger := logaff.GetNewLogger()
		logger.Println("GracefulShutdown")
		cancel()
	}()
	return ctx
}
