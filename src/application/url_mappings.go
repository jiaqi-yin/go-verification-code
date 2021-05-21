package application

import "github.com/jiaqi-yin/go-verification-code/src/controllers/ping"

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
