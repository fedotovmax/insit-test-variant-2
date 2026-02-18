package events

import (
	"context"
	"log/slog"
	"sync"

	"github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"
)

type Handler func(context.Context, *domain.ToProcessOperation) error

type Manager struct {
	log          *slog.Logger
	queue        chan *domain.ToProcessOperation
	ctx          context.Context
	cancel       context.CancelFunc
	isStopped    chan struct{}
	workersCount int
	handler      Handler
}

func New(buffer int, workersCount int, h Handler, log *slog.Logger) *Manager {
	ctx, cancel := context.WithCancel(context.Background())
	return &Manager{
		queue:        make(chan *domain.ToProcessOperation, buffer),
		isStopped:    make(chan struct{}),
		workersCount: workersCount,
		ctx:          ctx,
		cancel:       cancel,
		handler:      h,
		log:          log,
	}
}

func (m *Manager) Start() {

	wg := sync.WaitGroup{}

	for range m.workersCount {
		wg.Go(func() {
			for {
				select {
				case <-m.ctx.Done():
					return
				case value, ok := <-m.queue:
					if !ok {
						return
					}
					err := m.handler(m.ctx, value)

					if err != nil {
						m.log.Error(err.Error())
					}
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(m.isStopped)
	}()

}

func (m *Manager) Stop(ctx context.Context) error {
	m.cancel()

	select {
	case <-m.isStopped:
		close(m.queue)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}

func (m *Manager) Send(newOp *domain.ToProcessOperation) error {
	select {
	case <-m.ctx.Done():
		return m.ctx.Err()
	case m.queue <- newOp:
		return nil
	}
}
