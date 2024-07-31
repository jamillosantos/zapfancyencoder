# zapfancyencoder

The fancy encoder is a [zap](https://github.com/uber-go/zap) encoder that prints the logs in a colored human readable
format.

Below an example:

```
Level      : INFO
Message    : message 1
Timestamp  : 2024-07-30 21:20:37 -03:00
Fields     : 
      ├─ persistent: persisted value
      ├─ id        : 1
      └─ obj       : 
          ├─ prop1: value 1
          └─ prop2: 2
------------------------------------
Level      : WARN
Message    : message 2
Timestamp  : 2024-07-30 21:20:37 -03:00
Fields     : 
      ├─ persistent: persisted value
      ├─ id        : 123
      └─ obj       : 
          ├─ name: John
          ├─ age : 32
          └─ arr : 
              ├─ [0] : a
              ├─ [1] : b
              ├─ [0] : c1
              ├─ [1] : c2
              ├─ [3] : d
              ├─ name: John
              └─ age : 32
------------------------------------
Level      : ERROR
Message    : test error message 2
Timestamp  : 2024-07-30 21:20:37 -03:00
Fields     : 
      ├─ persistent: persisted value
      ├─ id        : 2
      └─ obj       : 
          ├─ name: John 2
          └─ age : 32
------------------------------------
Level      : ERROR
Message    : test error message 3
Timestamp  : 2024-07-30 21:20:37 -03:00
Fields     : 
      ├─ persistent: persisted value
      ├─ id        : 3
      └─ obj       : 
          ├─ name: John 3
          └─ age : 33
------------------------------------
Level      : ERROR
Message    : test error message 4
Timestamp  : 2024-07-30 21:20:37 -03:00
Fields     : 
      ├─ persistent: persisted value
      ├─ id        : 4
      └─ obj       : 
          ├─ name: John 4
          └─ age : 34
------------------------------------

```

## Usage

```go
package main

import (
    _ "github.com/jamillosantos/zapfancyencoder"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
    zapcfg := zap.NewDevelopmentConfig()
    zapcfg.Encoding = "fancy"
    logger, err := zapcfg.Build()
	if err != nil {
		panic(err)
	}
    
    logger.
        With(zap.String("field_1", "value_1")).
        Info("Hello, world!",
            zap.String("id", "123"),
            zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
                enc.AddString("name", "John")
                enc.AddInt("age", 32)
                return nil
            })),
        )
}
```