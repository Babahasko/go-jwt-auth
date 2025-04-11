package jwt

import (
	"crypto/rsa"
	"fmt"

	"github.com/Babahasko/go-jwt-auth/pkg/rsa_loader"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func NewJWT(privateKeyFile, publicKeyFile string) *JWT {
	rsaLoader, err := rsa_loader.NewRSA(privateKeyFile, publicKeyFile)
	if err != nil {
		fmt.Println("filed to initialize JWT")
	}
	return &JWT{
		PrivateKey: rsaLoader.PrivateKey,
		PublicKey: rsaLoader.PublicKey,
	}
}

// Создание JWT-токена
func (j *JWT) Create(email string) (string, error) {
    claims := jwt.MapClaims{
        "email": email,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    tokenString, err := token.SignedString(j.PrivateKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// Проверка JWT-токена
// func (j *JWT) verifyToken(tokenString string) (*jwt.Token, error) {
//     token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//         if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//             return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//         }
//         return publicKey, nil
//     })

//     if err != nil {
//         return nil, fmt.Errorf("failed to parse token: %w", err)
//     }

//     if !token.Valid {
//         return nil, fmt.Errorf("invalid token")
//     }
//     return token, nil
// }
