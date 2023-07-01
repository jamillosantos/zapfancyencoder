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
