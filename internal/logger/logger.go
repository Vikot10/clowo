package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger = zap.NewNop()

var level = zap.NewAtomicLevelAt(zap.InfoLevel)

func SetLevel(l zapcore.Level) {
	level.SetLevel(l)
}
func GetLogger() *zap.Logger {
	return defaultLogger
}

func Init(isDebug bool) {
	devEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	prodEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := zap.Config{
		Level:         level,
		Development:   false,
		Encoding:      "json",
		EncoderConfig: prodEncoderConfig,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if isDebug {
		level.SetLevel(zap.DebugLevel)
		loggerConfig.Level = level
		loggerConfig.Development = true
		loggerConfig.Encoding = "console"
		loggerConfig.EncoderConfig = devEncoderConfig
		loggerConfig.Sampling = nil
	}

	l, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	defaultLogger = l
}

func Debug(s string, fields ...zap.Field) {
	defaultLogger.Debug(s, fields...)
}

func Info(s string, fields ...zap.Field) {
	defaultLogger.Info(s, fields...)
}

func Warn(s string, fields ...zap.Field) {
	defaultLogger.Warn(s, fields...)
}

func Error(s string, fields ...zap.Field) {
	defaultLogger.Error(s, fields...)
}
