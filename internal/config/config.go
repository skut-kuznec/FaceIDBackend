package config

import (
	"time"

	"github.com/spf13/viper"
)

//
// api:
//   host: "0.0.0.0"
//   port: 8080
//   program_key: "ProgAuthDevKey"
//   read_time_out: 40s
//   write_time_out: 20s
//   read_head_time_out: 50s
//   release: false
//   admin_key: secret-admin-key
// log:
//   level: "debug"
//   to_file: false
//   filename: "log/epg.log"
// image:
//   upload_dir: "./upload"
//   model_dir: "./models"

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

type ConfigV2 struct {
	API      APIConfig
	Image    ImageConfig
	LogLevel string `mapstructure:"log_level"`
	Debug    bool
}

func NewConfigV2() (*ConfigV2, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cfg")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var ret ConfigV2

	if err := viper.Unmarshal(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

type Config struct {
	API
	Log
	Files
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ConfigFromViper()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) ConfigFromViper() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = c.ConfigureLog()
	if err != nil {
		return err
	}
	err = c.ConfigAAPI()
	if err != nil {
		return err
	}
	err = c.ConfigImage()
	if err != nil {
		return err
	}
	return nil
}
