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

// IOCloserWrap wraps any close function to convert to the io.Closer type
type IOCloserWrap struct {
	close func() error
}

// NewIOCloserWrap ctor
func NewIOCloserWrap(close func() error) io.Closer {
	return &IOCloserWrap{close: close}
}

// Close releases resources
func (closer IOCloserWrap) Close() error {
	return closer.close()
}

// MultiCloser closes all resources that should be closed
type MultiCloser struct {
	mutex             sync.Mutex
	once              sync.Once
	closers           []io.Closer
	allResourceClosed chan struct{}
}

// New creates a new multi closer
func New() *MultiCloser {
	return &MultiCloser{
		allResourceClosed: make(chan struct{}),
	}
}

// SetGlobalCloser set a new  multi closer
func SetGlobalCloser(multiCloser *MultiCloser) {
	globalCloser = multiCloser
}

// GetGlobalCloser get a global multi closer
func GetGlobalCloser() *MultiCloser {
	return globalCloser
}

// AddGlobal adds resource than need to close for global multi closer
func AddGlobal(closers ...io.Closer) {
	globalCloser.Add(closers...)
}

// CloseGlobal closes all resources registered in global multi closer
func CloseGlobal() error {
	return globalCloser.Close()
}

// WaitGlobal waits for all resource are closed in global multi closer
func WaitGlobal() {
	globalCloser.Wait()
}

// Add adds resource than need to close
func (multiCloser *MultiCloser) Add(closers ...io.Closer) {
	multiCloser.mutex.Lock()
	defer multiCloser.mutex.Unlock()

	multiCloser.closers = append(multiCloser.closers, closers...)
}

// Close all resources registered in global multi closer
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
					errsCh <- fmt.Errorf("an error occurred when closing the resource %T: %s", closer, err)
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

// Wait waits for all resource are closed
func (multiCloser *MultiCloser) Wait() {
	<-multiCloser.allResourceClosed
}
