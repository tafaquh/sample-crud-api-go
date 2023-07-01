package register

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type RegisterUserRequest struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmationToken string `json:confirmation_token`
	Role              string `json:"role"`
	Image             string `json:"image"`
	Nip               string `json:"nip"`
}

type RegisterAdminResponse struct {
	Email string `json:"email"`
}

func generateConfirmationToken() string {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}
	return token + time.Now().Format("20060102150405")
}
