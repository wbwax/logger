package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// only supports the following level
var (
	gDebug string = "debug"
	gInfo  string = "info"
	gWarn  string = "warn"
	gError string = "error"
)

var (
	gLogger        *zap.Logger
	gSugaredLogger *zap.SugaredLogger
)

// Init inits logger
func Init(cfg Config) error {
	var instances []string
	switch strings.ToLower(cfg.Level) {
	case gDebug:
		instances = []string{gDebug, gInfo, gWarn, gError}
	case gInfo:
		instances = []string{gInfo, gWarn, gError}
	case gWarn:
		instances = []string{gWarn, gError}
	case gError:
		instances = []string{gError}
	default: // default to info
		instances = []string{gInfo, gWarn, gError}
	}
	cores := make([]zapcore.Core, 0, len(instances))
	for _, instance := range instances {
		core, err := createZapCore(instance, cfg)
		if err != nil {
			return err
		}
		cores = append(cores, core)
	}
	zapLogger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	gLogger = zapLogger
	gSugaredLogger = zapLogger.Sugar()
	return nil
}

func createZapCore(instance string, cfg Config) (zapcore.Core, error) {
	level, err := levelStringToZapLevel(instance)
	if err != nil {
		return nil, err
	}
	writer := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", cfg.Path, instance),
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  true,
	}
	writeSyncer := zapcore.AddSync(writer)

	// build EncoderConfig for production environments.
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.TimeKey = "time"
	encoderConf.CallerKey = "line"
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	if cfg.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConf)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConf) // default to console encoding
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	return core, nil
}

func levelStringToZapLevel(levelStr string) (zapcore.Level, error) {
	switch levelStr {
	case gDebug:
		return zapcore.DebugLevel, nil
	case gInfo:
		return zapcore.InfoLevel, nil
	case gWarn:
		return zapcore.WarnLevel, nil
	case gError:
		return zapcore.ErrorLevel, nil
	}
	return zapcore.ErrorLevel, fmt.Errorf("unknown level '%s'", levelStr)
}

// Sync flushes any buffered log
func Sync() {
	if gLogger != nil {
		gLogger.Sync()
	}
	if gSugaredLogger != nil {
		gSugaredLogger.Sync()
	}
}

// Debugf calls Debugf in zap to log a templated message.
func Debugf(template string, args ...interface{}) {
	gSugaredLogger.Debugf(template, args...)
}

// Infof calls Infof in zap to log a templated message.
func Infof(template string, args ...interface{}) {
	gSugaredLogger.Infof(template, args...)
}

// Warnf calls Warnf in zap to log a templated message.
func Warnf(template string, args ...interface{}) {
	gSugaredLogger.Warnf(template, args...)
}

// Errorf calls Errorf in zap to log a templated message.
func Errorf(template string, args ...interface{}) {
	gSugaredLogger.Errorf(template, args...)
}

// Debug calls Debug in zap to log
func Debug(msg string, fields ...zap.Field) {
	gLogger.Debug(msg, fields...)
}

// Info calls Info in zap to log
func Info(msg string, fields ...zap.Field) {
	gLogger.Info(msg, fields...)
}

// Warn calls Warn in zap to log
func Warn(msg string, fields ...zap.Field) {
	gLogger.Warn(msg, fields...)
}

// Error calls Error in zap to log
func Error(msg string, fields ...zap.Field) {
	gLogger.Error(msg, fields...)
}
