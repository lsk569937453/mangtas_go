package main

import (
	"encoding/json"
	"fmt"
	"real_project/vojo"

	"github.com/valyala/fasthttp"
)

func main() {
	url := "http://127.0.0.1:9393/mangtas/parseTopText"
	req := &fasthttp.Request{}
	req.SetRequestURI(url)
	reqText := "a b c d abc d c b a cba ab bc ac ca ba abc cab cba a"

	topReq := vojo.TopRequst{
		SourceText: &reqText,
	}
	reqJson, err := json.Marshal(topReq)
	if err != nil {
		panic(err)
	}
	req.SetBody(reqJson)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		fmt.Println("request error:", err.Error())
		return
	}
	if resp.StatusCode() != 200 {
		fmt.Println("status code  error:", err.Error())
		return
	}

	b := resp.Body()

	resJson := vojo.TopResponse{}
	err = json.Unmarshal(b, &resJson)
	if err != nil {
		panic(err)
	}
	if resJson.Rescode == 0 {
		outPut, err := json.Marshal(resJson.ResMessage)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(outPut))
		//result is [{"key":"a","times":3},{"key":"abc","times":2},{"key":"cba","times":2},{"key":"d","times":2},{"key":"b","times":2},{"key":"c","times":2},{"key":"ca","times":1},{"key":"ba","times":1},{"key":"bc","times":1},{"key":"ab","times":1},{"key":"ac","times":1},{"key":"cab","times":1}]

	} else {
		panic(fmt.Errorf("response code is not zero"))
	}

}
