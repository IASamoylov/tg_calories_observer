// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/IASamoylov/tg_calories_observer/internal/pkg/types (interfaces: TelegramBotAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockTelegramBotAPI is a mock of TelegramBotAPI interface.
type MockTelegramBotAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramBotAPIMockRecorder
}

// MockTelegramBotAPIMockRecorder is the mock recorder for MockTelegramBotAPI.
type MockTelegramBotAPIMockRecorder struct {
	mock *MockTelegramBotAPI
}

// NewMockTelegramBotAPI creates a new mock instance.
func NewMockTelegramBotAPI(ctrl *gomock.Controller) *MockTelegramBotAPI {
	mock := &MockTelegramBotAPI{ctrl: ctrl}
	mock.recorder = &MockTelegramBotAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramBotAPI) EXPECT() *MockTelegramBotAPIMockRecorder {
	return m.recorder
}

// CopyMessage mocks base method.
func (m *MockTelegramBotAPI) CopyMessage(arg0 tgbotapi.CopyMessageConfig) (tgbotapi.MessageID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyMessage", arg0)
	ret0, _ := ret[0].(tgbotapi.MessageID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyMessage indicates an expected call of CopyMessage.
func (mr *MockTelegramBotAPIMockRecorder) CopyMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyMessage", reflect.TypeOf((*MockTelegramBotAPI)(nil).CopyMessage), arg0)
}

// GetChat mocks base method.
func (m *MockTelegramBotAPI) GetChat(arg0 tgbotapi.ChatInfoConfig) (tgbotapi.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChat", arg0)
	ret0, _ := ret[0].(tgbotapi.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChat indicates an expected call of GetChat.
func (mr *MockTelegramBotAPIMockRecorder) GetChat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChat", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetChat), arg0)
}

// GetChatAdministrators mocks base method.
func (m *MockTelegramBotAPI) GetChatAdministrators(arg0 tgbotapi.ChatAdministratorsConfig) ([]tgbotapi.ChatMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatAdministrators", arg0)
	ret0, _ := ret[0].([]tgbotapi.ChatMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatAdministrators indicates an expected call of GetChatAdministrators.
func (mr *MockTelegramBotAPIMockRecorder) GetChatAdministrators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatAdministrators", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetChatAdministrators), arg0)
}

// GetChatMember mocks base method.
func (m *MockTelegramBotAPI) GetChatMember(arg0 tgbotapi.GetChatMemberConfig) (tgbotapi.ChatMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatMember", arg0)
	ret0, _ := ret[0].(tgbotapi.ChatMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatMember indicates an expected call of GetChatMember.
func (mr *MockTelegramBotAPIMockRecorder) GetChatMember(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatMember", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetChatMember), arg0)
}

// GetChatMembersCount mocks base method.
func (m *MockTelegramBotAPI) GetChatMembersCount(arg0 tgbotapi.ChatMemberCountConfig) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatMembersCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatMembersCount indicates an expected call of GetChatMembersCount.
func (mr *MockTelegramBotAPIMockRecorder) GetChatMembersCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatMembersCount", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetChatMembersCount), arg0)
}

// GetFile mocks base method.
func (m *MockTelegramBotAPI) GetFile(arg0 tgbotapi.FileConfig) (tgbotapi.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", arg0)
	ret0, _ := ret[0].(tgbotapi.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockTelegramBotAPIMockRecorder) GetFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetFile), arg0)
}

// GetFileDirectURL mocks base method.
func (m *MockTelegramBotAPI) GetFileDirectURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileDirectURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileDirectURL indicates an expected call of GetFileDirectURL.
func (mr *MockTelegramBotAPIMockRecorder) GetFileDirectURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileDirectURL", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetFileDirectURL), arg0)
}

// GetGameHighScores mocks base method.
func (m *MockTelegramBotAPI) GetGameHighScores(arg0 tgbotapi.GetGameHighScoresConfig) ([]tgbotapi.GameHighScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameHighScores", arg0)
	ret0, _ := ret[0].([]tgbotapi.GameHighScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameHighScores indicates an expected call of GetGameHighScores.
func (mr *MockTelegramBotAPIMockRecorder) GetGameHighScores(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameHighScores", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetGameHighScores), arg0)
}

// GetInviteLink mocks base method.
func (m *MockTelegramBotAPI) GetInviteLink(arg0 tgbotapi.ChatInviteLinkConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInviteLink", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInviteLink indicates an expected call of GetInviteLink.
func (mr *MockTelegramBotAPIMockRecorder) GetInviteLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInviteLink", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetInviteLink), arg0)
}

// GetMe mocks base method.
func (m *MockTelegramBotAPI) GetMe() (tgbotapi.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMe")
	ret0, _ := ret[0].(tgbotapi.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMe indicates an expected call of GetMe.
func (mr *MockTelegramBotAPIMockRecorder) GetMe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMe", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetMe))
}

// GetMyCommands mocks base method.
func (m *MockTelegramBotAPI) GetMyCommands() ([]tgbotapi.BotCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyCommands")
	ret0, _ := ret[0].([]tgbotapi.BotCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyCommands indicates an expected call of GetMyCommands.
func (mr *MockTelegramBotAPIMockRecorder) GetMyCommands() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyCommands", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetMyCommands))
}

// GetMyCommandsWithConfig mocks base method.
func (m *MockTelegramBotAPI) GetMyCommandsWithConfig(arg0 tgbotapi.GetMyCommandsConfig) ([]tgbotapi.BotCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyCommandsWithConfig", arg0)
	ret0, _ := ret[0].([]tgbotapi.BotCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyCommandsWithConfig indicates an expected call of GetMyCommandsWithConfig.
func (mr *MockTelegramBotAPIMockRecorder) GetMyCommandsWithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyCommandsWithConfig", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetMyCommandsWithConfig), arg0)
}

// GetStickerSet mocks base method.
func (m *MockTelegramBotAPI) GetStickerSet(arg0 tgbotapi.GetStickerSetConfig) (tgbotapi.StickerSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStickerSet", arg0)
	ret0, _ := ret[0].(tgbotapi.StickerSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStickerSet indicates an expected call of GetStickerSet.
func (mr *MockTelegramBotAPIMockRecorder) GetStickerSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStickerSet", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetStickerSet), arg0)
}

// GetUpdates mocks base method.
func (m *MockTelegramBotAPI) GetUpdates(arg0 tgbotapi.UpdateConfig) ([]tgbotapi.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdates", arg0)
	ret0, _ := ret[0].([]tgbotapi.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdates indicates an expected call of GetUpdates.
func (mr *MockTelegramBotAPIMockRecorder) GetUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdates", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetUpdates), arg0)
}

// GetUpdatesChan mocks base method.
func (m *MockTelegramBotAPI) GetUpdatesChan(arg0 tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdatesChan", arg0)
	ret0, _ := ret[0].(tgbotapi.UpdatesChannel)
	return ret0
}

// GetUpdatesChan indicates an expected call of GetUpdatesChan.
func (mr *MockTelegramBotAPIMockRecorder) GetUpdatesChan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdatesChan", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetUpdatesChan), arg0)
}

// GetUserProfilePhotos mocks base method.
func (m *MockTelegramBotAPI) GetUserProfilePhotos(arg0 tgbotapi.UserProfilePhotosConfig) (tgbotapi.UserProfilePhotos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfilePhotos", arg0)
	ret0, _ := ret[0].(tgbotapi.UserProfilePhotos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfilePhotos indicates an expected call of GetUserProfilePhotos.
func (mr *MockTelegramBotAPIMockRecorder) GetUserProfilePhotos(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfilePhotos", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetUserProfilePhotos), arg0)
}

// GetWebhookInfo mocks base method.
func (m *MockTelegramBotAPI) GetWebhookInfo() (tgbotapi.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWebhookInfo")
	ret0, _ := ret[0].(tgbotapi.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebhookInfo indicates an expected call of GetWebhookInfo.
func (mr *MockTelegramBotAPIMockRecorder) GetWebhookInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebhookInfo", reflect.TypeOf((*MockTelegramBotAPI)(nil).GetWebhookInfo))
}

// HandleUpdate mocks base method.
func (m *MockTelegramBotAPI) HandleUpdate(arg0 *http.Request) (*tgbotapi.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdate", arg0)
	ret0, _ := ret[0].(*tgbotapi.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleUpdate indicates an expected call of HandleUpdate.
func (mr *MockTelegramBotAPIMockRecorder) HandleUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockTelegramBotAPI)(nil).HandleUpdate), arg0)
}

// IsMessageToMe mocks base method.
func (m *MockTelegramBotAPI) IsMessageToMe(arg0 tgbotapi.Message) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMessageToMe", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMessageToMe indicates an expected call of IsMessageToMe.
func (mr *MockTelegramBotAPIMockRecorder) IsMessageToMe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMessageToMe", reflect.TypeOf((*MockTelegramBotAPI)(nil).IsMessageToMe), arg0)
}

// ListenForWebhook mocks base method.
func (m *MockTelegramBotAPI) ListenForWebhook(arg0 string) tgbotapi.UpdatesChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenForWebhook", arg0)
	ret0, _ := ret[0].(tgbotapi.UpdatesChannel)
	return ret0
}

// ListenForWebhook indicates an expected call of ListenForWebhook.
func (mr *MockTelegramBotAPIMockRecorder) ListenForWebhook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenForWebhook", reflect.TypeOf((*MockTelegramBotAPI)(nil).ListenForWebhook), arg0)
}

// ListenForWebhookRespReqFormat mocks base method.
func (m *MockTelegramBotAPI) ListenForWebhookRespReqFormat(arg0 http.ResponseWriter, arg1 *http.Request) tgbotapi.UpdatesChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenForWebhookRespReqFormat", arg0, arg1)
	ret0, _ := ret[0].(tgbotapi.UpdatesChannel)
	return ret0
}

// ListenForWebhookRespReqFormat indicates an expected call of ListenForWebhookRespReqFormat.
func (mr *MockTelegramBotAPIMockRecorder) ListenForWebhookRespReqFormat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenForWebhookRespReqFormat", reflect.TypeOf((*MockTelegramBotAPI)(nil).ListenForWebhookRespReqFormat), arg0, arg1)
}

// MakeRequest mocks base method.
func (m *MockTelegramBotAPI) MakeRequest(arg0 string, arg1 tgbotapi.Params) (*tgbotapi.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", arg0, arg1)
	ret0, _ := ret[0].(*tgbotapi.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest.
func (mr *MockTelegramBotAPIMockRecorder) MakeRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockTelegramBotAPI)(nil).MakeRequest), arg0, arg1)
}

// Request mocks base method.
func (m *MockTelegramBotAPI) Request(arg0 tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", arg0)
	ret0, _ := ret[0].(*tgbotapi.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockTelegramBotAPIMockRecorder) Request(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockTelegramBotAPI)(nil).Request), arg0)
}

