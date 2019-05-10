package main

import (
	"encoding/json"

	"go.uber.org/zap"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// defer sugar.Sync()
	// sugar.Infow("failed to fetch URL",
	// 	"url", "http://example.com",
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("failed to fetch URL: %s", "http://example.com")
	// // By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering, calling Sync before letting your process exit is a good habit.
	// logger := zap.NewExample()
	// defer logger.Sync()
	// logger.Info("failed to fetch URL",
	// 	zap.String("url", "http://example.com"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
	// "outputPaths": ["stdout", "/Users/sourabhaggarwal/Desktop"],
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
    "outputPaths": ["stdout", "/Users/sourabhaggarwal/Desktop/log"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
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

	logger.Info("logger construction succeeded")
	// cfg := zap.NewProductionConfig()
	// cfg.OutputPaths = []string{
	// 	"Desktop/",
	// }
	// loggern, err := cfg.Build()
	// if err != nil {
	// 	log.Fatalf("can't initialize zap logger: %v", err)
	// }
	// defer loggern.Sync()
	// loggern.Info("failed to fetch URL",
	// 	zap.String("url", "http://example.com"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
}
