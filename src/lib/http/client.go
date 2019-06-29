package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	url     string
	reqType string
	header  map[string]string
	body    []byte
}

func (c *Client) Request() ([]byte, error) {
	req, err := http.NewRequest(c.reqType, c.url, bytes.NewReader(c.body))

	if err != nil{
		return nil, err
	}

	for k, v := range c.header{
		req.Header.Add(k, v)
	}

	resp, err := c.client.Do(req)

	if err != nil{
		return nil, err
	}

	defer func(){
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		return nil, err
	}

	return respBody, nil
}

func NewClient(url string, header map[string]string, reqType string, body []byte) *Client {
	client := &Client{
		client:  &http.Client{},
		url:     url,
		reqType: reqType,
		header:  header,
		body:    body,
	}

	return client
}