// Send mocks base method.
func (m *MockTelegramBotAPI) Send(arg0 tgbotapi.Chattable) (tgbotapi.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(tgbotapi.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockTelegramBotAPIMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTelegramBotAPI)(nil).Send), arg0)
}

// SendMediaGroup mocks base method.
func (m *MockTelegramBotAPI) SendMediaGroup(arg0 tgbotapi.MediaGroupConfig) ([]tgbotapi.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMediaGroup", arg0)
	ret0, _ := ret[0].([]tgbotapi.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMediaGroup indicates an expected call of SendMediaGroup.
func (mr *MockTelegramBotAPIMockRecorder) SendMediaGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMediaGroup", reflect.TypeOf((*MockTelegramBotAPI)(nil).SendMediaGroup), arg0)
}

// SetAPIEndpoint mocks base method.
func (m *MockTelegramBotAPI) SetAPIEndpoint(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAPIEndpoint", arg0)
}

// SetAPIEndpoint indicates an expected call of SetAPIEndpoint.
func (mr *MockTelegramBotAPIMockRecorder) SetAPIEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAPIEndpoint", reflect.TypeOf((*MockTelegramBotAPI)(nil).SetAPIEndpoint), arg0)
}

// StopPoll mocks base method.
func (m *MockTelegramBotAPI) StopPoll(arg0 tgbotapi.StopPollConfig) (tgbotapi.Poll, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopPoll", arg0)
	ret0, _ := ret[0].(tgbotapi.Poll)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopPoll indicates an expected call of StopPoll.
func (mr *MockTelegramBotAPIMockRecorder) StopPoll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPoll", reflect.TypeOf((*MockTelegramBotAPI)(nil).StopPoll), arg0)
}

// StopReceivingUpdates mocks base method.
func (m *MockTelegramBotAPI) StopReceivingUpdates() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopReceivingUpdates")
}

// StopReceivingUpdates indicates an expected call of StopReceivingUpdates.
func (mr *MockTelegramBotAPIMockRecorder) StopReceivingUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopReceivingUpdates", reflect.TypeOf((*MockTelegramBotAPI)(nil).StopReceivingUpdates))
}

// UploadFiles mocks base method.
func (m *MockTelegramBotAPI) UploadFiles(arg0 string, arg1 tgbotapi.Params, arg2 []tgbotapi.RequestFile) (*tgbotapi.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFiles", arg0, arg1, arg2)
	ret0, _ := ret[0].(*tgbotapi.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFiles indicates an expected call of UploadFiles.
func (mr *MockTelegramBotAPIMockRecorder) UploadFiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFiles", reflect.TypeOf((*MockTelegramBotAPI)(nil).UploadFiles), arg0, arg1, arg2)
}
