# sigctx - Signal Context

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
ctx := sx.Start().Context()

// TODO: work with the ctx.

// stop to listen signals. it make the ctx done.
sx.Stop()
```
