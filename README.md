# Shared Resources for Northwood Labs

This is only a library. It does not work on its own. It is consumed by the various apps.

## Usage

### Utilities

#### ExitErrorf

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

#### Pluralize

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

#### GetSpew

```go
import (
    "github.com/northwood-labs/golang-utils/debug"
    "github.com/northwood-labs/golang-utils/exiterrorf"
)

func main() {
    var results map[string]interface{}

    // Do some stuff to a variable.
    err := json.Unmarshal(data, &results)
    if err != nil {
        exiterrorf.ExitErrorf(err)
    }

    // Pretty-print the contents of the variable.
    pp := debug.GetSpew()
    pp.Dump(results)
}
```

Spew is a library that can print the contents of a data structure, much like `PrettyPrint` in Python, or `print_r()` or `var_dump()` in PHP. This implementation provides a specific configuration for Spew that is repeatable.
