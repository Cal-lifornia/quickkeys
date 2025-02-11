package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(config *Config, environment string) {

	var chosenLogLevel zapcore.Level = zapcore.InfoLevel

	switch config.LogLevel {
	case "debug":
		chosenLogLevel = zapcore.DebugLevel
	case "info":
		chosenLogLevel = zapcore.InfoLevel
	case "warn":
		chosenLogLevel = zapcore.WarnLevel
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return (lvl >= zapcore.ErrorLevel && lvl >= chosenLogLevel)
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return (lvl < zapcore.ErrorLevel && lvl >= chosenLogLevel)
	})

	var consoleEncoderConfig zapcore.EncoderConfig
	var logFile string

	if environment == "prod" {
		consoleEncoderConfig = zap.NewProductionEncoderConfig()
		logFile = "/var/tmp/quickkeys.log"

	} else if environment == "testing" {
		consoleEncoderConfig = zap.NewDevelopmentEncoderConfig()
		logFile = "./test.log"
	} else {
		consoleEncoderConfig = zap.NewDevelopmentEncoderConfig()
		logFile = "./debug.log"
	}

	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	consoleErrors := zapcore.Lock(os.Stderr)

	fileLog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     7,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, fileLog, lowPriority),
	)

	var logger *zap.Logger

	if environment == "prod" {
		logger = zap.New(core)
	} else {
		logger = zap.New(core, zap.AddCaller())
	}

	zap.ReplaceGlobals(logger)

	zap.L().Debug("logger initialised")
}
