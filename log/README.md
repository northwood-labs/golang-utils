# log

## GetStdTextLogger

Returns an instantiated <https://github.com/rs/zerolog> object.

```go
import "github.com/northwood-labs/golang-utils/log"

func main() {
    logger := log.GetStdTextLogger()
    log.Info().Str("foo", "bar").Msg("Hello World")

    // Output: 2006-01-02T15:04:05Z07:00 | INFO  | Hello World foo=bar
}
```

## GetStdJSONLogger

Returns an instantiated <https://github.com/rs/zerolog> object.

```go
import "github.com/northwood-labs/golang-utils/log"

func main() {
    logger := log.GetStdJSONLogger()

    err := errors.New("A repo man spends his life getting into tense situations")
    service := "myservice"

    log.Fatal().
        Err(err).
        Str("service", service).
        Msgf("Cannot start %s", service)

    // Output: {"time":1516133263,"level":"fatal","error":"A repo man spends his life getting into tense situations","service":"myservice","message":"Cannot start myservice"}
    //         exit status 1
}
```
