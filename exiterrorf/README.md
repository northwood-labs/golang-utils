# exiterrorf

Removed. Use error-wrapping (Go 1.13) instead.

Using <https://github.com/charmbracelet/log>, you should do:

```go
logger := log.NewWithOptions(os.Stderr, log.Options{
    ReportCaller: true,
    ReportTimestamp: true,
    TimeFormat: time.Kitchen,
})

log.With("err", err).Fatal("unable to start oven")
```
