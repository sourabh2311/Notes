= zap.NewExample().Sugar()
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