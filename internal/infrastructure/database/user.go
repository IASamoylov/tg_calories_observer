package database

import (
	"context"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
)

type UserRepository struct {
	PgxPool
	cryptor Cryptor
}

// NewUserRepository ctor
func NewUserRepository(pool PgxPool) UserRepository {
	return UserRepository{PgxPool: pool}
}

// UpsertAndGet creates or updates the user and returns his internal system ID
func (rep UserRepository) UpsertAndGet(ctx context.Context, user domain.User) (domain.User, error) {
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

	return domain.NewUser(
		userID,
		user.TelegramID(),
		user.UserName(),
		user.FirstName(),
		user.LastName(),
		user.Language(),
	), nil
}
