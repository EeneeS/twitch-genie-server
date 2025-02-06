package repositories

import "go.mongodb.org/mongo-driver/v2/mongo"

type Media struct {
  Source string `json:"source"`
  PositionX int `json:"position_x"`
  PositionY int `json:"position_y"`
}

type MediaRepository struct {
	db *mongo.Collection
}

func (repo *MediaRepository) SaveMedia(channelId, source string, xpos, ypos int) error {
  return nil
}

func (repo *MediaRepository) GetMedia(channelId string) error {
  return nil
}

func (repo *MediaRepository) RemoveAllMedia(channelId string) error {
  return nil
}
