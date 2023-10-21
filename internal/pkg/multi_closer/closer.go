package multicloser

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks io Closer

var globalCloser *MultiCloser

func init() {
	SetGlobalCloser(New())
}

// CloserFunc ...
type CloserFunc func() error

// IOCloserWrap врапер, позволяющие привести функцию вида CloserFunc к io.Closer
type IOCloserWrap struct {
	close CloserFunc
}

// NewIOCloserWrap ctor
func NewIOCloserWrap(close CloserFunc) io.Closer {
	return &IOCloserWrap{close: close}
}

// Close закрывает зарегистрированый ресурс
func (closer IOCloserWrap) Close() error {
	return closer.close()
}

// MultiCloser освобождает все ресурсы, которые должны быть закрыты
type MultiCloser struct {
	mutex             sync.Mutex
	once              sync.Once
	closers           []io.Closer
	allResourceClosed chan struct{}
}

// New ctor
func New() *MultiCloser {
	return &MultiCloser{
		allResourceClosed: make(chan struct{}),
	}
}

// SetGlobalCloser устанвливает новый комплексный обработчик освобожения ресурсов
func SetGlobalCloser(multiCloser *MultiCloser) {
	globalCloser = multiCloser
}

// GetGlobalCloser возвращает текущий комплексный обработчик освобожения ресурсов
func GetGlobalCloser() *MultiCloser {
	return globalCloser
}

// AddGlobal добавляет ресурс, который должен быть закрыт в глобальный обработчик
func AddGlobal(closers ...io.Closer) {
	globalCloser.Add(closers...)
}

// CloseGlobal завершает работу всех ресурсов зарегистрированных ранее
func CloseGlobal() error {
	return globalCloser.Close()
}

// WaitGlobal ожидает пока все ресурсы не будут закрыты
func WaitGlobal() {
	globalCloser.Wait()
}

// Add добавляет ресурс, который должен быть закрыт
func (multiCloser *MultiCloser) Add(closers ...io.Closer) {
	multiCloser.mutex.Lock()
	defer multiCloser.mutex.Unlock()

	multiCloser.closers = append(multiCloser.closers, closers...)
}

// Close завершает работу всех ресурсов зарегистрированных ранее
func (multiCloser *MultiCloser) Close() (err error) {
	multiCloser.once.Do(func() {
		defer close(multiCloser.allResourceClosed)

		multiCloser.mutex.Lock()
		closers := multiCloser.closers
		multiCloser.closers = nil
		multiCloser.mutex.Unlock()

		errsCh := make(chan error, len(closers))
		defer close(errsCh)

		for _, closer := range closers {
			closer := closer
			go func() {
				if err := closer.Close(); err != nil {
					errsCh <- fmt.Errorf("произошла ошибка в момент завершения работы ресурса %T: %s", closer, err)
				} else {
					errsCh <- nil
				}
			}()
		}

		var errs []error
		for i := 0; i < len(closers); i++ {
			errs = append(errs, <-errsCh)
		}

		err = errors.Join(errs...)
	})

	return err
}

// Wait ожидает пока все ресурсы не будут закрыты
func (multiCloser *MultiCloser) Wait() {
	<-multiCloser.allResourceClosed
}
