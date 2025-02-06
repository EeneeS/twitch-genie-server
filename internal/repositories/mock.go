package repositories

import (
	"fmt"
	"sync"
)

func NewMockRepository() *Repository {
	return &Repository{
		Token: &MockUserRepository{},
    Media: &MockMediaRepository{},
	}
}

type User struct {
	UserId       string
	Login        string
	accessToken  string
	refreshToken string
}

type MockUserRepository struct {
	Users []User
	Mutex sync.RWMutex
}

type MockMediaRepository struct {
  Media []Media
  Mutex sync.RWMutex
}

func (mockRepo *MockUserRepository) SaveUser(userId, login, accessToken, refreshToken string) error {
	mockRepo.Mutex.Lock()
	defer mockRepo.Mutex.Unlock()

	for i, user := range mockRepo.Users {
		if user.UserId == userId {
			mockRepo.Users[i].Login = login
			mockRepo.Users[i].accessToken = accessToken
			mockRepo.Users[i].refreshToken = refreshToken
			return nil
		}
	}

	newUser := User{
		UserId:       userId,
		Login:        login,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
	mockRepo.Users = append(mockRepo.Users, newUser)

	return nil
}

func (mockRepo *MockUserRepository) GetAccessToken(userId string) (string, error) {
	mockRepo.Mutex.RLock()
	defer mockRepo.Mutex.RUnlock()

	for _, user := range mockRepo.Users {
		if user.UserId == userId {
			return user.accessToken, nil
		}
	}

	return "", fmt.Errorf("user not found")
}

func (repo *MockMediaRepository) SaveMedia(channelId, source string, xpos, ypos int) error {
  return nil
}

func (repo *MockMediaRepository) GetMedia(channelId string) error {
  return nil
}

func (repo *MockMediaRepository) RemoveAllMedia(channelId string) error {
  return nil
}
