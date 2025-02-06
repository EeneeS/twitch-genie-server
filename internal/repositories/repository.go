package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repository struct {
	Token interface {
		SaveUser(userId, login, accessToken, refreshToken string) error
		GetAccessToken(userId string) (string, error)
	}
  Media interface {
    SaveMedia(channelId, source string, xpos, ypos int) error
    GetMedia(channelId string) error
  }
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Token: &UserRepository{db: db.Database("twitch-genie-db").Collection("users")},
    Media: &MediaRepository{db: db.Database("twitch-genie-db").Collection("media")},
	}
}
