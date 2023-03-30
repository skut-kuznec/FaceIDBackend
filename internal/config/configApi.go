package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type API struct {
	apiHost            string
	apiPort            int
	apiProgramKey      string
	apiAdminKey        string
	apiReadTimeOut     time.Duration
	apiWriteTimeOut    time.Duration
	apiReadHeadTimeOut time.Duration
}

func (a *API) ConfigAAPI() error {
	a.apiHost = viper.GetString("api.host")
	if a.apiHost == "" {
		a.apiHost = "127.0.0.1"
	}
	a.apiPort = viper.GetInt("api.port")
	if a.apiPort == 0 {
		a.apiPort = 8080
	}
	a.apiProgramKey = viper.GetString("api.program_key")
	if a.apiProgramKey == "" {
		return fmt.Errorf("api.program_key not in config")
	}
	a.apiAdminKey = viper.GetString("api.admin_key")
	if a.apiAdminKey == "" {
		return fmt.Errorf("api.admin_key not in config")
	}
	a.apiReadTimeOut = viper.GetDuration("api.read_time_out")
	if a.apiReadTimeOut == 0 {
		a.apiReadTimeOut = 30 * time.Second
	}
	a.apiWriteTimeOut = viper.GetDuration("api.write_time_out")
	if a.apiWriteTimeOut == 0 {
		a.apiWriteTimeOut = 30 * time.Second
	}
	a.apiReadHeadTimeOut = viper.GetDuration("api.read_head_time_out")
	if a.apiReadHeadTimeOut == 0 {
		a.apiReadHeadTimeOut = 30 * time.Second
	}
	return nil
}

func (a *API) APIHost() string {
	return a.apiHost
}

func (a *API) APIPort() int {
	return a.apiPort
}

func (a *API) APIProgramKey() string {
	return a.apiProgramKey
}

func (a *API) APIAdminKey() string {
	return a.apiAdminKey
}

func (a *API) APIReadTimeOut() time.Duration {
	return a.apiReadTimeOut
}

func (a *API) APIWriteTimeOut() time.Duration {
	return a.apiWriteTimeOut
}

func (a *API) APIReadHeadTimeOut() time.Duration {
	return a.apiReadHeadTimeOut
}
