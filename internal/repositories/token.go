package repositories

type TokenRepository struct {
	db string
}

// rename to SaveUser
func (repo *TokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
  return nil
}

func (repo *TokenRepository) GetAccessToken(userId string) (string, error) {
	return "", nil
}
