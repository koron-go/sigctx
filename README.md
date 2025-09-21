# sigctx - Signal Context

[![GoDoc](https://godoc.org/github.com/koron-go/sigctx?status.svg)](https://godoc.org/github.com/koron-go/sigctx)
[![Actions/Go](https://github.com/koron-go/sigctx/workflows/Go/badge.svg)](https://github.com/koron-go/sigctx/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/sigctx)](https://goreportcard.com/report/github.com/koron-go/sigctx)

sigctx is a context of signal listening loop.

In Go 1.16 and later, [singal.NotifyContext][signal] can be used to achieve the same effect as sigctx.WithCancelSignal,
so it is recommended to use [signal.NotifyContext][signal] instead.

[signal]:https://pkg.go.dev/os/signal#NotifyContext

## Example

Using `WithCancelSignal()`, it is similar to `context.WithCancel()`

```golang
import (
    "os"
    "github.com/koron-go/sigctx"
)

ctx, cancel := sigctx.WithCancelSignal(ctx.Background(), os.Interrupt)
defer cancel() // call cancel() at least once.

// TODO: work with the ctx.
```

Or try `*sigctx.Sigctx` to control in detail.

```golang
import (
    "os"
    "github.com/koron-go/sigctx"
)

// create signal listening loop
sx := sigctx.New(os.Interrupt)

// start to listen signals and get its context.
// the ctx will be done when receive signals.
// if parent context be done, it terminates listening and the ctx.
ctx := sx.Start(context.Background()).Context()

// TODO: work with the ctx.

// stop to listen signals. it make the ctx done.
sx.Stop()
```
