package application

import (
	"github.com/gin-gonic/gin"
	appconfig "github.com/jiaqi-yin/go-verification-code/src/app_config"
	"github.com/jiaqi-yin/go-verification-code/src/clients/redis"
)

var (
	router = gin.Default()
)

func StartApplication() {
	appconfig.LoadConfig(".")

	redis.Init()

	mapUrls()

	router.Run(appconfig.Config.ServerAddr)
}
