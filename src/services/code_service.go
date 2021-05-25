package services

import (
	"github.com/jiaqi-yin/go-verification-code/src/domain/code"
	"github.com/jiaqi-yin/go-verification-code/src/utils"
)

var (
	CodeService codeServiceInterface = &codeService{}
)

type codeServiceInterface interface {
	Generate(code.CodeGenerator) (bool, utils.RestErr)
	Verify(string) bool
}

type codeService struct{}

func (s *codeService) Generate(codeGenerator code.CodeGenerator) (bool, utils.RestErr) {
	if err := codeGenerator.Validate(); err != nil {
		return false, err
	}

	if err := codeGenerator.Generate(); err != nil {
		return false, err
	}

	return true, nil
}

func (s *codeService) Verify(code string) bool {
	return true
}
