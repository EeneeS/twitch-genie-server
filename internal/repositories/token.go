package repositories

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepository struct {
	db *mongo.Collection
}

func (repo *UserRepository) SaveUser(userId, login, accessToken, refreshToken string) error {
    filter := bson.M{"userId": userId}

    update := bson.M{
        "$set": bson.M{
            "login":        login,
            "accessToken":  accessToken,
            "refreshToken": refreshToken,
            "updatedAt":    time.Now(),
        },
    }

    opts := options.UpdateOne().SetUpsert(true)

    _, err := repo.db.UpdateOne(context.TODO(), filter, update, opts)
    if err != nil {
        return fmt.Errorf("failed to save user: %w", err)
    }

    return nil
}

func (repo *UserRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
