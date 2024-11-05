package helper

import (
	"errors"
	"time"

	"github.com/MohdAhzan/auth-svc/pkg/models"
	"github.com/golang-jwt/jwt/v4"
)

type JwtWrapper struct{
  SecretKey string
  Issuer string
  ExpiryHours int64
}

type jwtClaims struct{

    jwt.RegisteredClaims
    Id int64
    Email string
}


func (w *JwtWrapper) GenerateToken(user models.Users) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,

		RegisteredClaims:  jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(w.ExpiryHours))),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now().Local()){
		return nil, errors.New("JWT is expired")
	}

	return claims, nil

}
