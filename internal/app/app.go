package app

import (
	"context"

	"github.com/Nikby53/image-converter/internal/logs"

	"github.com/Nikby53/image-converter/internal/handler"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/Nikby53/image-converter/internal/service"
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/spf13/viper"
)

// Start starts the server.
func Start() error {
	var logger = logs.NewLogger()
	if err := initConfig(); err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	logger.Info("connected to db")
	repos := repository.New(db)
	reposImage := repository.New(db)
	services := service.New(repos, reposImage)
	st, err := storage.New(storage.Config{
		BucketName: viper.GetString("awsS3.bucketName"),
		AccID:      viper.GetString("awsS3.accId"),
		SecretKey:  viper.GetString("awsS3.secretKey"),
		Region:     viper.GetString("awsS3.region"),
	})
	if err != nil {
		logger.Fatalf("failed to initialize awsS3 storage: %s", err.Error())
	}
	srv := handler.NewServer(services, st)
	if err := srv.Run(viper.GetString("port"), srv); err != nil {
		logger.Fatalf("error occurred while running http server: %s", err.Error())
	}
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	return nil
}

func initConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
