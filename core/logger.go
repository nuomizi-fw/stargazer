package core

import (
	"os"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

type StargazerLogger struct {
	*zap.SugaredLogger
}

type FxLogger struct {
	*StargazerLogger
}

type FiberLogger struct {
	*StargazerLogger
}

func NewStargazerLogger() StargazerLogger {
	return newStargazerLogger(NewStargazerConfig())
}

func newStargazerLogger(config StargazerConfig) StargazerLogger {
	var zapConfig zap.Config

	if _, err := os.Stat(config.Logger.LogPath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.Logger.LogPath, 0o750); err != nil {
			panic(err)
		}
	}

	logOutputPath := config.Logger.LogPath + "/" + config.Logger.LogName + "." + config.Logger.LogExt

	if config.Server.Debug {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		zapConfig.OutputPaths = []string{logOutputPath}
	}

	switch config.Logger.LogLevel {
	case "debug":
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	case "info":
		zapConfig.Level.SetLevel(zapcore.InfoLevel)
	case "warn":
		zapConfig.Level.SetLevel(zapcore.WarnLevel)
	case "error":
		zapConfig.Level.SetLevel(zapcore.ErrorLevel)
	case "fatal":
		zapConfig.Level.SetLevel(zapcore.FatalLevel)
	default:
		zapConfig.Level.SetLevel(zapcore.PanicLevel)
	}

	zapLogger, _ = zapConfig.Build()

	return *newStargazerSugaredLogger(zapLogger)
}

func newStargazerSugaredLogger(logger *zap.Logger) *StargazerLogger {
	return &StargazerLogger{
		SugaredLogger: logger.Sugar(),
	}
}

// GetFxLogger gets logger for go-fx
func (l *StargazerLogger) GetFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return &FxLogger{newStargazerSugaredLogger(logger)}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.StargazerLogger.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.StargazerLogger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.StargazerLogger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.StargazerLogger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.StargazerLogger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.StargazerLogger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.StargazerLogger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.StargazerLogger.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.StargazerLogger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.StargazerLogger.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.StargazerLogger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.StargazerLogger.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
		}
	}
}

// Printf prints go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	if len(args) > 0 {
		l.Debugf(str, args)
	}
	l.Debug(str)
}

// Writer interface implementation for fiber-framework
func (l FiberLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
