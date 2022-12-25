package u_monitor

import (
	"context"
	"fmt"
	"github.com/tikivn/ultrago/u_logger"
	"regexp"
	"runtime"
	"time"
)

// TimeTrack call defer TimeTrack(time.Now) at beginning of function that you want to track monitor
func TimeTrack(start time.Time) {
	logger := u_logger.NewLogger()
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	logger.Println(fmt.Sprintf("%s took %s", name, elapsed))
}

func TimeTrackWithCtx(ctx context.Context, start time.Time) {
	ctx, logger := u_logger.GetLogger(ctx)
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	logger.Info(fmt.Sprintf("%s took %s", name, elapsed))
}
