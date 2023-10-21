package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gookit/slog"
)

// BotTelegramClient client for telegram bot API
type BotTelegramClient interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type channelErrorLoggerHandler struct {
	// Formatter log message formatter. default use TextFormatter
	Formatter slog.Formatter
	channelID int64
	client    BotTelegramClient
}

//
//
//func (t *channelErrorLoggerHandler) Handle(record *slog.Record) error {
//	log, err := t.Formatter.Format(record)
//	if err != nil {
//		return err
//	}
//	var prettyJSON bytes.Buffer
//	if err = json.Indent(&prettyJSON, log, "", "    "); err != nil {
//		return err
//	}
//	msg := tgbotapi.NewMessage(t.channelID, prettyJSON.String())
//	_, err = t.client.Send(msg)
//	return err
//}
//
//func NewChannelErrorLoggerHandler(channelID int64, client BotTelegramClient) zapcore.Core {
//	return &channelErrorLoggerHandler{
//		channelID: channelID,
//		client:    client,
//		Formatter: slog.NewJSONFormatter(),
//	}
//}
