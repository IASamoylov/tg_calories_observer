//go:build e2e
// +build e2e

package asserts

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/IASamoylov/tg_calories_observer/e2e/app"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
	. "github.com/onsi/gomega"
)

var UserCreatedInStorage = func(ctx context.Context, app *app.TestableApp, user dto.User) {
	sql := `select exists(select 1 from "user" where telegram_id = $1);`
	var isExist bool
	Expect(pgxscan.Get(ctx, app.Pool(), &isExist, sql, user.TelegramID())).Should(Succeed())
	Expect(isExist).Should(BeTrue(), "Пользователь #%d не найден", user.TelegramID())
}
