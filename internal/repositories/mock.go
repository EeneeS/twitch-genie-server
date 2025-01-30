package repositories

func NewMockRepository() *Repository {
	return &Repository{
		Token: &MockTokenRepository{},
	}
}

type MockTokenRepository struct {
	Users []User
}

func (mockRepo *MockTokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
	newUser := User{
		UserId:       userId,
		Login:        login,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
	mockRepo.Users = append(mockRepo.Users, newUser)
	return nil
}
