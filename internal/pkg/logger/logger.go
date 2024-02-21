package logger

import (
	"io"
	"os"

	config "github.com/IASamoylov/tg_calories_observer/internal/config/debug"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger
var defaultLvl = zap.NewAtomicLevel()

func init() {
	SetLogger(NewDefault())
}

// Sync ...
func Sync() error {
	return log.Sync()
}

// SetLogger заменяет по умолчанию инцилизированный логер
func SetLogger(logger *zap.SugaredLogger) {
	log = logger
}

// SetLogLvl переопределяет минимальный уровень логирования
//func SetLogLvl(lvl string) error {
//	return defaultLvl.UnmarshalText([]byte(lvl))
//}

// NewDefault создает новый экземпляр zap.SugaredLogger
func NewDefault(additionalCores ...zapcore.Core) *zap.SugaredLogger {
	return New(os.Stdout, additionalCores...)
}

// New создает новый экземпляр zap.SugaredLogger с возможность переопределить stdout
func New(writer io.Writer, additionalCores ...zapcore.Core) *zap.SugaredLogger {
	conf := zap.NewProductionEncoderConfig()
	conf.TimeKey = "time"
	conf.EncodeTime = zapcore.RFC3339TimeEncoder

	baseCores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewJSONEncoder(conf), zapcore.AddSync(writer), defaultLvl),
	}

	cores := append(baseCores, additionalCores...)
	core := zapcore.
		NewTee(cores...).
		With([]zap.Field{
			{Key: "app_name", String: config.AppName, Type: zapcore.StringType},
			{Key: "version", String: config.Version, Type: zapcore.StringType},
		})

	logger := zap.New(core)

	return logger.Sugar()
}

// Info ...
func Info(msg string, args ...interface{}) {
	log.Infow(msg, args...)
}

// Infof ...
func Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

// Warn ...
func Warn(msg string, args ...interface{}) {
	log.Warnw(msg, args...)
}

// Error ...
func Error(msg string, args ...interface{}) {
	log.Errorw(msg, args...)
}

// Errorf ...
func Errorf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

// Panicf ...
func Panicf(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

// Fatalf ...
func Fatalf(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}
