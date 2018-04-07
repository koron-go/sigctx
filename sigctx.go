package sigctx

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

// Sigctx is a context wrapper for signal handler.
type Sigctx struct {
	m     sync.Mutex
	sigs  []os.Signal
	funcs map[os.Signal]Handler
	ctx   context.Context
	cncl  context.CancelFunc
	c     chan os.Signal
}

// Handler is signal handler.
type Handler func(os.Signal)

// New creates a new Sigctx.
func New(signals ...os.Signal) *Sigctx {
	return &Sigctx{
		sigs: signals,
	}
}

// SetHandler sets a handler for a signal. Signals which have handler won't
// stop the context.
func (sx *Sigctx) SetHandler(sig os.Signal, h Handler) *Sigctx {
	sx.funcs[sig] = h
	return sx
}

// Start starts signal listening loop.
func (sx *Sigctx) Start(ctx context.Context) *Sigctx {
	sx.m.Lock()
	defer sx.m.Unlock()
	if sx.ctx != nil {
		return sx
	}
	sx.ctx, sx.cncl = context.WithCancel(ctx)
	sx.c = make(chan os.Signal, 1)
	signal.Notify(sx.c, sx.sigs...)
	go sx.loop()
	return sx
}

func (sx *Sigctx) loop() {
L:
	for {
		select {
		case <-sx.ctx.Done():
			break L
		case s := <-sx.c:
			if h, ok := sx.funcs[s]; ok {
				go h(s)
				continue
			}
			break L
		}
	}
	signal.Stop(sx.c)
	close(sx.c)
	sx.cncl()
}

// Stop stops signal listening loop.
func (sx *Sigctx) Stop() {
	sx.m.Lock()
	defer sx.m.Unlock()
	if sx.cncl == nil {
		return
	}
	sx.cncl()
}

// Context returns a context of signal listening loop
func (sx *Sigctx) Context() context.Context {
	return sx.ctx
}
