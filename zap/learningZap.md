# Learning Zap

* Use Logger than SugaredLogger as it is very fast. Example below:
  ```go
  logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
  ```
* By default, loggers are unbuffered. However, since zap's low-level APIs allow buffering, calling Sync before letting your process exit is a good habit.
* The simplest way to build a Logger is to use zap's opinionated presets: NewExample, NewProduction, and NewDevelopment.
* Presets are fine for small projects, but larger projects and organizations naturally require a bit more customization. For most users, zap's Config struct strikes the right balance between flexibility and convenience.
