package graceful

import (
	"context"
	"syscall"
	"testing"

	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer/mocks"
	"go.uber.org/mock/gomock"
)

func TestShutdown(t *testing.T) {
	t.Parallel()

	t.Run("without the transmitted signals, graceful shutdown will not work", func(t *testing.T) {
		closer := mocks.NewMockCloser(gomock.NewController(t))
		closer.EXPECT().Close().Times(0)

		Shutdown(context.Background(), closer)
	})

	t.Run("closes all resources after cancelling context", func(t *testing.T) {
		t.Parallel()

		closer := mocks.NewMockCloser(gomock.NewController(t))
		closer.EXPECT().Close().Times(1)

		multiCloser := multicloser.New()
		multiCloser.Add(closer)
		ctx, cancel := context.WithCancel(context.Background())

		Shutdown(ctx, multiCloser, syscall.SIGINT)

		cancel()
		multiCloser.Wait()
	})
}
