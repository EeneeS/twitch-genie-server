package repositories

type Repository struct {
	Token interface {
		SaveToken(userId, login, accessToken, refreshToken string) error
	}
}

func NewRepository(db string) *Repository {
	return &Repository{
		Token: &TokenRepository{db: db},
	}
}
