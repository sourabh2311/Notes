package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logger1() {
	// Below is method one
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
      "outputPaths": ["stdout", "/Users/sourabhaggarwal/Desktop/log"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
		"levelEncoder": "lowercase",
		"timeEncoder": "ISO8601",
		"lineEnding": "\n--------------------\n",
		"timeKey": "timestamp",
		"nameKey": "whatisthis",
		"callerKey": " file and line ",
		"nameEncoder": "hi",
		"callerEncoder": "this",
		"stacktraceKey": "stack trace"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("LOGGER construction succeeded")
}

func logger2() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"/Users/sourabhaggarwal/Desktop/log3", "stdout",
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func logger3() {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"/Users/sourabhaggarwal/Desktop/log3", "stdout",
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func logger4() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"some": "bar"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "CurrentTime",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,

			LineEnding: "\n----***------\n",

			StacktraceKey: "Stack Trace",
		},
	}
	logger, _ := cfg.Build()
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
}

func logger5() {
	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"some": "bar"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,

			LineEnding: "\n----------\n",

			StacktraceKey: "Stack Trace",
		},
	}
	logger, _ := cfg.Build()
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
}

func main() {
	logger1()
	logger2()
	logger3()
	fmt.Println("\n\nFollowing is very modified json encoding\n\n")
	logger4()
	fmt.Println("\n\nFollowing is very modified console encoding\n\n")
	logger5()
}
