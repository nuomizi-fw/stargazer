package logger

import (
	"github.com/nuomizi-fw/stargazer/core"
)

var logger core.StargazerLogger

// Debug logs the provided arguments at [DebugLevel].
// Spaces are added between arguments when neither is a string.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info logs the provided arguments at [InfoLevel].
// Spaces are added between arguments when neither is a string.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn logs the provided arguments at [WarnLevel].
// Spaces are added between arguments when neither is a string.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error logs the provided arguments at [ErrorLevel].
// Spaces are added between arguments when neither is a string.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// DPanic logs the provided arguments at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are added between arguments when neither is a string.
func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

// Panic constructs a message with the provided arguments and panics.
// Spaces are added between arguments when neither is a string.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Fatal constructs a message with the provided arguments and calls os.Exit.
// Spaces are added between arguments when neither is a string.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Debugf formats the message according to the format specifier
// and logs it at [DebugLevel].
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Infof formats the message according to the format specifier
// and logs it at [InfoLevel].
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Warnf formats the message according to the format specifier
// and logs it at [WarnLevel].
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Errorf formats the message according to the format specifier
// and logs it at [ErrorLevel].
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// DPanicf formats the message according to the format specifier
// and logs it at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

// Panicf formats the message according to the format specifier
// and panics.
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}

// Debugln logs a message at [DebugLevel].
// Spaces are always added between arguments.
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Infoln logs a message at [InfoLevel].
// Spaces are always added between arguments.
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Warnln logs a message at [WarnLevel].
// Spaces are always added between arguments.
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Errorln logs a message at [ErrorLevel].
// Spaces are always added between arguments.
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

// DPanicln logs a message at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are always added between arguments.
func DPanicln(args ...interface{}) {
	logger.DPanicln(args...)
}

// Panicln logs a message at [PanicLevel] and panics.
// Spaces are always added between arguments.
func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}

// Fatalln logs a message at [FatalLevel] and calls os.Exit.
// Spaces are always added between arguments.
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}
