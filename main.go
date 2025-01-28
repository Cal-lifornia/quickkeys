/*
Copyright © 2025 William Hobson willhobson@live.com.au

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"os"

	"github.com/Cal-lifornia/quickkeys/cmd"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var environment string

func initLogger() {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// TODO: Set log level option
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel

	})

	var consoleEncoderConfig zapcore.EncoderConfig
	var logFile string

	if environment == "prod" {
		consoleEncoderConfig = zap.NewProductionEncoderConfig()

		logFile = "/var/tmp/quickkeys.log"

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

func main() {
	initLogger()
	cmd.Execute()
}
