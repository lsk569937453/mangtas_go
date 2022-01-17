package service

import (
	"fmt"
	"mapreduce_service/vojo"
	"sort"
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

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Times > ss[j].Times
	})
	return nil, ss

}
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}
