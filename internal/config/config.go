package config

import (
	"time"

	"github.com/spf13/viper"
)

type APIConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	SecretKey    string        `mapstructure:"secret_key"`
}

type ImageConfig struct {
	UploadDir string `mapstructure:"upload_dir"`
	ModelDir  string `mapstructure:"model_dir"`
}

type Config struct {
	API      APIConfig
	Image    ImageConfig
	LogLevel string `mapstructure:"log_level"`
	Debug    bool
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cfg")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var ret Config

	if err := viper.Unmarshal(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
