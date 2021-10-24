package configs

import (
	"os"
)

// Config struct contains all configs.
type Config struct {
	DBConf       *DBConfig
	APIPort      string
	AWSConf      *AWSConfig
	RabbitMQConf *RabbitMQConfig
}

// DBConfig is config of the database.
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// AWSConfig is config for aws s3 storage.
type AWSConfig struct {
	BucketName string
	AccID      string
	SecretKey  string
	Region     string
}

// RabbitMQConfig is config for rabbitMQ message broker.
type RabbitMQConfig struct {
	RabbitURL string
}

// NewConfig is constructor for Config that sets up all configs.
func NewConfig() *Config {
	return &Config{
		DBConf: &DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
		AWSConf: &AWSConfig{
			BucketName: os.Getenv("AWS_BUCKET_NAME"),
			AccID:      os.Getenv("AWS_ACC_ID"),
			SecretKey:  os.Getenv("AWS_SECRET_KEY"),
			Region:     os.Getenv("AWS_REGION"),
		},
		APIPort: os.Getenv("API_PORT"),
		RabbitMQConf: &RabbitMQConfig{
			RabbitURL: os.Getenv("RABBIT_URL"),
		},
	}
}
