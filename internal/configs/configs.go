package configs

import (
	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	log "github.com/sirupsen/logrus"
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
		log.Fatalf("failed to create new session %v", err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	paramsDB, err := ssmsvc.GetParameters(&ssm.GetParametersInput{
		Names:          aws.StringSlice([]string{"DB_HOST", "DB_PORT", "DB_USERNAME", "DB_PASSWORD", "DB_NAME", "DB_SSL_MODE"}),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		log.Fatalf("failed to get db parameters %v", err)
	}

	paramsStorage, err := ssmsvc.GetParameters(&ssm.GetParametersInput{
		Names:          aws.StringSlice([]string{"BUCKET_NAME", "ACC_ID", "SECRET_KEY", "REGION"}),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		log.Fatalf("failed to get storage parameters %v", err)
	}

	paramsJWT, err := ssmsvc.GetParameters(&ssm.GetParametersInput{
		Names:          aws.StringSlice([]string{"JWT_TOKEN_TTL", "JWT_SIGNING_KEY"}),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		log.Fatalf("failed to get jwt parameters %v", err)
	}

	paramMinio, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("MINIO_ENDPOINT"),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		log.Fatalf("failed to get minio parameters %v", err)
	}

	paramPort, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("API_PORT"),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		log.Fatalf("failed to get api parameters %v", err)
	}

	return &Config{
		DBConf: &DBConfig{
			Host:     *paramsDB.Parameters[0].Value,
			Port:     *paramsDB.Parameters[3].Value,
			Username: *paramsDB.Parameters[5].Value,
			Password: *paramsDB.Parameters[2].Value,
			DBName:   *paramsDB.Parameters[1].Value,
			SSLMode:  *paramsDB.Parameters[4].Value,
		},
		AWSConf: &storage.AWSConfig{
			BucketName: *paramsStorage.Parameters[1].Value,
			AccID:      *paramsStorage.Parameters[0].Value,
			SecretKey:  *paramsStorage.Parameters[3].Value,
			Region:     *paramsStorage.Parameters[2].Value,
		},
		JWTConf: &JWTConfig{
			TokenTTL:   *paramsJWT.Parameters[1].Value,
			SigningKey: *paramsJWT.Parameters[0].Value,
		},
		MinioConf: &storage.MinioConfig{
			BucketName: *paramsStorage.Parameters[1].Value,
			AccID:      *paramsStorage.Parameters[0].Value,
			SecretKey:  *paramsStorage.Parameters[3].Value,
			Region:     *paramsStorage.Parameters[2].Value,
			Endpoint:   *paramMinio.Parameter.Value,
		},
		APIPort: *paramPort.Parameter.Value,
	}
}
