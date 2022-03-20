package config

import (
	"sabariram.com/goserverbase/config"
	"sabariram.com/goserverbase/utils"
)

type Config struct {
	Logger *config.LoggerConfig
	App    *config.ServerConfig
	Mongo  *config.MongoConfig
	KMS    *config.AWSConfig
	SNS    *config.AWSConfig
}

func NewConfig() *Config {
	return &Config{

		Logger: &config.LoggerConfig{
			Version:     utils.GetEnv("LOG_VERSION", "1.1"),
			Host:        utils.GetEnv("HOST", "localhost"),
			ServiceName: utils.GetEnv("SERVICE_NAME", "API"),
			LogLevel:    utils.GetEnvInt("LOG_LEVEL", 6),
			BufferSize:  utils.GetEnvInt("LOG_BUFFER_SIZE", 1),
			GrayLog: &config.GraylogConfig{
				URL:     utils.GetEnv("GRAYLOG_URL", "http://localhost:12201/gelf"),
				Address: utils.GetEnv("GRAYLOG_ADD", "localhost"),
				Port:    uint(utils.GetEnvInt("GRAYLOG_PORT", 12201)),
			},
			AuthHeaderKeyList: utils.GetEnvAsSlice("AUTH_HEADER_LIST", []string{}, ";"),
		},
		App: &config.ServerConfig{
			Host:        utils.GetEnv("HOST", "localhost"),
			Port:        utils.GetEnv("APP_PORT", "8080"),
			ServiceName: utils.GetEnv("SERVICE_NAME", "API"),
		},
		Mongo: &config.MongoConfig{
			ConnectionString: utils.GetEnv("MONGO_URL", ""),
		},

		KMS: &config.AWSConfig{
			Arn: utils.GetEnv("KMS_ARN", ""),
		},
		SNS: &config.AWSConfig{
			Arn: utils.GetEnv("SNS_ARN", ""),
		},
	}
}
