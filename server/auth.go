package server

type User struct {
	Username string
	Email    string
	Password string
	RootPath string
}
type UserProvider interface {
	GetUser(username string) (*User, error)
	Exists(username string) (bool, error)
}

type AuthService interface {
	Verify(username string, password string) (bool, error)
}

type auth struct {
	UserProvider UserProvider
	AuthService  AuthService
}

var DefaultAuth *auth

func InitAuth(userProvider UserProvider, authService AuthService) {
	DefaultAuth = &auth{
		UserProvider: userProvider,
		AuthService:  authService,
	}
}

func GetUser(username string) (*User, error) {
	return DefaultAuth.UserProvider.GetUser(username)
}
