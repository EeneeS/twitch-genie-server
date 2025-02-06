package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

type TokenRepository struct {
	db *mongo.Client
}

func (repo *TokenRepository) SaveUser(userId, login, accessToken, refreshToken string) error {
  return nil
}

func (repo *TokenRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
