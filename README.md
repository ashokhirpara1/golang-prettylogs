# Go Prettylogs
Golang structured logging library

## Usage

### Get the go-prettylogs library

```
$ go get github.com/ashokhirpara1/golang-prettylogs
```

```go
package main

import (
    prettylogs "github.com/ashokhirpara1/golang-prettylogs"
)

func main() {
    logs.Info("main function to boot up everything")
}
```

### Expected output

```
{"level":"info","method":"main","msg":"main function to boot up everything","time":"2022-09-23T22:59:01+05:30"}
```

### Measure function's execution time in logs
```
func functionName() {
    defer prettylogs.Exit(prettylogs.Enter())
}
```