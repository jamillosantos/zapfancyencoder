package zapfancyencoder

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestFancyEncoder(t *testing.T) {
	b := zap.NewDevelopmentConfig()
	b.Encoding = "fancy"
	logger, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Sync()

	logger = logger.With(zap.String("persistent", "value"))

	logger.Info("Hello, world!",
		zap.String("id", "123"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("name", "John")
			enc.AddInt("age", 32)
			return nil
		})),
	)
	logger.Warn("test error message",
		zap.String("id", "123"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("name", "John")
			enc.AddInt("age", 32)
			_ = enc.AddArray("arr", zapcore.ArrayMarshalerFunc(func(arr zapcore.ArrayEncoder) error {
				arr.AppendString("a")
				arr.AppendString("b")
				_ = arr.AppendArray(zapcore.ArrayMarshalerFunc(func(arr zapcore.ArrayEncoder) error {
					arr.AppendString("c1")
					arr.AppendString("c2")
					return nil
				}))
				arr.AppendString("d")
				_ = arr.AppendObject(zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
					enc.AddString("name", "John")
					enc.AddInt("age", 32)
					return nil
				}))
				return nil
			}))
			return nil
		})),
	)
	logger.Error("test error message",
		zap.String("id", "123"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("name", "John")
			enc.AddInt("age", 32)
			return nil
		})),
	)
}
