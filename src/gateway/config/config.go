package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gosports/lib/logger"
)

var config *Config

type Config struct {
	EnvConfig *EnvConfig           `toml:"env"`
	DBConfigs map[string]*DBConfig `toml:"dbs"`
	LogConfig *logger.Config       `toml:"log"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

type EnvConfig struct {
	Debug bool     `toml:"debug"`
	Host  string   `toml:"host"`
	Port  int      `toml:"port"`
	IPs   []string `toml:"ips"`
	Ids   []int    `toml:"ids"`
}

func (c *EnvConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}

type DBConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func (c *DBConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}

func GetConfig() *Config {
	return config
}

func GetEnvConfig() *EnvConfig {
	return config.EnvConfig
}

func GetDBConfigs() map[string]*DBConfig {
	return config.DBConfigs
}

func GetLogConfig() *logger.Config {
	return config.LogConfig
}

func Init(file string) error {
	config = &Config{}

	_, err := toml.DecodeFile(file, config)
	if err != nil {
		return err
	}
	return nil
}
