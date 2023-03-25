package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Api struct {
	apiHost            string
	apiPort            int
	apiProgramKey      string
	apiAdminKey        string
	apiReadTimeOut     int
	apiWriteTimeOut    int
	apiReadHeadTimeOut int
}

func (c *Config) ConfigApi() error {
	c.apiHost = viper.GetString("api.host")
	if c.apiHost == "" {
		c.apiHost = "127.0.0.1"
	}
	c.apiPort = viper.GetInt("api.port")
	if c.apiPort == 0 {
		c.apiPort = 8080
	}
	c.apiProgramKey = viper.GetString("api.program_key")
	if c.apiProgramKey == "" {
		return fmt.Errorf("api.program_key not in config")
	}
	c.apiAdminKey = viper.GetString("api.admin_key")
	if c.apiAdminKey == "" {
		return fmt.Errorf("api.admin_key not in config")
	}
	c.apiReadTimeOut = viper.GetInt("api.read_time_out")
	if c.apiReadTimeOut == 0 {
		c.apiReadTimeOut = 30
	}
	c.apiWriteTimeOut = viper.GetInt("api.write_time_out")
	if c.apiWriteTimeOut == 0 {
		c.apiWriteTimeOut = 30
	}
	c.apiReadHeadTimeOut = viper.GetInt("api.read_head_time_out")
	if c.apiReadHeadTimeOut == 0 {
		c.apiReadHeadTimeOut = 30
	}
	return nil
}

func (c *Config) ApiHost() string {
	return c.apiHost
}

func (c *Config) ApiPort() int {
	return c.apiPort
}

func (c *Config) ApiProgramKey() string {
	return c.apiProgramKey
}

func (c *Config) ApiAdminKey() string {
	return c.apiAdminKey
}

func (c *Config) ApiReadTimeOut() int {
	return c.apiReadTimeOut
}

func (c *Config) ApiWriteTimeOut() int {
	return c.apiWriteTimeOut
}

func (c *Config) ApiReadHeadTimeOut() int {
	return c.apiReadHeadTimeOut
}
