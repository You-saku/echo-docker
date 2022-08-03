package token

import (
	"time"

	// jwtのライブラリ
	"github.com/golang-jwt/jwt/v4"
)

const Key = "my-secure-jwt-key" // 秘密鍵(ダミー)

type JwtCustomClaims struct {
	Email  string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	jwt.StandardClaims
}

// golangは(引数1 型, 引数2 型) (返り値1 型, 返り値2 型)とかで定義できる
func CreateJwt(name string, password string) (string, error){
	// Set custom claims
	claims := &JwtCustomClaims{
		name,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(Key));
}
