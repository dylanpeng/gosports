package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var config *Config

type Config struct {
	EnvConfig *EnvConfig           `toml:"env"`
	DBConfigs map[string]*DBConfig `toml:"dbs"`
	Fruits    []*Fruit             `toml:"fruit"`
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

type Fruit struct {
	Name      string     `toml:"name"`
	Physical  *Physical  `toml:"physical"`
	Varieties []*Variety `toml:"variety"`
}

func (c *Fruit) String() string {
	return fmt.Sprintf("%+v", *c)
}

type Physical struct {
	Color string `toml:"color"`
	Shape string `toml:"shape"`
}

func (c *Physical) String() string {
	return fmt.Sprintf("%+v", *c)
}

type Variety struct {
	Name string `toml:"name"`
}

func (c *Variety) String() string {
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
