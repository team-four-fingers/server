package grouper

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"sync"
)

type Grouper interface {
	Go(f func() error)
	Wait() error
}

type GroupFunc func() Grouper

type panicSafeGroup struct {
	mutex sync.Mutex
	err   *multierror.Error
	wg    sync.WaitGroup
}

var _ Grouper = (*panicSafeGroup)(nil)

func NewPanicSafeGroup() Grouper {
	return &panicSafeGroup{}
}

func (g *panicSafeGroup) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()
		defer func() {
			var err error
			if x := recover(); x != nil {
				switch x := x.(type) {
				case error:
					err = x
				default:
					err = errors.Errorf("%s", x)
				}
			}
			if err != nil {
				g.mutex.Lock()
				g.err = multierror.Append(g.err, errors.Wrap(err, "recovered"))
				g.mutex.Unlock()
			}
		}()

		if err := f(); err != nil {
			g.mutex.Lock()
			g.err = multierror.Append(g.err, err)
			g.mutex.Unlock()
		}
	}()
}

func (g *panicSafeGroup) Wait() error {
	g.wg.Wait()
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.err.ErrorOrNil()
}
