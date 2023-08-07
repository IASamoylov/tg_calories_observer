package domain

// TelegramID ...
type TelegramID int64

type User struct {
	id         int64
	telegramID int64
	userName   string
	firstName  string
	lastName   string
	language   string
}

type SecureUser User

func (u User) ID() int64 {
	return u.id
}

func (u User) TelegramID() int64 {
	return u.telegramID
}

func (u User) UserName() string {
	return u.userName
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u User) Language() string {
	return u.language
}

func NewDefaultUser(
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

func NewUser(
	ID int64,
	telegramID int64,
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
