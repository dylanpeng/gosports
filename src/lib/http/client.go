package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	Url     string
	ReqType string
	Header  map[string]string
	Body    []byte
}

func (c *Client) Request() ([]byte, error) {
	req, err := http.NewRequest(c.ReqType, c.Url, bytes.NewReader(c.Body))

	if err != nil{
		return nil, err
	}

	for k, v := range c.Header{
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
		Url:     url,
		ReqType: reqType,
		Header:  header,
		Body:    body,
	}

	return client
}
