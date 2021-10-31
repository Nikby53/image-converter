package configs

import (
	"os"
	"testing"

	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Clearenv()
	err := os.Setenv("APi_PORT", "8000")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("DB_USERNAME", "postgres")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("DB_PASSWORD", "password")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("DB_NAME", "converter")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("DB_HOST", "postgres")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("DB_PORT", "5432")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("DB_SSL_MODE", "disable")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("AWS_REGION", "eu-central-1")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("AWS_ACC_ID", "REWPOUGHEROPFGFOVNSDFGUWREHSG")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("AWS_SECRET_KEY", "ASDPOJHADPOKHJASPODJA{OSDJAPOIHDQWEIQWJEJ")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("AWS_BUCKET_NAME", "converter")
	if err != nil {
		assert.NoError(t, err)
	}
	err = os.Setenv("RABBIT_URL", "amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
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
		AWSConf: &storage.AWSConfig{
			Region:     "eu-central-1",
			AccID:      "REWPOUGHEROPFGFOVNSDFGUWREHSG",
			SecretKey:  "ASDPOJHADPOKHJASPODJA{OSDJAPOIHDQWEIQWJEJ",
			BucketName: "converter",
		},
	}
	assert.Equal(t, expected, actual)
}
