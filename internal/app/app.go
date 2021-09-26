package app

import (
	"net/http"

	"github.com/Nikby53/image-converter/internal/storage"

	"github.com/Nikby53/image-converter/internal/handler"
	"github.com/Nikby53/image-converter/internal/repository"
	"github.com/Nikby53/image-converter/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() error {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
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
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.New(db)
	reposImage := repository.New(db)
	services := service.New(repos, reposImage)
	st, err := storage.New(storage.Config{
		BucketName: viper.GetString("awsS3.bucketName"),
		AccId:      viper.GetString("awsS3.accId"),
		SecretKey:  viper.GetString("awsS3.secretKey"),
		Region:     viper.GetString("awsS3.region"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize awsS3 storage: %s", err.Error())
	}
	srv := handler.NewServer(services, st)

	if err := http.ListenAndServe(viper.GetString("port"), srv); err != nil && err != http.ErrServerClosed {
		logrus.Printf("ListenAndServe(): %s", err)
	}
	return nil
}

func initConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
