# aclog

App Container JSON formatted logging built on Uber's Zap.

aclog is intended for Go applications designed to run specifically within container runtimes. The aim is to provide a logging facility for introspection of the executing container and its runtime.

aclog utilizes Uber's [Zap](https://github.com/uber-go/zap) logging package for performance.

## Installation
`go get -u github.com/christianvozar/aclog`

## Usage
```go
import (
  _ "github.com/christianvozar/aclog"
)
```

