// Golang SwaggerUI image-service
//
// Documentation of our awesome API
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package main

import (
	"github.com/Nikby53/image-converter/internal/app"
	"github.com/Nikby53/image-converter/internal/logs"
	_ "github.com/lib/pq"
)

func main() {
	var logger = logs.NewLogger()
	logger.Infoln("Server is starting")
	err := app.Start()
	if err != nil {
		logger.Fatalf("failed to start app: %v", err)
	}
}
