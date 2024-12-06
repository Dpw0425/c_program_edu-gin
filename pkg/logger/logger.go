package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log *zap.SugaredLogger

type Level = zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = -1
	// InfoLevel is the default logging priority.
	InfoLevel Level = iota - 1
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

func InitLogger(filepath string, topic string) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	lowLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < ErrorLevel && l >= DebugLevel
	})
	highLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= ErrorLevel
	})

	lowFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath + "/low.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	lowFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(lowFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowLevel)

	highFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath + "/high.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
	highFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(highFileWriteSyncer, zapcore.AddSync(os.Stdout)), highLevel)

	core := zapcore.NewTee(lowFileCore, highFileCore)
	logger := zap.New(core, zap.AddCaller())
	log = logger.Sugar().With("topic", topic)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
