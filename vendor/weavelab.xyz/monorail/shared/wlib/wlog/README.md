# Example

```go

import (
  "weavelab.xyz/monorail/shared/wlib/werror"
  "weavelab.xyz/monorail/shared/wlib/wlog"
  "weavelab.xyz/monorail/shared/wlib/wlog/tag"
)

func main() {
  ctx := context.Background()

  wlog.InfoC(ctx, "Hello!", tag.Int("key", 5))
  wlog.DebugC(ctx, "Only when debugging enabled!", tag.String("key", "value"))

  err := werror.New("disconnected").Add("when", "yesterday")
  wlog.WErrorC(ctx, err)
}

```
