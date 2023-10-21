package telegram

import (
	"bytes"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gookit/slog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// BotTelegramClient client for telegram bot API
type BotTelegramClient interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type writeSyncer struct {
	// Formatter log message formatter. default use TextFormatter
	Formatter slog.Formatter
	channelID int64
	client    BotTelegramClient
}

func (ws writeSyncer) Write(p []byte) (n int, err error) {
	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, p, "", "    "); err != nil {
		return 0, err
	}
	msg := tgbotapi.NewMessage(ws.channelID, prettyJSON.String())
	_, err = ws.client.Send(msg)

	return len(p), err
}

func (ws writeSyncer) Sync() error {
	return nil
}

func newWriteSyncer(channelID int64, client BotTelegramClient) zapcore.WriteSyncer {
	return &writeSyncer{
		channelID: channelID,
		client:    client,
		Formatter: slog.NewJSONFormatter(),
	}
}

func NewChannelErrorLoggerCore(channelID int64, client BotTelegramClient) zapcore.Core {
	conf := zap.NewProductionEncoderConfig()
	conf.TimeKey = "time"
	conf.EncodeTime = zapcore.RFC3339TimeEncoder
	writer := newWriteSyncer(channelID, client)

	return zapcore.NewCore(zapcore.NewJSONEncoder(conf), writer, zapcore.ErrorLevel)
}
