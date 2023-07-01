package services

import (
	"coba-go/repository/login"
	"coba-go/utils/errors"
)

var (
	LoginService loginServiceInterface = &loginService{}
)

type loginService struct {
}

type loginServiceInterface interface {
	Login(*login.LoginRequest) (interface{}, *errors.RestErr)
}

func (s *loginService) Login(payload *login.LoginRequest) (interface{}, *errors.RestErr) {
	result, err := login.LoginRepository.Login(payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}
