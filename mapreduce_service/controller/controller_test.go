package controller

import (
	"encoding/json"
	"fmt"
	"mapreduce_service/constants"
	"mapreduce_service/util"
	"mapreduce_service/vojo"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

const ParseTopNUrl = "/mangtas/parseTopText"

func TestParseJsonWrongJsonStruct(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := ParseTopNUrl
	body := struct {
		wrongText string
	}{"100 200 300 400"}

	w = util.PostJson(urlIndex, body, router, nil)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)
	assert.Equal("pare json  body error", baseRes.ResMessage)
}
func TestParseJsonEmptyInput(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := ParseTopNUrl
	sourceText := ""
	body := vojo.TopRequst{
		SourceText: &sourceText,
	}
	w = util.PostJson(urlIndex, body, router, nil)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)
	assert.Equal("empty input", baseRes.ResMessage)
}

func TestParseJsonSpaceInput(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := ParseTopNUrl
	sourceText := "    "
	body := vojo.TopRequst{
		SourceText: &sourceText,
	}
	w = util.PostJson(urlIndex, body, router, nil)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)
	assert.Equal("no word input", baseRes.ResMessage)

}

func TestTopServiceOk(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := ParseTopNUrl
	sourceText := "test1 test2 test3 a b c abc a b c test3 test2 test1"
	body := vojo.TopRequst{
		SourceText: &sourceText,
	}
	w = util.PostJson(urlIndex, body, router, nil)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()

	type topNResponse struct {
		vojo.BaseRes
		ResMessage []vojo.TopServiceKV `json:"resMessage"`
	}
	var baseRes topNResponse
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)

	if len(baseRes.ResMessage) != 7 {
		t.Error("error result length")
	}
	for _, value := range baseRes.ResMessage {
		if value.Key == "test1" {
			if value.Times != 2 {
				t.Error("test1 result error")
			}
		}
		if value.Key == "test2" {
			if value.Times != 2 {
				t.Error("test2 result error")
			}
		}
		if value.Key == "test3" {
			if value.Times != 2 {
				t.Error("test3 result error")
			}
		}
		if value.Key == "a" {
			if value.Times != 2 {
				t.Error("a result error")
			}
		}
		if value.Key == "b" {
			if value.Times != 2 {
				t.Error("b result error")
			}
		}
		if value.Key == "c" {
			if value.Times != 2 {
				t.Error("c result error")
			}
		}
		if value.Key == "abc" {
			if value.Times != 1 {
				t.Error("abc result error")
			}
		}
	}
}
func TestTopServiceTen(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := ParseTopNUrl
	sourceText := "test1 test2 test3 a b c abc a b c test3 test2 test1 a a a  a a a a a a a a b b b b b b b b b  c c c c c c  c c c c c c c c c cc j k l i m p o h f e"
	body := vojo.TopRequst{
		SourceText: &sourceText,
	}
	w = util.PostJson(urlIndex, body, router, nil)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	fmt.Println(resBody)

	type topNResponse struct {
		vojo.BaseRes
		ResMessage []vojo.TopServiceKV `json:"resMessage"`
	}
	var baseRes topNResponse
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)

	if len(baseRes.ResMessage) != 10 {
		t.Error("error result length")
	}
	for _, value := range baseRes.ResMessage {
		if value.Key == "abc" {
			if value.Times != 1 {
				t.Error("abc result error")
			}
		}
		if value.Key == "test2" {
			if value.Times != 2 {
				t.Error("test2 result error")
			}
		}
		if value.Key == "test3" {
			if value.Times != 2 {
				t.Error("test3 result error")
			}
		}
		if value.Key == "p" {
			if value.Times != 1 {
				t.Error("p result error")
			}
		}
		if value.Key == "test1" {
			if value.Times != 2 {
				t.Error("test1 result error")
			}
		}
		if value.Key == "b" {
			if value.Times != 11 {
				t.Error("b result error")
			}
		}
		if value.Key == "c" {
			if value.Times != 17 {
				t.Error("c result error")
			}
		}
		if value.Key == "cc" {
			if value.Times != 1 {
				t.Error("cc result error")
			}
		}
		if value.Key == "o" {
			if value.Times != 1 {
				t.Error("o result error")
			}
		}
		if value.Key == "a" {
			if value.Times != 13 {
				t.Error("a result error")
			}
		}
	}

}
func TestMain(m *testing.M) {
	fmt.Println("test begin")
	InitRouter()
	m.Run()
	fmt.Println("test end")
}

func InitRouter() {
	r := gin.New()
	r.POST("/mangtas/parseTopText", ParseTopText)

	router = r

}
