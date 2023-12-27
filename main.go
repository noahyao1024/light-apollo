package main

import (
	"light-apollo/handler"
	"light-apollo/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	storage.Init()

	r.GET("configfiles/json/:app_id/:cluster/:namespace", handler.Configs)
	r.Run(":80")
}
