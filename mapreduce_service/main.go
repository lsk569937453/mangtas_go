package main

import (
	"fmt"
	"mapreduce_service/controller"
	"mapreduce_service/log"
	"runtime"

	"github.com/gin-gonic/gin"
)

const PORT = "9393"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	err := initController()
	if err != nil {
		log.Error("initController error", err.Error())
	}

}

func initController() error {

	gin.DefaultWriter = log.BaseGinLog()
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[GIN] %v |%3d| %13v | %15s | %-7s  %#v %s |\"%s\" \n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
			param.Request.UserAgent(),
		)
	}))

	r.POST("/mangtas/parseTopText", controller.ParseTopText)

	err := r.Run(":" + PORT) // listen and serve

	return err

}
