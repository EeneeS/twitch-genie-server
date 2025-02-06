package repositories

import (
	"github.com/jackc/pgx/v5"
)

type TokenRepository struct {
	db *pgx.Conn
}

type User struct {
	UserId       string
	Login        string
	accessToken  string
	refreshToken string
}

func (repo *TokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
	return nil
}

func (repo *TokenRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
