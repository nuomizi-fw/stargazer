package core

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

type StargazerLogger struct {
	*zap.Logger
}

func NewStargazerLogger() StargazerLogger {
	zapLogger, _ = zap.NewProduction()

	return StargazerLogger{zapLogger}
}
