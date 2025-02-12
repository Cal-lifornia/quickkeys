package appkeys

import "go.uber.org/zap"

var logger *zap.Logger = zap.L().With(
	zap.String("service", "appkeys"),
)
