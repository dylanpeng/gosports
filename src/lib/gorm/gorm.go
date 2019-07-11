package gorm

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

type Pool struct {
	locker  sync.RWMutex
	clients map[string]*gorm.DB
}

func (p *Pool) Add(name string, c *Config) error {
	p.locker.Lock()
	defer p.locker.Unlock()

	orm, err := gorm.Open("mysql", c.GetConnectString())

	if err != nil {
		return err
	}

	p.clients[name] = orm

	return nil
}

func (p *Pool) Get(name string) (*gorm.DB, error) {
	p.locker.RLock()
	defer p.locker.RUnlock()

	client, ok := p.clients[name]

	if ok {
		return client, nil
	}

	return nil, errors.New("no mysql gorm client")
}

func NewPool() *Pool {
	return &Pool{clients: make(map[string]*gorm.DB, 64)}
}

type Config struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Charset  string `toml:"charset"`
	Database string `toml:"database"`
	Debug    bool   `toml:"debug"`
}

func (c *Config) GetConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}
