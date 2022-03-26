# archstring

```go
import (
    "fmt"
    "runtime"

    "github.com/northwood-labs/golang-utils/archstring"
)

func main() {
    fmt.Println(
        archstring.GetFriendlyName(runtime.GOOS, runtime.GOARCH),
    )
}
```
