# sigctx - Signal Context

[![GoDoc](https://godoc.org/github.com/koron-go/sigctx?status.svg)](https://godoc.org/github.com/koron-go/sigctx)
[![CircleCI](https://circleci.com/gh/koron-go/sigctx.svg?style=svg)](https://circleci.com/gh/koron-go/sigctx)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/sigctx)](https://goreportcard.com/report/github.com/koron-go/sigctx)

sigctx is a context of signal listening loop.

## Example

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
