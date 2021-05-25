package code

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiaqi-yin/go-verification-code/src/domain/code"
	"github.com/jiaqi-yin/go-verification-code/src/services"
	"github.com/jiaqi-yin/go-verification-code/src/utils"
)

var (
	CodeController codeControllerInterface = &codeController{}
)

type codeControllerInterface interface {
	Generate(*gin.Context)
	Verify(*gin.Context)
}

type codeController struct{}

func (controller *codeController) Generate(c *gin.Context) {
	var codeGenerator code.CodeGenerator
	if err := c.ShouldBindJSON(&codeGenerator); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.CodeService.Generate(codeGenerator)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (controller *codeController) Verify(c *gin.Context) {
	isValid := services.CodeService.Verify("123456")
	c.JSON(http.StatusOK, isValid)
}