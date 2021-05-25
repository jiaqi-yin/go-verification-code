package application

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaqi-yin/go-verification-code/src/clients/redis"
)

var (
	router = gin.Default()
)

func StartApplication() {
	redis.Init()

	mapUrls()

	router.Run(":8080")
}
