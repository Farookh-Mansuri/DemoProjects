package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger1, _ := zap.NewDevelopment()
	// defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "blabla",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "blabla")

	defer logger1.Sync()
	logger1.Info("Hello Zap!")
	logger1.Warn("Beware of getting Zapped! (Pun)")
	logger1.Error("I'm out of Zap joke!")
}
