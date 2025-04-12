package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"os"

	"github.com/Babahasko/go-jwt-auth/pkg/rsa_loader"
	"github.com/golang-jwt/jwt/v5"
)

const (
	ErrorInvalidToken  = "invalid token"
	ErrorInvalidClaims = "invalid claims"
	ErrorEmailClaim    = "email claim is missing or invalid"
)

type JWTData struct {
	Email string
}

type JWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewJWT(privateKeyFile, publicKeyFile string) *JWT {
	rsaLoader, err := rsa_loader.NewRSA(privateKeyFile, publicKeyFile)
	if err != nil {
		var a []any = []any{err}
		fmt.Fprintf(os.Stdout, "filed to initialize JWT:%w", a...)
	}
	return &JWT{
		PrivateKey: rsaLoader.PrivateKey,
		PublicKey:  rsaLoader.PublicKey,
	}
}

// Создание JWT-токена
func (j *JWT) Create(data JWTData) (string, error) {
	claims := jwt.MapClaims{
		"email": data.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(j.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) Parse(tokenString string) (*JWTData, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected sigining method: %v", t.Header["alg"])
		}
		return j.PublicKey, nil
	})
	if err != nil {
		return nil, errors.New(ErrorInvalidToken)
	}
	if !token.Valid {
		return nil, errors.New(ErrorInvalidToken)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(ErrorInvalidClaims)
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New(ErrorEmailClaim)
	}
	return &JWTData{
		Email: email,
	}, nil
}
