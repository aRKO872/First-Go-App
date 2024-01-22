package authcontrollers

import (
	"context"
	"time"

	"github.com/aRKO872/first-go-app/config"
	pb "github.com/aRKO872/first-go-app/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb.RegisterUserInput) (*pb.LoginUserResponse, error) {

	valid := validName (req.GetName()) && validEmail (req.GetEmail()) && validPassword (req.GetPassword())

	if !valid {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: "Has incorrect name, password or email type!",
		}, nil
	}

	// Checking if User trying to sign up already exists
	_, userExistsError := checkUserExists(req, false)

	if userExistsError != nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: userExistsError.Error(),
		}, nil
	}

	// Encrypting Provided Password and Storing User info in DB.
	savedUser, persistUserError := persistUser(req);

	if persistUserError != nil || savedUser == nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: persistUserError.Error(),
		}, nil
	}

	// Generate Access and Refresh Tokens 
	access_token, accessTokenErr := generateToken(
		config.GetConfig().JWT_SECRET_KEY, 
		time.Duration(config.GetConfig().ACCESS_TOKEN_EXPIRY),
		savedUser.GetId(),
	)

	refresh_token, refreshTokenErr := generateToken(
		config.GetConfig().JWT_SECRET_KEY, 
		time.Duration(config.GetConfig().REFRESH_TOKEN_EXPIRY),
		savedUser.GetId(),
	)

	if accessTokenErr != nil || refreshTokenErr != nil {
		errStr := ""
		if accessTokenErr != nil && refreshTokenErr != nil {
			errStr = accessTokenErr.Error()
		} else if accessTokenErr != nil {
			errStr = accessTokenErr.Error()
		} else {
			errStr = refreshTokenErr.Error()
		}
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: errStr,
		}, nil
	}

	// Successful Sign Up Returns Access and Refresh Tokens :
	return &pb.LoginUserResponse{
		Status: true,
		AccessToken: access_token,
		RefreshToken: refresh_token,
	}, nil
}