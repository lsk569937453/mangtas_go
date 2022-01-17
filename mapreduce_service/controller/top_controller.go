package controller

import (
	"mapreduce_service/log"
	"mapreduce_service/service"
	"mapreduce_service/vojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseTopText(c *gin.Context) {

	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS

	err, resJson := service.ParseTopText(c)
	if err != nil {
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		res.ResMessage = err.Error()
		log.Error("ParseTopText error", err.Error())

	} else {
		res.ResMessage = resJson
	}
	c.JSON(http.StatusOK, res)

}
