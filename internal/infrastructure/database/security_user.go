package database

import (
	"context"
	"encoding/hex"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/dto"
)

// SecurityUserRepository ...
type SecurityUserRepository struct {
	userRepository UserRepository
	cryptor        Cryptor
}

// NewSecurityUserRepository ctor
func NewSecurityUserRepository(userRepository UserRepository, cryptor Cryptor) SecurityUserRepository {
	return SecurityUserRepository{userRepository: userRepository, cryptor: cryptor}
}

// UpsertAndGet шифрует персональные данные пользователя перед тем как созданить или обновить
// и возвращает текущего пользователя
func (rep SecurityUserRepository) UpsertAndGet(ctx context.Context, user dto.User) (dto.User, error) {
	userName, err := rep.cryptor.Encrypt([]byte(user.UserName()))
	if err != nil {
		return user, err
	}
	firstName, err := rep.cryptor.Encrypt([]byte(user.FirstName()))
	if err != nil {
		return user, err
	}
	lastName, err := rep.cryptor.Encrypt([]byte(user.LastName()))
	if err != nil {
		return user, err
	}

	newUser, err := rep.userRepository.UpsertAndGet(ctx, dto.NewDefaultUser(
		user.TelegramID(),
		hex.EncodeToString(userName),
		hex.EncodeToString(firstName),
		hex.EncodeToString(lastName),
		user.Language(),
	))

	return dto.NewUser(
		newUser.ID(),
		user.TelegramID(),
		user.UserName(),
		user.FirstName(),
		user.LastName(),
		user.Language(),
	), err
}
