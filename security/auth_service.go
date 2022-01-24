package security

type AuthService struct {
	context      *AuthContext
	userProvider *SQLiteUserProvider
}

func NewAuthService(ctx *AuthContext, userProvider *SQLiteUserProvider) *AuthService {
	return &AuthService{
		context:      ctx,
		userProvider: userProvider,
	}
}

func (a *AuthService) Verify(username string, password string) (bool, error) {
	return true, nil
}
