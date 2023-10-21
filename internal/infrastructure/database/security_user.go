package database

import (
	"context"
	"encoding/hex"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
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

// UpsertAndGet creates or updates the user and returns his internal system ID
func (rep SecurityUserRepository) UpsertAndGet(ctx context.Context, user domain.User) (domain.User, error) {
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

	newUser, err := rep.userRepository.UpsertAndGet(ctx, domain.NewDefaultUser(
		user.TelegramID(),
		hex.EncodeToString(userName),
		hex.EncodeToString(firstName),
		hex.EncodeToString(lastName),
		user.Language(),
	))

	return domain.NewUser(
		newUser.ID(),
		user.TelegramID(),
		user.UserName(),
		user.FirstName(),
		user.LastName(),
		user.Language(),
	), err
}
