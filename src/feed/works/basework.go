package works

import (
	"errors"
	"gosports/lib/http"
)

type BaseWork struct {
	Client *http.Client
}

func (b *BaseWork) Request() (string, error) {
	if b.Client == nil {
		return "", errors.New("http Client is nil")
	}

	resp, err := b.Client.Request()

	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func NewBaseWork(url string, header map[string]string, reqType string, body []byte) *BaseWork {
	return &BaseWork{Client: http.NewClient(url, header, reqType, body)}
}
