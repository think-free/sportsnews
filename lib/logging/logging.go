/*
	This is a simple logging helper that use zap
	Usage:
		- Init the logger with the desired log level, if not default to debug
			Availables levels are: debug, info, warn, error, fatal, panic
			-> logging.Init("info")
		- Use the logger with the context
			-> logging.L(ctx).Infof("my message")
		- Add tags to the context in any part of the code, the tags will be added to the following log messages that use the context
			-> ctx = logging.SetTag(ctx, "mytag", "myvalue")
		- Use the logger with the context after adding tags
			The tags will be added to the log message
			-> logging.L(ctx).Infof("my message")

	2022 - Christophe Meurice (meumeu1402@gmail.com)
*/

package logging

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type t string

var tgs t = "tags"
var defaultLogger *zap.Logger

func init() {
	Init("debug")
}

func Init(l string) {
	level := getZapLevelFromString(l)

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.Level = zap.NewAtomicLevelAt(level)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	defaultLogger, err = config.Build()
	if err != nil {
		panic(err)
	}
	// Set StdLog and globals to use the default logger
	_ = zap.RedirectStdLog(defaultLogger)
	_ = zap.ReplaceGlobals(defaultLogger)
}

func L(ctx context.Context) *zap.SugaredLogger {
	logger := defaultLogger.Sugar()
	if tags := ctx.Value(tgs); tags != nil {
		for k, v := range tags.(map[string]interface{}) {
			logger = logger.With(k, v)
		}
	}
	return logger
}

func SetTag(ctx context.Context, key string, value interface{}) context.Context {
	values := ctx.Value(tgs)
	if values == nil {
		values = make(map[string]interface{})
	}
	values.(map[string]interface{})[key] = value
	return context.WithValue(ctx, tgs, values)
}

func getZapLevelFromString(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	case "panic":
		return zap.PanicLevel
	}
	return zap.InfoLevel
}
