package main

import (
	"fmt"
	"gosports/common/consts"
	"gosports/lib/http"
)

func main(){
	url := "http://www.baidu.com"

	rep, err := http.Get(url)
	if err != nil{
		fmt.Printf("get url failed url: %s | err: %s \n", url, err)
	}else {
		fmt.Printf("get url seccess. content: %s \n", string(rep))
	}

	rep, err = http.Post(url, nil)
	if err != nil{
		fmt.Printf("post url failed url: %s | err: %s \n", url, err)
	}else {
		fmt.Printf("post url seccess. content: %s \n", string(rep))
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["User-Agent"] = "ORide-driver/2.11.111"
	body := "{\"phone_number\":\"01111122222\",\"password\":\"123456\"}"
	rep, err = http.Requset("http://dev.api.o-pay.in/driver/login", header, consts.HttpPost, []byte(body))
	if err != nil{
		fmt.Printf("request get url failed url: %s | err: %s \n", url, err)
	}else {
		fmt.Printf("request get url seccess. content: %s \n", string(rep))
	}

	client := http.NewClient("http://dev.api.o-pay.in/driver/login", header, consts.HttpPost, []byte(body))
	rep, err = client.Request()
	if err != nil{
		fmt.Printf("request get url failed url: %s | err: %s \n", url, err)
	}else {
		fmt.Printf("request get url seccess. content: %s \n", string(rep))
	}
}
