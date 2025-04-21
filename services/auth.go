package services

import (
	"errors"
	"medods/database/dao"
	"medods/database/model"
	"medods/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrTokenNotExist error = errors.New("token not exist")

type AuthService interface {
	CreateTokens(uuid string) (string, string, error)
	TryRefreshToken(token, uuid string) (string, string, error)
}

func Auth() authService {
	return authService{}
}

type authService struct{}

func (this authService) CreateTokens(uuid string) (string, string, error) {
	access, err := utils.CreateToken(jwt.MapClaims{
		"uuid":      uuid,
		"expire_at": time.Now().Add(time.Minute * 15).Unix(),
	})
	if err != nil {
		return "", "", err
	}

	refresh, err := utils.CreateToken(jwt.MapClaims{
		"uuid":      uuid,
		"expire_at": time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		return "", "", nil
	}

	if err := dao.Auth().CreateOrUpdate(&model.Auth{Uuid: uuid, RefreshToken: refresh}); err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (this authService) TryRefreshToken(token, uuid string) (string, string, error) {
	exists, err := dao.Auth().IsExistsToken(token)
	if err != nil {
		return "", "", err
	}

	if !exists {
		return "", "", ErrTokenNotExist
	}

	return this.CreateTokens(uuid)
}
