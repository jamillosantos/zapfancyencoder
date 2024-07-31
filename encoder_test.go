package zapfancyencoder

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TestFancyEncoder is not really testing but rendering the test for visual inspection.
// A proper test should be written to assert the output.
func TestFancyEncoder(t *testing.T) {
	b := zap.NewDevelopmentConfig()
	b.Encoding = "fancy"
	logger, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	logger = logger.With(zap.String("persistent", "persisted value"))

	logger.Info("message 1",
		zap.String("id", "1"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("prop1", "value 1")
			enc.AddInt("prop2", 2)
			return nil
		})),
	)

	logger.Warn("message 2",
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

	logger.Error("test error message 2",
		zap.String("id", "2"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("name", "John 2")
			enc.AddInt("age", 32)
			return nil
		})),
	)

	logger.Error("test error message 3",
		zap.String("id", "3"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("name", "John 3")
			enc.AddInt("age", 33)
			return nil
		})),
	)

	logger.Error("test error message 4",
		zap.String("id", "4"),
		zap.Object("obj", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("name", "John 4")
			enc.AddInt("age", 34)
			return nil
		})),
		zap.Any("extra", []byte("extra value")),
	)
}
