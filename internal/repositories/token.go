package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TokenRepository struct {
	db *pgx.Conn
}

// rename to SaveUser
func (repo *TokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
  sql := `INSERT INTO users (user_id, login, access_token, refresh_token)
  VALUES ($1, $2, $3, $4)
  ON CONFLICT (user_id)
  DO UPDATE SET
  login = EXCLUDED.login, 
  access_token = EXCLUDED.access_token, 
  refresh_token = EXCLUDED.refresh_token`

  _, err := repo.db.Exec(context.Background(), sql, userId, login, accessToken, refreshToken)
  if err != nil {
    return fmt.Errorf("failed to save user: %v", err)
  }

  return nil
}

func (repo *TokenRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
