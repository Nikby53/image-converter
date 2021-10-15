package app

import (
	"context"

	"github.com/Nikby53/image-converter/internal/rabbitMQ"
	"github.com/joho/godotenv"

	"github.com/Nikby53/image-converter/internal/configs"

	"github.com/Nikby53/image-converter/internal/logs"

	"github.com/Nikby53/image-converter/internal/handler"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
)

// Start starts the server.
func Start() error {
	var logger = logs.NewLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	conf := configs.NewConfig()
	db, err := repository.NewPostgresDB(conf.DBConf)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	logger.Infoln("connected to db")
	repos := repository.New(db)
	reposImage := repository.New(db)
	services := service.New(repos, reposImage)
	st, err := storage.New(conf.AWSConf)
	if err != nil {
		logger.Fatalf("failed to initialize awsS3 storage: %s", err.Error())
	}
	logger.Infoln("connected to awsS3 storage")
	broker, err := rabbitMQ.NewRabbitMQ(conf.RabbitMQConf)
	if err != nil {
		logger.Fatalf("failed to initialize rabbitMQ message broker: %s", err.Error())
	}
	logger.Infoln("connected to rabbitMQ")
	srv := handler.NewServer(services, st, broker)
	if err := srv.Run(conf.APIPort, srv); err != nil {
		logger.Fatalf("error occurred while running http server: %s", err.Error())
	}
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	return nil
}
