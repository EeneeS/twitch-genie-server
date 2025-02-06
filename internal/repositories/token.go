package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

type UserRepository struct {
	db *mongo.Client
}

func (repo *UserRepository) SaveUser(userId, login, accessToken, refreshToken string) error {
  return nil
}

func (repo *UserRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
