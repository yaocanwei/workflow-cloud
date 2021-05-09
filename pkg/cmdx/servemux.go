package cmdx

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type ServeMux struct {
	mu sync.RWMutex
	m  map[string]muxEntry
	es []muxEntry
}

type muxEntry struct {
	h       Handler
	pattern string
}

func NewServeMux() *ServeMux {
	return new(ServeMux)
}

func (mux *ServeMux) ProcessTask(ctx context.Context, t *Task) error {
	h, _ := mux.Handler(t)
	return h.ProcessTask(ctx, t)
}

func (mux *ServeMux) Handler(t *Task) (h Handler, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	h, pattern = mux.match(t.Type)
	if h == nil {
		h, pattern = NotFoundHandler(), ""
	}
	//for i := len(mux.mws) - 1; i >= 0; i-- {
	//	h = mux.mws[i](h)
	//}
	return h, pattern
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("asynq: invalid pattern")
	}
	if handler == nil {
		panic("asynq: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("asynq: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	e := muxEntry{h: handler, pattern: pattern}
	mux.m[pattern] = e
	mux.es = appendSorted(mux.es, e)
}

func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
	n := len(es)
	i := sort.Search(n, func(i int) bool {
		return len(es[i].pattern) < len(e.pattern)
	})
	if i == n {
		return append(es, e)
	}
	es = append(es, muxEntry{})
	copy(es[i+1:], es[i:])
	es[i] = e
	return es
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(context.Context, *Task) error) {
	if handler == nil {
		panic("asynq: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}

func (mux *ServeMux) match(t string) (Handler, string) {
	v, ok := mux.m[t]
	if ok {
		return v.h, v.pattern
	}

	for _, e := range mux.es {
		if strings.HasPrefix(t, e.pattern) {
			return e.h, e.pattern
		}
	}
	return nil, ""
}

func NotFound(ctx context.Context, task *Task) error {
	return fmt.Errorf("handler not found for task %q", task.Type)
}

// NotFoundHandler returns a simple task handler that returns a ``not found`` error.
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
