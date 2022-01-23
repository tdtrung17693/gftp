package security

type UserProvider interface {
	GetUser(username string)
	Exists(username string)
}

type Auth interface {
	Verify(username string, password string)
}
