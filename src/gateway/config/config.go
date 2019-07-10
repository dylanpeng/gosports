package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var config *Config

type Config struct {
	EnvConfig *EnvConfig           `toml:"env"`
	DBConfigs map[string]*DBConfig `toml:"dbs"`
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
