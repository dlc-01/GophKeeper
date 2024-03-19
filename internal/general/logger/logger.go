package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	DevelopMode = "develop"
	ProductMode = "product"
)

type ConfigLog struct {
	AppMode              string
	LoggerDirectory      string
	LoggerFileMaxSize    int
	LoggerFileMaxBackups int
	LoggerFileMaxAge     int
	LoggerFileCompress   bool
}

type Logger struct {
	zap.SugaredLogger
}

// InitLogger — initialization function of zap logger.
func Initialize(cfg ConfigLog) (*Logger, error) {
	switch cfg.AppMode {
	case DevelopMode:
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("cannot init logger: %w", err)
		}
		lgr := &Logger{
			SugaredLogger: *logger.Sugar(),
		}
		return lgr, nil

	case ProductMode:
		loggerConf := zap.NewProductionEncoderConfig()
		loggerConf.EncodeTime = zapcore.ISO8601TimeEncoder
		fileEncoder := zapcore.NewJSONEncoder(loggerConf)
		defaultLogLevel := zapcore.DebugLevel
		dir := fmt.Sprintf("%v/", cfg.LoggerDirectory)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return nil, err
		}
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%v/%v.log", dir, time.Now().Format("2022-02-24")),
			MaxSize:    cfg.LoggerFileMaxSize,
			MaxBackups: cfg.LoggerFileMaxBackups,
			MaxAge:     cfg.LoggerFileMaxAge,
			Compress:   cfg.LoggerFileCompress,
		})
		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		lgr := &Logger{
			SugaredLogger: *zap.New(zapcore.NewTee(core)).Sugar(),
		}
		return lgr, nil
	default:
		return nil, fmt.Errorf("error while: creating logger, not supported mode")
	}
}

// Fatalf — function that equals Fatalf error in the logger.
func (l Logger) Fatalf(format string, opts ...any) {
	l.SugaredLogger.Fatalf(format, opts)
}

// Errorf — function that equals Errorf error in the logger.
func (l Logger) Errorf(format string, opts ...any) {
	l.SugaredLogger.Errorf(format, opts)
}

// Infof — function that equals Infof error in the logger.
func (l Logger) Infof(format string, opts ...any) {
	l.SugaredLogger.Infof(format, opts)
}

// Warnf — function that equals Warnf error in the logger.
func (l Logger) Warnf(format string, opts ...any) {
	l.SugaredLogger.Warnf(format, opts)
}

// Panicf — function that equals Panicf error in the logger.
func (l Logger) Panicf(format string, opts ...any) {
	l.SugaredLogger.Panicf(format, opts)
}

// Info — function that equals Info error in the logger.
func (l Logger) Info(msg string) {
	l.SugaredLogger.Info(msg)
}
