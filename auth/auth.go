package auth

type Auth interface {
	GetToken() (now int64, jwtToken string, err error)
}
