# Build

Provides the application build metadata

## Usage

Print the application build meta

```go
build.Print()
```

Return the application build meta

```go
engine.GET("/v1/metadata", build.Handle)
engine.GET("/v1/headers", build.Header)
```

