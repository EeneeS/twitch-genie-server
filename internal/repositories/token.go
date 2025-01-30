package repositories

type TokenRepository struct {
	db string
}

type User struct {
	UserId       string
	Login        string
	accessToken  string
	refreshToken string
}

func (repo *TokenRepository) SaveToken(userId, login, accessToken, refreshToken string) error {
	return nil
}
