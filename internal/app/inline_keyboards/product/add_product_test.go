package product

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestAddProductInlineButton(t *testing.T) {
	t.Parallel()

	mockCmd := NewMockcommand(gomock.NewController(t))
	mockCmd.EXPECT().Alias().Return("/add_product")
	btn := NewAddProductInlineButton(mockCmd)

	t.Run("возврщается название кнопки", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "➕ Добавить", btn.Text())
	})

	t.Run("возврщается комманду действие, которую необходимо выполнить", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "/add_product", btn.Callback())
	})
}
