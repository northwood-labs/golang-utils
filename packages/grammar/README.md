# grammar

## Pluralize

```go
import "github.com/northwood-labs/golang-utils/grammar"

func main() {
    results := DatabaseThing()

    fmt.Printf(
        "There were %d %s.\n",
        results.NumberOfResults,
        grammar.Pluralize(results.NumberOfResults, "result", "results")
    )
}
```

Based on the count of whatever the thing is, you can pass a singular and plural version of a string so that you can maintain proper grammar in your application.

No more _“There were 1 results.”_
