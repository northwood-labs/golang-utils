# exiterrorf

## ExitErrorf

```go
import "github.com/northwood-labs/golang-utils/exiterrorf"

func main() {
    results, err := DoAThing()
    if err != nil {
        exiterrorf.ExitErrorf(err)
    }
}
```

It will exit the CLI tool with the error message. It will also display the **function** and **line number** from where the exit was triggered.
