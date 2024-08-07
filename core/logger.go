package core

import "go.uber.org/zap"

type StargazerLogger struct {
	*zap.SugaredLogger
}

func NewStargazerLogger() StargazerLogger {
	logger, _ := zap.NewProduction()
	return StargazerLogger{logger.Sugar()}
}
