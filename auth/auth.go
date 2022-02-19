package auth

import "context"

type Auth interface {
	GetToken() (now int64, jwtToken string, err error)
	GetSecretValue(ctx context.Context, key string) interface{}
}
