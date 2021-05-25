package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(*gin.Context)
}

type pingController struct{}

func (controller *pingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
