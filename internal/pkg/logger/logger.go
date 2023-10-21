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
	err := defaultLvl.UnmarshalText([]byte(lvl))

	if err != nil {
		Error("failed to set log level as "+lvl, "reason", err.Error())
		Warn("log level is set as ", defaultLvl.String())
	} else {
		Debug("log level was successfully set as " + lvl)
	}

	return err
}

// New создает новый экземпляр zap.SugaredLogger
func New(additionalCores ...zapcore.Core) *zap.SugaredLogger {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.RFC3339TimeEncoder

	baseCores := []zapcore.Core{
		// json console logger
		zapcore.NewCore(zapcore.NewJSONEncoder(conf), zapcore.AddSync(os.Stdout), defaultLvl),
	}

	cores := append(baseCores, additionalCores...)
	core := zapcore.
		NewTee(cores...).
		With([]zap.Field{
			{Key: "app_name", String: config.AppName, Type: zapcore.StringType},
			{Key: "version", String: config.Version, Type: zapcore.StringType},
			{Key: "build_time", String: config.BuildTime, Type: zapcore.StringType},
			{Key: "github_sha", String: config.GithubSHA, Type: zapcore.StringType},
			{Key: "github_sha_short", String: config.GithubSHAShort, Type: zapcore.StringType},
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
