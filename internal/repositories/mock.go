package repositories

import "fmt"

func NewMockRepository() *Repository {
	return &Repository{
		Token: &MockTokenRepository{},
	}
}

type MockTokenRepository struct{}

func (mockRepo *MockTokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
	fmt.Printf("saving user: %v\n", login)
	return nil
}
