package appjwt

import (
	"time"

	"github.com/dany0814/go-apisolutions/pkg/apperrors"
	"github.com/dany0814/go-apisolutions/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type TokenCustomClaims struct {
	ID *string `json:"id"`
	jwt.StandardClaims
}

func CreateToken(id *string) (string, error) {
	now := time.Now().Local()
	claims := TokenCustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour * time.Duration(config.Cfg.ExpiredHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Cfg.Secret))

	return signedToken, err
}

func DecodeToken(tokenPart string) (*TokenCustomClaims, error) {
	claims := &TokenCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Secret), nil
	})

	if err != nil {
		apier := apperrors.NewErrorApi(apperrors.Unauthorized, err.Error())
		return nil, apier
	}

	if !token.Valid {
		apier := apperrors.NewErrorApi(apperrors.Unauthorized, "ID token is invalid")
		return nil, apier
	}

	claims, ok := token.Claims.(*TokenCustomClaims)

	if !ok {
		apier := apperrors.NewErrorApi(apperrors.Unauthorized, "ID token valid but couldn't parse claims")
		return nil, apier
	}

	return claims, err
}
