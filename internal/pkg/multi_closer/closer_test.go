package multicloser

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer/mocks"
)

// nolint
func TestAdd(t *testing.T) {
	t.Run("adds multiple closers", func(t *testing.T) {
		SetGlobalCloser(New())

		closer1 := mocks.NewMockCloser(gomock.NewController(t))
		closer2 := mocks.NewMockCloser(gomock.NewController(t))
		closer3 := mocks.NewMockCloser(gomock.NewController(t))

		AddGlobal(closer1, closer2)
		AddGlobal(closer3)

		assert.ElementsMatch(t, GetGlobalCloser().closers, []io.Closer{closer1, closer2, closer3})
	})

	// nolint
	t.Run("adding to the unique multi closer does not change the global", func(t *testing.T) {
		SetGlobalCloser(New())

		uniqueCloser := New()

		closer1 := mocks.NewMockCloser(gomock.NewController(t))
		closer2 := mocks.NewMockCloser(gomock.NewController(t))
		closer3 := mocks.NewMockCloser(gomock.NewController(t))

		uniqueCloser.Add(closer1, closer2)
		AddGlobal(closer3)

		assert.ElementsMatch(t, uniqueCloser.closers, []io.Closer{closer1, closer2})
		assert.ElementsMatch(t, GetGlobalCloser().closers, []io.Closer{closer3})
	})
}

// nolint
func TestClose(t *testing.T) {
	t.Run("closing a global closer does not affect the unique", func(t *testing.T) {
		SetGlobalCloser(New())
		uniqueCloser := New()

		closer1 := mocks.NewMockCloser(gomock.NewController(t))
		closer2 := mocks.NewMockCloser(gomock.NewController(t))
		closer3 := mocks.NewMockCloser(gomock.NewController(t))

		closer1.EXPECT().Close().DoAndReturn(func() error {
			time.Sleep(300 * time.Millisecond)

			return nil
		}).Times(1)
		closer2.EXPECT().Close().Times(0)
		closer3.EXPECT().Close().DoAndReturn(func() error {
			time.Sleep(300 * time.Millisecond)

			return nil
		}).Times(1)

		uniqueCloser.Add(closer2)
		AddGlobal(closer1, closer3)

		err := CloseGlobal()
		assert.NoError(t, err)
		WaitGlobal()
	})

	t.Run("an error that occurred when closing one resource does not affect the closure of the rest", func(t *testing.T) {
		uniqueCloser := New()

		closer1 := mocks.NewMockCloser(gomock.NewController(t))
		closer2 := mocks.NewMockCloser(gomock.NewController(t))
		closer3 := mocks.NewMockCloser(gomock.NewController(t))

		closer1.EXPECT().Close().Return(nil).Times(1)
		closer2.EXPECT().Close().Return(fmt.Errorf("test error")).Times(1)
		closer3.EXPECT().Close().Return(nil).Times(1)

		uniqueCloser.Add(closer1, closer2, closer3)

		err := uniqueCloser.Close()
		assert.EqualError(t, err, "an error occurred when closing the resource *mocks.MockCloser: test error")
		WaitGlobal()
	})
}
