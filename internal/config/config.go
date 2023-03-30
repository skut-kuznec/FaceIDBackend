package config

import (
	"github.com/spf13/viper"
)

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
