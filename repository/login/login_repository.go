package login

import (
	"coba-go/database/schema"
	"coba-go/utils/errors"

	"coba-go/utils/password"
)

var (
	LoginRepository loginRepositoryInterface = &loginRepository{}
)

type loginRepositoryInterface interface {
	Login(payload *LoginRequest) (*LoginResponse, *errors.RestErr)
	GetUserByEmail(email string) (*schema.Users, *errors.RestErr)
}
type loginRepository struct{}

func (r *loginRepository) Login(payload *LoginRequest) (*LoginResponse, *errors.RestErr) {
	res := LoginResponse{}
	user, err := r.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}
	res.IsSuccess = false
	res.UserUUID = user.ID
	res.UserNip = user.Nip
	// Check password
	if password.ComparePasswords(user.Password, password.GetPwd(payload.Password)) {
		// Create the token
		res.Token = payload.CreateToken()
		if res.Token == "" {
			return nil, errors.NewInternalServerError("cannot generate token")
		}
	} else {
		return nil, errors.NewBadRequestError("email or password not match")
	}
	res.IsSuccess = true
	return &res, nil
}

func (r *loginRepository) GetUserByEmail(email string) (*schema.Users, *errors.RestErr) {
	var user schema.Users
	if err := schema.Database.First(&user, "email = ?", email); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to get user by confirmation token: " + err.Error.Error())
	}
	if user.StatusID < 2 {
		return nil, errors.NewBadRequestError("You must confirm your email first")
	}
	return &user, nil
}
