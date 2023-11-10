package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	t.Run("creates user with system ID", func(t *testing.T) {
		t.Parallel()

		user := NewUser(100500, "A", "B", "C", "en")

		assert.Equal(t, User{
			telegramID: 100500,
			userName:   "A",
			firstName:  "B",
			lastName:   "C",
			language:   "en",
		}, user)

		t.Run("getters return values correctly", func(t *testing.T) {
			assert.Equalf(t, int64(100500), user.TelegramID(), "telegram_id")
			assert.Equalf(t, "A", user.UserName(), "user_name")
			assert.Equalf(t, "B", user.FirstName(), "first_name")
			assert.Equalf(t, "C", user.LastName(), "last_name")
			assert.Equalf(t, "en", user.Language(), "language")
		})
	})
}
