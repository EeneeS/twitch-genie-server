package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func (repo *UserRepository) SaveUser(userId, login, accessToken, refreshToken string) error {
  user := bson.M{
    "userId": userId,
    "login": login,
    "accessToken": accessToken,
    "refreshToken": refreshToken,
  }
  _, err := repo.db.InsertOne(context.TODO(), user)
  return err
}

func (repo *UserRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
