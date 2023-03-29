package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DB struct {
	dbEnable   bool
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string
	dbSSLMode  bool
}

func (c *DB) ConfigDB() error {
	c.dbEnable = viper.GetBool("db.enable")

	c.dbHost = viper.GetString("db.host")
	if c.dbHost == "" {
		c.dbHost = "127.0.0.1"
	}
	c.dbPort = viper.GetInt("db.port")
	if c.dbPort == 0 {
		c.dbPort = 5432
	}
	c.dbUser = viper.GetString("db.user")
	if c.dbUser == "" {
		return fmt.Errorf("db.users not in config")
	}
	c.dbPassword = viper.GetString("db.password")
	if c.dbPassword == "" {
		return fmt.Errorf("db.password not in config")
	}
	c.dbName = viper.GetString("db.dbname")
	if c.dbName == "" {
		return fmt.Errorf("db.dbname not in config")
	}
	c.dbSSLMode = viper.GetBool("db.ssl_mode")
	return nil
}

func (c *DB) DBEnable() bool {
	return c.dbEnable
}

func (c *DB) DBHost() string {
	return c.dbHost
}

func (c *DB) DBPort() int {
	return c.dbPort
}

func (c *DB) DBUser() string {
	return c.dbUser
}

func (c *DB) DBPassword() string {
	return c.dbPassword
}

func (c *DB) DBName() string {
	return c.dbName
}

func (c *DB) DBSSLMode() bool {
	return c.dbSSLMode
}

func (c *DB) DBConfig() DB {
	return *c
}
