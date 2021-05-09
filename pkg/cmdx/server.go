package cmdx

import (
	"context"
	"fmt"
	"sync"
)

type Server struct {
	wg        sync.WaitGroup
	processor *processor
}

type Handler interface {
	ProcessTask(context.Context, *Task) error
}

type HandlerFunc func(context.Context, *Task) error

func (h HandlerFunc) ProcessTask(ctx context.Context, task *Task) error {
	return h(ctx, task)
}

func (srv *Server) Run(task *Task, handler Handler) error {
	if err := srv.Start(task, handler); err != nil {
		return err
	}
	return nil
}

func (srv *Server) Start(task *Task, handler Handler) error {
	if handler == nil {
		return fmt.Errorf("cmdx: server cannot run nil handler")
	}
	srv.processor.handler = handler
	srv.processor.start(task, &srv.wg)
	return nil
}
