package verify_auth

import (
	"context"
	"errors"
	"strings"

	"github.com/aRKO872/first-go-app/config"
	authcontrollers "github.com/aRKO872/first-go-app/controllers/authController"
	"google.golang.org/grpc/metadata"
)

func VerifyAuthToken (ctx context.Context) (*authcontrollers.Claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("missing metadata")
	}

	authArr := md.Get("authorization")
	if len(authArr) == 0 {
		return nil, errors.New("empty metadata")
	}

	authData := authArr[0]
	authToken := strings.Fields(authData)[1];

	tokenDerivedPayload, tokenErr := authcontrollers.ParseToken(authToken, config.GetConfig().JWT_SECRET_KEY)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return tokenDerivedPayload, nil;
}

