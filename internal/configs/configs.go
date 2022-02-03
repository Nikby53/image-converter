package configs

import (
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/Nikby53/image-converter/internal/storage"
)

// Config struct contains all configs.
type Config struct {
	DBConf    *DBConfig
	APIPort   string
	AWSConf   *storage.AWSConfig
	JWTConf   *JWTConfig
	MinioConf *storage.MinioConfig
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

// JWTConfig is config of the jwt.
type JWTConfig struct {
	TokenTTL   string
	SigningKey string
}

// NewConfig is constructor for Config that sets up all configs.
func NewConfig() *Config {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-east-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	paramDBHost, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("DB_HOST"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramDBPort, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("DB_PORT"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramDBUser, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("DB_USERNAME"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramDBPass, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("DB_PASSWORD"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramDBName, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("DB_NAME"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramDBSSL, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("DB_SSL_MODE"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramBucket, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("BUCKET_NAME"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramAccID, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("ACC_ID"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramSecKey, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("SECRET_KEY"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramRegion, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("REGION"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramJWTTtl, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("JWT_TOKEN_TTL"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramJWTSign, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("JWT_SIGNING_KEY"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramMinio, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("MINIO_ENDPOINT"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	paramPort, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("API_PORT"),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	return &Config{
		DBConf: &DBConfig{
			Host:     *paramDBHost.Parameter.Value,
			Port:     *paramDBPort.Parameter.Value,
			Username: *paramDBUser.Parameter.Value,
			Password: *paramDBPass.Parameter.Value,
			DBName:   *paramDBName.Parameter.Value,
			SSLMode:  *paramDBSSL.Parameter.Value,
		},
		AWSConf: &storage.AWSConfig{
			BucketName: *paramBucket.Parameter.Value,
			AccID:      *paramAccID.Parameter.Value,
			SecretKey:  *paramSecKey.Parameter.Value,
			Region:     *paramRegion.Parameter.Value,
		},
		JWTConf: &JWTConfig{
			TokenTTL:   *paramJWTTtl.Parameter.Value,
			SigningKey: *paramJWTSign.Parameter.Value,
		},
		MinioConf: &storage.MinioConfig{
			BucketName: *paramBucket.Parameter.Value,
			AccID:      *paramAccID.Parameter.Value,
			SecretKey:  *paramSecKey.Parameter.Value,
			Region:     *paramRegion.Parameter.Value,
			Endpoint:   *paramMinio.Parameter.Value,
		},
		APIPort: *paramPort.Parameter.Value,
	}
}
