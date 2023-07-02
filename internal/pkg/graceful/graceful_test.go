package graceful

import (
	"syscall"
	"testing"

	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer/mocks"
	"github.com/golang/mock/gomock"
)

func TestShutdown(t *testing.T) {
	t.Parallel()

	t.Run("without the transmitted signals, graceful shutdown will not work", func(t *testing.T) {
		closer := mocks.NewMockCloser(gomock.NewController(t))
		closer.EXPECT().Close().Times(0)

		Shutdown(closer)
	})

	t.Run("closes all resources after receiving stop signal", func(t *testing.T) {
		closer := mocks.NewMockCloser(gomock.NewController(t))
		closer.EXPECT().Close().Times(1)

		multiCloser := multicloser.New()
		multiCloser.Add(closer)

		Shutdown(multiCloser, syscall.SIGINT)

		done <- syscall.SIGINT
		multiCloser.Wait()
	})
}
