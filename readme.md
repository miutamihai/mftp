# Mihai's Funky Telemetry Package - MFTP

MFTP is a golang common logging library, with configurable drivers & format.

## Example use

You can check the `some_app` project for an use example, but the TL;DR is:

```go
package main

import (
	"mihaimiuta/mftp/pkg/mftp"
	"mihaimiuta/mftp/pkg/mftp/driver"
)

func main() {
	loggerInstance := mftp.Logger{}
    loggerInstance.Initialize(mftp.InitializationInput{
		ParentContext:  context.Background(),
		DriverInstance: &driver.StandardOutputDriver{},
		ConfigPath:     "./config.json",
	})

    loggerInstance.Info("some info", nil)
}
```

One can use the following `make` commands:

* `make run` - will run the `some_app` project
* `make test` - will run all the tests in `mftp`

## Drivers

Storing the logs is done via drivers. Essentially, any struct that implements
the `driver.Driver` interface can serve as a driver. For brevity, here
is its definition:

```go
type Driver interface {
	Write(logs []types.Log, encodeLog func(types.Log) string) error
	SupportsANSIColors() bool
	GetBufferSize() int
}
```

Some more details on those methods are:

* `Write` - pretty self explanatory, takes a list of logs, applies a function
on each log & writes the result
* `SupportsANSIColors` - determines whether the driver supports ANSI colors,
which `mftp` can then add to the different log levels for an easier reading
of the logs
* `GetBufferSize` - determines the driver's preferred buffer size (i.e. how
many logs are in the `logs` list when `Write` is called).
Can be overridden via the [Config file](<readme.md#Config file>).

These drivers are already included:

* `StandardOutputDriver` - logs to the standard output
* `TextFileDriver` - logs to a text file

## Logging format

`mftp`'s logging format can be configured via the `log_format` field of the
config file. It should be a comma delimited list of the following values:

* `LEVEL` - the log level (info, error or debug)
* `TIMESTAMP` - the current timestamp
* `TRACE_ID` - the current trace id (see
[Transaction Style logs](<readme.md#Transaction Style logs>))
* `ATTRIBUTES` - a string map containing extra attributes

## Config file

Some of `mftp`'s behaviour can be overridden via a json config file. An
example config file is already provided in the `config.json` file.
Its fields' meaning is the following:

1. `log_format` - the log format to be used. Leave blank for the default format
2. `driver` - driver configs, can be omitted
    1. `override_buffer_size` - override the driver's preferred buffer size

## Transaction Style logs

Transaction Style logging (i.e. logs containing a trace id that
can be used to correlate logs happening in the same requests)
can be achieved by using the `parentContext` input of the logger.
This context can be overridden after initialization by using the `WithContext` function.
