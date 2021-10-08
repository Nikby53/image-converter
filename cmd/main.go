package main

import (
	"github.com/Nikby53/image-converter/internal/app"
	"github.com/Nikby53/image-converter/internal/logs"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	var logging = logs.NewLogger()
	logging.Info("Server is starting")
	err := app.Start()
	if err != nil {
		logrus.Fatal("failed to start app: ", err)
	}
}
