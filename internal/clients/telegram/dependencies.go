package telegram

import (
	"net/http"

	telegrambotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . ClinetTelegramAPI

// TelegramBotAPI telegram client
type TelegramBotAPI interface {
	SetAPIEndpoint(apiEndpoint string)
	MakeRequest(endpoint string, params telegrambotapi.Params) (*telegrambotapi.APIResponse, error)
	UploadFiles(endpoint string, params telegrambotapi.Params,
		files []telegrambotapi.RequestFile) (*telegrambotapi.APIResponse, error)
	GetFileDirectURL(fileID string) (string, error)
	GetMe() (telegrambotapi.User, error)
	IsMessageToMe(message telegrambotapi.Message) bool
	Request(c telegrambotapi.Chattable) (*telegrambotapi.APIResponse, error)
	Send(c telegrambotapi.Chattable) (telegrambotapi.Message, error)
	SendMediaGroup(config telegrambotapi.MediaGroupConfig) ([]telegrambotapi.Message, error)
	GetUserProfilePhotos(config telegrambotapi.UserProfilePhotosConfig) (telegrambotapi.UserProfilePhotos, error)
	GetFile(config telegrambotapi.FileConfig) (telegrambotapi.File, error)
	GetUpdates(config telegrambotapi.UpdateConfig) ([]telegrambotapi.Update, error)
	GetWebhookInfo() (telegrambotapi.WebhookInfo, error)
	GetUpdatesChan(config telegrambotapi.UpdateConfig) telegrambotapi.UpdatesChannel
	StopReceivingUpdates()
	ListenForWebhook(pattern string) telegrambotapi.UpdatesChannel
	ListenForWebhookRespReqFormat(w http.ResponseWriter, r *http.Request) telegrambotapi.UpdatesChannel
	HandleUpdate(r *http.Request) (*telegrambotapi.Update, error)
	GetChat(config telegrambotapi.ChatInfoConfig) (telegrambotapi.Chat, error)
	GetChatAdministrators(config telegrambotapi.ChatAdministratorsConfig) ([]telegrambotapi.ChatMember, error)
	GetChatMembersCount(config telegrambotapi.ChatMemberCountConfig) (int, error)
	GetChatMember(config telegrambotapi.GetChatMemberConfig) (telegrambotapi.ChatMember, error)
	GetGameHighScores(config telegrambotapi.GetGameHighScoresConfig) ([]telegrambotapi.GameHighScore, error)
	GetInviteLink(config telegrambotapi.ChatInviteLinkConfig) (string, error)
	GetStickerSet(config telegrambotapi.GetStickerSetConfig) (telegrambotapi.StickerSet, error)
	StopPoll(config telegrambotapi.StopPollConfig) (telegrambotapi.Poll, error)
	GetMyCommands() ([]telegrambotapi.BotCommand, error)
	GetMyCommandsWithConfig(config telegrambotapi.GetMyCommandsConfig) ([]telegrambotapi.BotCommand, error)
	CopyMessage(config telegrambotapi.CopyMessageConfig) (telegrambotapi.MessageID, error)
}
