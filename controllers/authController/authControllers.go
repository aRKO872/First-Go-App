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
	_, userExistsError := checkUserExists(req.GetEmail(), false)

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


func (s *AuthServer) LogIn(ctx context.Context, req *pb.LoginUserInput) (*pb.LoginUserResponse, error) {
	valid := validEmail (req.GetEmail()) && validPassword (req.GetPassword())

	if !valid {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: "Has incorrect password or email type!",
		}, nil
	}

	loginUserData, userExistsError := checkUserExists(req.GetEmail(), true)

	if userExistsError != nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: userExistsError.Error(),
		}, nil
	}

	passwordsMatch := comparePasswords (req.GetPassword(), loginUserData.Password);

	if !passwordsMatch {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: "user unauthorized. please check password or email",
		}, nil
	}

	// Generate Access and Refresh Tokens 
	access_token, accessTokenErr := generateToken(
		config.GetConfig().JWT_SECRET_KEY, 
		time.Duration(config.GetConfig().ACCESS_TOKEN_EXPIRY),
		loginUserData.GetId(),
	)

	refresh_token, refreshTokenErr := generateToken(
		config.GetConfig().JWT_SECRET_KEY, 
		time.Duration(config.GetConfig().REFRESH_TOKEN_EXPIRY),
		loginUserData.GetId(),
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

	// Successful Sign In Returns Access and Refresh Tokens :
	return &pb.LoginUserResponse{
		Status: true,
		AccessToken: access_token,
		RefreshToken: refresh_token,
	}, nil
}

func (s *AuthServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenInput) (*pb.LoginUserResponse, error) {
	input_access_token, input_refresh_token := req.GetOldAccessToken(), req.GetRefreshToken()

	oldTokenDerivedPayload, oldTokenErr := ParseExpiredToken(input_access_token, config.GetConfig().JWT_SECRET_KEY);

	if oldTokenErr != nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: oldTokenErr.Error(),
		}, nil
	}

	refreshTokenPayload, refreshTokenErr := ParseToken(input_refresh_token, config.GetConfig().JWT_SECRET_KEY)

	if refreshTokenErr != nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: refreshTokenErr.Error(),
		}, nil
	}

	if oldTokenDerivedPayload.UserId != refreshTokenPayload.UserId {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: "invalid user",
		}, nil
	}

	// Generate Access and Refresh Tokens 
	access_token, accessTokenErr := generateToken(
		config.GetConfig().JWT_SECRET_KEY, 
		time.Duration(config.GetConfig().ACCESS_TOKEN_EXPIRY),
		oldTokenDerivedPayload.UserId,
	)

	refresh_token, refreshTokenErr := generateToken(
		config.GetConfig().JWT_SECRET_KEY, 
		time.Duration(config.GetConfig().REFRESH_TOKEN_EXPIRY),
		oldTokenDerivedPayload.UserId,
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

	// Successful Sign In Returns Access and Refresh Tokens :
	return &pb.LoginUserResponse{
		Status: true,
		AccessToken: access_token,
		RefreshToken: refresh_token,
	}, nil
}