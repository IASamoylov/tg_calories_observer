package database

import (
	"context"
	"encoding/hex"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
)

// Cryptor ...
type Cryptor interface {
	Encrypt(message []byte) ([]byte, error)
	Decrypt(message []byte) ([]byte, error)
}

// SecurityUserRepository ...
type SecurityUserRepository struct {
	userRepository UserRepository
	cryptor        Cryptor
}

// NewSecurityUserRepository ctor
func NewSecurityUserRepository(userRepository UserRepository, cryptor Cryptor) SecurityUserRepository {
	return SecurityUserRepository{userRepository: userRepository, cryptor: cryptor}
}

// Upsert шифрует персональные данные пользователя перед тем как созданить или обновить
func (rep SecurityUserRepository) Upsert(ctx context.Context, user dto.User) error {
	userName, err := rep.cryptor.Encrypt([]byte(user.UserName()))
	if err != nil {
		return err
	}
	firstName, err := rep.cryptor.Encrypt([]byte(user.FirstName()))
	if err != nil {
		return err
	}
	lastName, err := rep.cryptor.Encrypt([]byte(user.LastName()))
	if err != nil {
		return err
	}

	return rep.userRepository.Upsert(ctx, dto.NewUser(
		user.TelegramID(),
		hex.EncodeToString(userName),
		hex.EncodeToString(firstName),
		hex.EncodeToString(lastName),
		user.Language(),
	))
}

// ApplyAgreementExchangePersonalData сохраняет сиогласие об хранение персональных данных
func (rep SecurityUserRepository) ApplyAgreementExchangePersonalData(ctx context.Context, user dto.User) error {
	return rep.userRepository.ApplyAgreementExchangePersonalData(ctx, user)
}
