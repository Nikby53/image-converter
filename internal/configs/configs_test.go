package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {

	os.Clearenv()

	os.Setenv("APi_PORT", "8000")

	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "converter")
	os.Setenv("DB_HOST", "postgres")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SSL_MODE", "disable")

	os.Setenv("AWS_REGION", "eu-central-1")
	os.Setenv("AWS_ACC_ID", "REWPOUGHEROPFGFOVNSDFGUWREHSG")
	os.Setenv("AWS_SECRET_KEY", "ASDPOJHADPOKHJASPODJA{OSDJAPOIHDQWEIQWJEJ")
	os.Setenv("AWS_BUCKET_NAME", "converter")

	os.Setenv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")

	actual := NewConfig()
	expected := &Config{
		APIPort: "8000",
		DBConf: &DBConfig{
			Username: "postgres",
			Password: "password",
			DBName:   "converter",
			Host:     "postgres",
			Port:     "5432",
			SSLMode:  "disable",
		},
		AWSConf: &AWSConfig{
			Region:     "eu-central-1",
			AccID:      "REWPOUGHEROPFGFOVNSDFGUWREHSG",
			SecretKey:  "ASDPOJHADPOKHJASPODJA{OSDJAPOIHDQWEIQWJEJ",
			BucketName: "converter",
		},
		RabbitMQConf: &RabbitMQConfig{
			RabbitURL: "amqp://guest:guest@localhost:5672/",
		},
	}
	assert.Equal(t, expected, actual)
}
