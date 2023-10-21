package logger

import (
	"os"

	config "github.com/IASamoylov/tg_calories_observer/internal/config/debug"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger
var defaultLvl = zap.NewAtomicLevel()

func init() {
	SetLogger(New())
}

func Sync() error {
	return log.Sync()
}

// SetLogger заменяет по умолчанию инцилизированный логер
func SetLogger(logger *zap.SugaredLogger) {
	log = logger
}

// SetLogLvl переопределяет минимальный уровень логирования
func SetLogLvl(lvl string) error {
	return defaultLvl.UnmarshalText([]byte(lvl))
}

// New создает новый экземпляр zap.SugaredLogger
func New(additionalCores ...zapcore.Core) *zap.SugaredLogger {
	conf := zap.NewProductionEncoderConfig()
	conf.TimeKey = "time"
	conf.EncodeTime = zapcore.RFC3339TimeEncoder

	baseCores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewJSONEncoder(conf), zapcore.AddSync(os.Stdout), defaultLvl),
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

func Info(msg string, args ...interface{}) {
	log.Infow(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Debugw(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Warnw(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Errorw(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

func Panic(msg string, args ...interface{}) {
	log.Panicw(msg, args...)
}

func Panicf(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	log.Fatalw(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}
