package main

import (
	"github.com/Nikby53/image-converter/internal/app"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	err := app.Run()
	if err != nil {
		logrus.Fatal("failed to start app: ", err)
	}
}
