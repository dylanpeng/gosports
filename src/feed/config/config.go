package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gosports/lib/gorm"
)

var config *Config

type Config struct {
	DBConfigs   map[string]*gorm.Config `toml:"dbs"`
	WorkConfigs *WorkConfig             `toml:"works"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

type WorkConfig struct {
	MatchUrl      string `toml:"match_url"`
	MatchInterval int64  `toml:"match_interval"`
	TeamUrl       string `toml:"team_url"`
	TeamInterval  int64  `toml:"team_interval"`
}

func (c *WorkConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}

func Init(file string) error {
	config = &Config{}
	_, err := toml.DecodeFile(file, config)

	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}

func GetDBConfig() map[string]*gorm.Config {
	return config.DBConfigs
}

func GetWorkConfig() *WorkConfig {
	return config.WorkConfigs
}
