package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleWriter := zapcore.AddSync(os.Stdout)

	file, _ := os.Create("./test.log")
	fileWriter := zapcore.AddSync(file)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleWriter, zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriter, zapcore.DebugLevel),
	)
	Log = zap.New(core)
}
