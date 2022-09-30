# Simple logging overlay for Arcalot

This is a very simple API abstracting away other logging libraries, such as Go log and the Go test logger. You
can use it by running:

```
go get go.flow.arcalot.io/log
```

## Why create a custom logging library?

> Why create a custom logging library, when there are a ton of them already out there?

The answer is simple: to create a clean API we can rely on. Even the `log` and `testing` packages in Go don't implement a common API for logging, not to mention the myriad libraries out there. We don't proscribe what library an application should use, but we do wish to standardize how we communicate with these libraries.

## Using the logger

The easiest way to create a logger is to use the [`Config`](config.go) struct:

```go
package yourapplication

import "go.flow.arcalot.io/log"

func main() {
    logConfig := log.Config{
        Level: log.LevelInfo,
        Destination: log.DestinationStdout,
        Stdout: os.Stdout,
    }
    logger := log.New(logConfig)
    logger.Infof("Hello world!")
}
```

You should now see a log message on your console. If you wish to log from a test utility, you can do so like this:

```go
package yourapplication_test

import "go.flow.arcalot.io/log"

func TestSomething(t *testing.T) {
    logConfig := log.Config{
        Level: log.LevelInfo,
        Destination: log.DestinationTest,
        T: t,
    }
    logger := log.New(logConfig)
    logger.Infof("Hello world!")
}
```

You can also apply labels to the log messages. This will add metadata to your log messages. Extending the previous example:

```go
package yourapplication_test

import "go.flow.arcalot.io/log"

func TestSomething(t *testing.T) {
    logConfig := log.Config{
        Level: log.LevelInfo,
        Destination: log.DestinationTest,
        T: t,
    }
    logger := log.New(logConfig)
    logger = logger.WithLabel("source", "mytest")
    logger.Infof("Hello world!")
}
```

For a more detailed explanation please see [our documentation](https://godoc.org/go.flow.arcalot.io/log).
