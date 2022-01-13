package configs

import (
	"os"
	"testing"

	"github.com/Nikby53/image-converter/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Clearenv()

	err := os.Setenv("API_PORT", "8000")
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

	err = os.Setenv("MINIO_REGION", "eu-central-1")
	if err != nil {
		assert.NoError(t, err)
	}

	err = os.Setenv("MINIO_ACC_ID", "REWPOUGHEROPFGFOVNSDFGUWREHSG")
	if err != nil {
		assert.NoError(t, err)
	}

	err = os.Setenv("MINIO_SECRET_KEY", "ASDPOJHADPOKHJASPODJA{OSDJAPOIHDQWEIQWJEJ")
	if err != nil {
		assert.NoError(t, err)
	}

	err = os.Setenv("MINIO_BUCKET_NAME", "converter")
	if err != nil {
		assert.NoError(t, err)
	}

	err = os.Setenv("MINIO_ENDPOINT", "localhost")
	if err != nil {
		assert.NoError(t, err)
	}

	err = os.Setenv("JWT_SIGNING_KEY", "qwetdfrydfgyesrafxzf")
	if err != nil {
		assert.NoError(t, err)
	}

	err = os.Setenv("JWT_TOKEN_TTL", "1h")
	if err != nil {
		assert.NoError(t, err)
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
		MinioConf: &storage.MinioConfig{
			BucketName: "converter",
			AccID:      "REWPOUGHEROPFGFOVNSDFGUWREHSG",
			SecretKey:  "ASDPOJHADPOKHJASPODJA{OSDJAPOIHDQWEIQWJEJ",
			Region:     "eu-central-1",
			Endpoint:   "localhost",
		},
		JWTConf: &JWTConfig{
			TokenTTL:   "1h",
			SigningKey: "qwetdfrydfgyesrafxzf",
		},
	}
	assert.Equal(t, expected, actual)
}
