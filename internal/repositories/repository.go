package repositories

import "github.com/jackc/pgx/v5"

type Repository struct {
	Token interface {
		SaveToken(userId, login, accessToken, refreshToken string) error
		GetAccessToken(userId string) (string, error)
	}
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Token: &TokenRepository{db: db},
	}
}
