package core

import (
	"os"
	"sync"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger  *zap.Logger
	loggerOnce sync.Once
	loggerErr  error
)

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
	loggerOnce.Do(func() {
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

		zapLogger, loggerErr = zapConfig.Build()

		if loggerErr != nil {
			panic(loggerErr)
		}
	})

	return *newSugaredLogger(zapLogger)
}

func newSugaredLogger(logger *zap.Logger) *StargazerLogger {
	return &StargazerLogger{
		SugaredLogger: logger.Sugar(),
	}
}

// GetFxLogger gets logger for go-fx
func (l *StargazerLogger) GetFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(zap.WithCaller(false))
	return &FxLogger{newSugaredLogger(logger)}
}

func (l *StargazerLogger) GetFiberLogger() *FiberLogger {
	logger := zapLogger.WithOptions(zap.WithCaller(false))
	return &FiberLogger{newSugaredLogger(logger)}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
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
