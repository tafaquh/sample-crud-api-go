package services

import (
	"coba-go/repository/register"
	"coba-go/utils/errors"
)

var (
	RegisterService registerServiceInterface = &registerService{}
)

type registerService struct {
}

type registerServiceInterface interface {
	CreateUser(*register.RegisterUserRequest) (interface{}, *errors.RestErr)
}

func (s *registerService) CreateUser(payload *register.RegisterUserRequest) (interface{}, *errors.RestErr) {
	result, err := register.RegisterRepository.CreateUserToDB(payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}
