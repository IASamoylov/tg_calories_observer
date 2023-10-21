package database

import (
	"context"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/dto"
)

// UserRepository ...
type UserRepository struct {
	PgxPool
}

// NewUserRepository ctor
func NewUserRepository(pool PgxPool) UserRepository {
	return UserRepository{PgxPool: pool}
}

// UpsertAndGet создает нового или обновляет существующего пользователя в базе данных по его внутреннему ID
// и возвращает текущего пользователя
func (rep UserRepository) UpsertAndGet(ctx context.Context, user dto.User) (dto.User, error) {
	sql := `insert into "user" (telegram_id, user_name, first_name, last_name, language)
			values ($1, $2, $3, $4, $5)
			on conflict (telegram_id) do update set user_name  = excluded.user_name,
													first_name = excluded.first_name,
													last_name  = excluded.last_name,
													language   = excluded.language
			returning id;`

	row := rep.QueryRow(
		ctx,
		sql,
		user.TelegramID(),
		user.UserName(),
		user.FirstName(),
		user.LastName(),
		user.Language(),
	)

	var userID int64
	if err := row.Scan(&userID); err != nil {
		return user, err
	}

	return dto.NewUser(
		userID,
		user.TelegramID(),
		user.UserName(),
		user.FirstName(),
		user.LastName(),
		user.Language(),
	), nil
}
