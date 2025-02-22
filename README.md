# Go Utility

### Instrumentation

#### Logging

- [x] Implement logger
- [x] Implement tracing auto associated logs

## Server usage

### Default logger

```go
import (
    ...
    log "go-utility/logging/slog"
)

func main()  {
	log.SetLogger(log.NewDefaultLogger(
		log.WithLevel(log.LEVEL_DEBUG),
		log.WithRotateFile("console.log"),
	))
	log.Debug("hello world")
	log.Info("hello world")
	log.Warn("hello world")
	log.Error("hello world")
}

```


#### view log

```log
{"time":"2024-08-27T14:27:08.284733+07:00","level":"DEBUG","msg":"hello world","trace_id":"unknown","file":"main.go:12"}
{"time":"2024-08-27T14:27:08.284895+07:00","level":"INFO","msg":"hello world","trace_id":"unknown","file":"main.go:13"}
{"time":"2024-08-27T14:27:08.284927+07:00","level":"WARN","msg":"hello world","trace_id":"unknown","file":"main.go:14"}
{"time":"2024-08-27T14:27:08.284946+07:00","level":"ERROR","msg":"hello world","trace_id":"unknown","file":"main.go:15"}
```

### Log with context

```go
import (
    ...
    log "go-utility/logging/slog"
)

func main()  {
	log.SetLogger(log.NewDefaultLogger(
		log.WithLevel(log.LEVEL_DEBUG),
		log.WithRotateFile("console.log"),
	))
	ctx := context.Background()
	ctx = log.WithTraceId(ctx, "123")

	log.DebugContext(ctx, "hello world")
	log.InfoContext(ctx, "hello world")
	log.WarnContext(ctx, "hello world")
	log.ErrorContext(ctx, "hello world")
}

```

#### view log

```log
{"time":"2024-08-27T14:37:01.098351+07:00","level":"DEBUG","msg":"hello world","trace_id":"123","file":"main.go:22"}
{"time":"2024-08-27T14:37:01.098362+07:00","level":"INFO","msg":"hello world","trace_id":"123","file":"main.go:23"}
{"time":"2024-08-27T14:37:01.098371+07:00","level":"WARN","msg":"hello world","trace_id":"123","file":"main.go:24"}
{"time":"2024-08-27T14:37:01.09838+07:00","level":"ERROR","msg":"hello world","trace_id":"123","file":"main.go:25"}
```
