package repositories

import "fmt"

func NewMockRepository() *Repository {
	return &Repository{
		Token: &MockTokenRepository{},
	}
}

type MockTokenRepository struct {
	Users []User
}

func (mockRepo *MockTokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
	// TODO: if the users exists just update the information, don't create new entry
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
	for _, user := range mockRepo.Users {
		if user.UserId == userId {
			return user.accessToken, nil
		}
	}
	return "", fmt.Errorf("user not found")
}
