package database

import (
	"context"

	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
)

// UserRepository ...
type UserRepository struct {
	types.PgxPool
}

// NewUserRepository ctor
func NewUserRepository(pool types.PgxPool) UserRepository {
	return UserRepository{PgxPool: pool}
}

// Upsert создает нового или обновляет существующего пользователя по telegramID
func (rep UserRepository) Upsert(ctx context.Context, user dto.User) error {
	sql := `insert into "user" (telegram_id, user_name, first_name, last_name, language)
			values ($1, $2, $3, $4, $5)
			on conflict (telegram_id) do update set user_name  = excluded.user_name,
													first_name = excluded.first_name,
													last_name  = excluded.last_name,
													language   = excluded.language;`

	_, err := rep.Exec(
		ctx,
		sql,
		user.TelegramID(),
		user.UserName(),
		user.FirstName(),
		user.LastName(),
		user.Language(),
	)

	return err
}

// ApplyAgreementExchangePersonalData сохраняет сиогласие об хранение персональных данных
func (rep UserRepository) ApplyAgreementExchangePersonalData(ctx context.Context, user dto.User) error {
	sql := `update "user" set agreement = $2 where telegram_id = $1`

	_, err := rep.Exec(
		ctx,
		sql,
		user.TelegramID(),
		user.Agreement(),
	)

	return err
}
