package http

import (
	"bytes"
	"gosports/common/consts"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func Post(url string, body []byte) ([]byte, error) {
	resp, err := http.Post(url, consts.HeaderContentTypeJson, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func Request(url string, header map[string]string, reqType string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(reqType, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respBody, nil
}
