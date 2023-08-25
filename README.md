# zapfancyencoder

The fancy encoder is a [zap](https://github.com/uber-go/zap) encoder that prints the logs in a colored human readable
format.

Below an example:

```
Level      : INFO
Message    : Hello, world!
Timestamp  : 2023-08-24 22:20:29 -03:00
Fields     : 
      ├─ persistent: value
      ├─ id        : 123
      └─ obj       : 
          ├─ name: John
          └─ age : 32
------------------------------------
Level      : WARN
Message    : test error message
Timestamp  : 2023-08-24 22:20:29 -03:00
Fields     : 
      ├─ persistent: value
      ├─ id        : 123
      ├─ obj       : 
      |   ├─ name: John
      |   └─ age : 32
      ├─ id        : 123
      └─ obj       : 
          ├─ name: John
          ├─ age : 32
          └─ arr : 
              ├─ [0]: a
              ├─ [1]: b
              ├─ [2]: 
              |   ├─ [0]: c1
              |   └─ [1]: c2
              ├─ [3]: d
              └─ [4]: 
                  ├─ name: John
                  └─ age : 32
------------------------------------
Level      : ERROR
Message    : test error message
Timestamp  : 2023-08-24 22:20:29 -03:00
Fields     : 
      ├─ persistent: value
      ├─ id        : 123
      ├─ obj       : 
      |   ├─ name: John
      |   └─ age : 32
      ├─ id        : 123
      ├─ obj       : 
      |   ├─ name: John
      |   ├─ age : 32
      |   └─ arr : 
      |       ├─ [0]: a
      |       ├─ [1]: b
      |       ├─ [2]: 
      |       |   ├─ [0]: c1
      |       |   └─ [1]: c2
      |       ├─ [3]: d
      |       └─ [4]: 
      |           ├─ name: John
      |           └─ age : 32
      ├─ id        : 123
      └─ obj       : 
          ├─ name: John
          └─ age : 32
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