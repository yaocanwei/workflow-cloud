package cmdx

import (
	"context"
	"fmt"
	"sync"
)

//TODO: error handler
type processor struct {
	handler Handler
	done    chan struct{}
}

func (p *processor) start(task *Task, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-p.done:
				return
			default:
				p.exec(task)
			}
		}
	}()
}

func (p *processor) exec(task *Task) {
	resCh := make(chan error, 1)
	go func() {
		// TODO: add context with deadline
		resCh <- perform(context.Background(), task, p.handler)
	}()
}

// perform call the handler with given task
func perform(ctx context.Context, task *Task, handler Handler) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("panic: %v", x)
		}
	}()
	return handler.ProcessTask(ctx, task)
}
