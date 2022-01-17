package service

import (
	"fmt"
	"mapreduce_service/util"
	"mapreduce_service/vojo"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseTopText(c *gin.Context) (error, []vojo.TopServiceKV) {

	var req vojo.TopRequst
	err := c.ShouldBind(&req)
	if err != nil {
		return fmt.Errorf("pare json  body error"), nil
	}
	if *req.SourceText == "" {
		return fmt.Errorf("empty input"), nil
	}
	resJson := WordCount(*req.SourceText)

	if len(resJson) == 0 {
		return fmt.Errorf("no word input"), nil
	}

	var ss []vojo.TopServiceKV
	for k, v := range resJson {
		ss = append(ss, vojo.TopServiceKV{k, v})
	}
	heapLength := 10
	if len(ss) < 10 {
		heapLength = len(ss)
	}
	startNums := ss[:heapLength]
	util.MiniHeap(startNums)
	for _, data := range ss[heapLength:] {
		if data.Times <= startNums[0].Times {
			continue
		} else {
			startNums[0] = data
			util.MiniHeap(startNums)
		}
	}
	return nil, startNums

}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}
