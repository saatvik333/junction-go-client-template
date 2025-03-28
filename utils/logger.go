package utils

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger initializes the global logger
func InitLogger() {
	if Log != nil {
		return // Avoid re-initialization if already set
	}

	var err error
	Log, err = zap.NewProduction() // Use zap.NewDevelopment() for development
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// Ensure logs are properly flushed on program exit
	defer Log.Sync()
}
