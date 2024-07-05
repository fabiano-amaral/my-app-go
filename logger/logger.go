package logger

import (
	"fmt"
	"os"

	"github.com/fabiano-amaral/my-app-go/slack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func printHello() {
	fmt.Println("Hello, World!")
}

func PrintHello() {
	fmt.Println("Hello, World!")
	slack.SendMessage()
}

// InitLog initializes the logger of the application.
func InitLog() *zap.Logger {
	printHello()
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "body",
		LevelKey:    "severity",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "time",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
	}

	hostname, _ := os.Hostname()

	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"host.name":        hostname,
			"service.language": "golang",
		},
		EncoderConfig: encoderConfig,
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
