package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Nikby53/image-converter/internal/configs"
	"github.com/Nikby53/image-converter/internal/handler"
	"github.com/Nikby53/image-converter/internal/logs"
	"github.com/Nikby53/image-converter/internal/rabbitMQ"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/joho/godotenv"
)

// Start starts the server.
func Start(logger *logs.Logger) error {
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	conf := configs.NewConfig()
	db, err := repository.NewPostgresDB(conf.DBConf)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	logger.Infoln("connected to db")
	repo := repository.New(db)
	st, err := storage.New(conf.AWSConf)
	if err != nil {
		logger.Fatalf("failed to initialize awsS3 storage: %s", err.Error())
	}
	services := service.New(repo, st)
	logger.Infoln("connected to storage")
	broker, err := rabbitMQ.NewRabbitMQ(conf.RabbitMQConf)
	if err != nil {
		logger.Fatalf("failed to initialize rabbitMQ message broker: %s", err.Error())
	}
	logger.Infoln("connected to rabbitMQ")
	srv := handler.NewServer(services, st, broker)
	go func() {
		if err := srv.Run(conf.APIPort, srv); err != nil {
			logger.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Infoln("server exited properly")
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logger.Errorf("error occurred on db connection close: %s", err.Error())
	}
	return nil
}
