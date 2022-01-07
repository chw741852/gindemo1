package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func Init() {
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleWriter := zapcore.AddSync(os.Stdout)

	// file, _ := os.Create("./test.log")
	// fileWriter := zapcore.AddSync(file)
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
		LocalTime:  true,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleWriter, zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriter, zapcore.DebugLevel),
	)
	Log = zap.New(core, zap.AddCaller())
}
