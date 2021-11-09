package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Nikby53/image-converter/internal/configs"
	"github.com/Nikby53/image-converter/internal/handler"
	"github.com/Nikby53/image-converter/internal/logs"
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
	logger.Infoln("connected to storage")
	services := service.New(repo, st)
	srv := handler.NewServer(services, st)
	go func() {
		if err := srv.Run(conf.APIPort, srv); err != nil && err != http.ErrServerClosed {
			logger.Errorf("error occurred while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Infoln("Server shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	logger.Infoln("Server stopped")
	if err := db.Close(); err != nil {
		logger.Errorf("error occurred on db connection close: %s", err.Error())
	}
	logger.Infoln("Database stopped")
	return nil
}
