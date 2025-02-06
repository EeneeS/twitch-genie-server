package repositories

import (
	"fmt"
	"sync"
)

func NewMockRepository() *Repository {
	return &Repository{
		Token: &MockTokenRepository{},
	}
}

type User struct {
	UserId       string
	Login        string
	accessToken  string
	refreshToken string
}

type MockTokenRepository struct {
	Users []User
	Mutex sync.RWMutex
}

func (mockRepo *MockTokenRepository) SaveUser(userId, login, accessToken, refreshToken string) error {
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

func (mockRepo *MockTokenRepository) GetAccessToken(userId string) (string, error) {
	mockRepo.Mutex.RLock()
	defer mockRepo.Mutex.RUnlock()

	for _, user := range mockRepo.Users {
		if user.UserId == userId {
			return user.accessToken, nil
		}
	}

	return "", fmt.Errorf("user not found")
}
