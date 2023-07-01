package login

import (
	goErr "errors"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserUUID  string `json:"id"`
	UserNip   string `json:"user_nip"`
	IsSuccess bool   `json:"is_success"`
	Token     string `json:"token"`
}

func (s *LoginRequest) CreateToken() string {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = false
	atClaims["id"] = s.ID
	atClaims["email"] = s.Email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET_RECIPE")))
	if err != nil {
		return ""
	}
	return token
}

func ExtractToken(bearer string) (res string) {
	str := strings.Split(bearer, " ")
	if len(str) == 2 {
		res = str[1]
		return
	}
	return
}

func VerifyToken(key string, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, goErr.New("ErrSessDecode")
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, goErr.New("ErrSessVerifation")
	}
	return token, nil
}

func TokenValid(token *jwt.Token) error {
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return goErr.New("ErrSessInvalid")
	}
	return nil
}

func ExtractTokenMetadata(token *jwt.Token) map[string]interface{} {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims
	}
	return nil
}
