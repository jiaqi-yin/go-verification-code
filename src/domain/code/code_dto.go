package code

import (
	"strings"

	"github.com/jiaqi-yin/go-verification-code/src/utils"
)

type CodeGenerator struct {
	Phone string `json:"phone_number"`
}

type CodeVerifier struct {
	Phone string `json:"phone_number"`
	Code  string `json:"verification_code"`
}

func (codeGenerator *CodeGenerator) Validate() utils.RestErr {
	codeGenerator.Phone = strings.TrimSpace(codeGenerator.Phone)
	if codeGenerator.Phone == "" {
		return utils.NewBadRequestError("missing phone_number")
	}
	return nil
}

func (codeVerifier *CodeVerifier) Validate() utils.RestErr {
	codeVerifier.Phone = strings.TrimSpace(codeVerifier.Phone)
	codeVerifier.Code = strings.TrimSpace(codeVerifier.Code)
	if codeVerifier.Phone == "" {
		return utils.NewBadRequestError("missing phone_number")
	}
	if codeVerifier.Code == "" {
		return utils.NewBadRequestError("missing verification_code")
	}
	return nil
}
