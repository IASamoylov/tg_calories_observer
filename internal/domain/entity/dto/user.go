package dto

// User пользователь сервиса
type User struct {
	telegramID int64
	userName   string
	firstName  string
	lastName   string
	language   string
	agreement  bool
}

// TelegramID returns user telegram id
func (u User) TelegramID() int64 {
	return u.telegramID
}

// UserName returns name received from telegram
func (u User) UserName() string {
	return u.userName
}

// FirstName returns first name received from telegram
func (u User) FirstName() string {
	return u.firstName
}

// LastName returns last name received from telegram
func (u User) LastName() string {
	return u.lastName
}

// Language returns user language selected in telegram
func (u User) Language() string {
	return u.language
}

// Agreement agreement exchange personal data
func (u User) Agreement() bool {
	return u.agreement
}

// NewUser создает пользователя
func NewUser(
	telegramID int64,
	userName string,
	firstName string,
	lastName string,
	language string,
) User {
	return User{
		telegramID: telegramID,
		userName:   userName,
		firstName:  firstName,
		lastName:   lastName,
		language:   language,
	}
}

// ApplyAgreementExchangePersonalData пользователь согласился с хранением персональных данных
func (u User) ApplyAgreementExchangePersonalData() User {
	u.agreement = true

	return u
}
