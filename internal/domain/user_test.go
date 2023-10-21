package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultUser(t *testing.T) {
	t.Parallel()

	t.Run("creates user without system ID", func(t *testing.T) {
		t.Parallel()

		defaultUser := NewDefaultUser(100500, "A", "B", "C", "en")

		assert.Equal(t, User{
			id:         0,
			telegramID: 100500,
			userName:   "A",
			firstName:  "B",
			lastName:   "C",
			language:   "en",
		}, defaultUser)
	})
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	t.Run("creates user with system ID", func(t *testing.T) {
		t.Parallel()

		user := NewUser(2323592345, 100500, "A", "B", "C", "en")

		assert.Equal(t, User{
			id:         2323592345,
			telegramID: 100500,
			userName:   "A",
			firstName:  "B",
			lastName:   "C",
			language:   "en",
		}, user)

		t.Run("getters return values correctly", func(t *testing.T) {
			assert.Equalf(t, int64(2323592345), user.ID(), "id")
			assert.Equalf(t, TelegramID(100500), user.TelegramID(), "telegram_id")
			assert.Equalf(t, "A", user.UserName(), "user_name")
			assert.Equalf(t, "B", user.FirstName(), "first_name")
			assert.Equalf(t, "C", user.LastName(), "last_name")
			assert.Equalf(t, "en", user.Language(), "language")
		})
	})
}
