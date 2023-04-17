package jwt

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	Role   int32 `json:"role"`
	jwt.RegisteredClaims
}

//go:embed secret/id_rsa
var _prvKey []byte

//go:embed secret/id_rsa.pub
var _publicKey []byte

func GenerateJwt(userID int64, permission int32) string {
	var err error

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(_prvKey)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Create the claims
	claims := Claims{
		userID,
		permission,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Println(err)
		return ""
	}
	return signedToken
}

func VerifyJwt(token string, permission Permission) (*Claims, error) {
	var err error

	if err != nil {
		return nil, err
	}
	parsedPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(_publicKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	decodedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return parsedPublicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := decodedToken.Claims.(*Claims); ok && decodedToken.Valid {

		if claims.Role >= int32(permission) {
			return claims, nil
		} else {
			return nil, errors.New("invalid JWT")
		}

	} else {
		return nil, errors.New("invalid JWT")
	}
}
