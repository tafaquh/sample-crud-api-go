package register

import (
	"coba-go/database/schema"
	"coba-go/utils/errors"

	"coba-go/utils/password"
)

var (
	RegisterRepository registerRepositoryInterface = &registerRepository{}
)

type registerRepositoryInterface interface {
	CreateUserToDB(payload *RegisterUserRequest) (*RegisterAdminResponse, *errors.RestErr)
	GetUserByConfirmationToken(token string) (*schema.Users, *errors.RestErr)
	UpdateUser(user *schema.Users) *errors.RestErr
}
type registerRepository struct{}

func (r *registerRepository) CreateUserToDB(payload *RegisterUserRequest) (*RegisterAdminResponse, *errors.RestErr) {
	pass := password.HashAndSalt(password.GetPwd(payload.Password))
	userData := &schema.Users{
		Name:      payload.FirstName + " " + payload.LastName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  pass,
		//StatusID:          1,
		ConfirmationToken: payload.ConfirmationToken,
		AuthToken:         "",
		Image:             payload.Image,
		Role:              payload.Role,
		Nip:               payload.Nip,
	}
	userData.ConfirmationToken = generateConfirmationToken()

	if err := schema.Database.Create(userData); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to save callback user to db: " + err.Error.Error())
	}
	return &RegisterAdminResponse{Email: payload.Email}, nil
}

func (r *registerRepository) GetUserByConfirmationToken(token string) (*schema.Users, *errors.RestErr) {
	var user schema.Users
	if err := schema.Database.First(&user, "confirmation_token = ?", token); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to get user by confirmation token: " + err.Error.Error())
	}
	return &user, nil
}

func (r *registerRepository) UpdateUser(user *schema.Users) *errors.RestErr {
	if err := schema.Database.Save(user); err.Error != nil {
		return errors.NewInternalServerError("error when tryin to update users: " + err.Error.Error())
	}
	return nil
}
