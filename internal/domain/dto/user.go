package dto

// TelegramID ...
type TelegramID int64

// User describes users in the service
type User struct {
	id         int64
	telegramID TelegramID
	userName   string
	firstName  string
	lastName   string
	language   string
}

// ID returns user internal id
func (u User) ID() int64 {
	return u.id
}

// TelegramID returns user telegram id
func (u User) TelegramID() TelegramID {
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

// NewDefaultUser creates a new user without internal id
func NewDefaultUser(
	telegramID TelegramID,
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

// NewUser creates a new user with internal id
func NewUser(
	ID int64,
	telegramID TelegramID,
	userName string,
	firstName string,
	lastName string,
	language string,
) User {
	return User{
		id:         ID,
		telegramID: telegramID,
		userName:   userName,
		firstName:  firstName,
		lastName:   lastName,
		language:   language,
	}
}
