package services

import (
	"github.com/jiaqi-yin/go-verification-code/src/clients/sms"
	"github.com/jiaqi-yin/go-verification-code/src/domain/code"
	"github.com/jiaqi-yin/go-verification-code/src/utils"
)

var (
	CodeService codeServiceInterface = &codeService{}
)

type codeServiceInterface interface {
	Generate(code.CodeGenerator, sms.SmsClientInterface) (bool, utils.RestErr)
	Verify(code.CodeVerifier) (bool, utils.RestErr)
}

type codeService struct{}

func (s *codeService) Generate(codeGenerator code.CodeGenerator, smsClient sms.SmsClientInterface) (bool, utils.RestErr) {
	if err := codeGenerator.Validate(); err != nil {
		return false, err
	}

	code, err := codeGenerator.Generate()
	if err != nil {
		return false, err
	}

	if err := smsClient.SendMessage(codeGenerator.Phone, code); err != nil {
		return false, err
	}

	return true, nil
}

func (s *codeService) Verify(codeVerifier code.CodeVerifier) (bool, utils.RestErr) {
	if err := codeVerifier.Validate(); err != nil {
		return false, err
	}

	if err := codeVerifier.Verify(); err != nil {
		return false, err
	}

	return true, nil
}
