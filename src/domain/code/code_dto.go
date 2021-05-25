package code

import (
	"strings"

	"github.com/jiaqi-yin/go-verification-code/src/utils"
)

type CodeGenerator struct {
	Phone string `json:"phone_number"`
}

func (codeGenerator *CodeGenerator) Validate() utils.RestErr {
	codeGenerator.Phone = strings.TrimSpace(codeGenerator.Phone)
	if codeGenerator.Phone == "" {
		return utils.NewBadRequestError("missing phone_number")
	}
	return nil
}
