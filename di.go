package ultrago

import (
	"github.com/google/wire"
	"github.com/tikivn/ultrago/u_middleware"
)

var DiMiddlewares = wire.NewSet(
	u_middleware.NewLogMiddleware,
	u_middleware.NewMetricMiddleware,
)
