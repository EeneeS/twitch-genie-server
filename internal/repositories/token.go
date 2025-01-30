package repositories

type TokenRepository struct {
	db string
}

func (repo *TokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
	return nil
}
