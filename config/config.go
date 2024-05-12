package config

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kelseyhightower/envconfig"
	"log"
	"memoo-backend/constants"
	"os"
)

var AppConfig *Config

// Config represents the composition of yml settings.
type Config struct {
	Database struct {
		Dsn string `envconfig:"PG_DSN" default:""`
	}
	JwtAttribute struct {
		SignValidTimeMs int64 `yaml:"sign_valid_time_ms" envconfig:"SIGN_VALID_TIME_MS" default:"30000"`
	}

	AwsAttribute struct {
		AwsRegion          string `yaml:"aws_region" envconfig:"AWS_REGION" default:""`
		AwsAccessKeyId     string `yaml:"aws_access_key_id" envconfig:"AWS_ACCESS_KEY_ID" default:""`
		AwsSecretAccessKey string `yaml:"aws_secret_access_key" envconfig:"AWS_SECRET_ACCESS_KEY" default:""`
		AwsBucketName      string `yaml:"aws_bucket_name" envconfig:"AWS_BUCKET_NAME" default:""`
		AwsExpirationTime  int64  `yaml:"aws_expiration_time" envconfig:"AWS_EXPIRATION_TIME" default:"168"`
	}
	Extension struct {
		MasterGenerator bool   `yaml:"master_generator" default:"false"`
		CorsEnabled     bool   `yaml:"cors_enabled" default:"false"`
		SecurityEnabled bool   `yaml:"security_enabled" default:"false"`
		GinMode         string `yaml:"gin_mode" envconfig:"GIN_MODE" default:"test"`
		ServerPod       string `yaml:"server_pod" envconfig:"SERVER_POD" default:"defaultPod"`
	}
	Log struct {
		RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status}"`
	}
}
type JwtConfig struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

// LoadAppConfig reads the settings written to the yml file
func LoadAppConfig() *Config {

	//file, err := yamlFile.ReadFile(constants.AppConfigPath)
	//if err != nil {
	//	fmt.Printf("Failed to read application.yml: %s", err)
	//	os.Exit(constants.ErrExitStatus)
	//}
	config := &Config{}

	err1 := envconfig.Process("", config)
	if err1 != nil {
		log.Fatal("Failed to process environment variables:", err1)
	}

	//if err := yaml.Unmarshal(file, config); err != nil {
	//	fmt.Printf("Failed to read application.yml: %s", err)
	//	os.Exit(constants.ErrExitStatus)
	//}
	AppConfig = config
	return config
}

func loadPublicKey() (interface{}, error) {
	publicKeyFile, err := os.ReadFile(constants.JwtPublicKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func loadPrivateKey() (interface{}, error) {
	privateKeyFile, err := os.ReadFile(constants.JwtPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
