package main

import (
	"github.com/Nikby53/image-converter/internal/app"
	"github.com/Nikby53/image-converter/internal/logs"
	_ "github.com/lib/pq"
)

func main() {
	var logger = logs.NewLogger()
	logger.Infoln("Server is starting")
	err := app.Start(logger)
	if err != nil {
		logger.Fatalf("failed to start app: %v", err)
	}
}
