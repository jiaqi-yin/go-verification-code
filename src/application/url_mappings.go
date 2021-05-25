package application

import (
	"github.com/jiaqi-yin/go-verification-code/src/controllers/code"
	"github.com/jiaqi-yin/go-verification-code/src/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.PingController.Ping)
	router.POST("/generate", code.CodeController.Generate)
	router.POST("/verify", code.CodeController.Verify)
}
