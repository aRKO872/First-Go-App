package authcontrollers

import (
	"context"

	
	pb "github.com/aRKO872/first-go-app/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb.RegisterUserInput) (*pb.LoginUserResponse) {

	valid := validName (req.GetName()) && validEmail (req.GetEmail()) && validPassword (req.GetPassword())

	if !valid {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: "Has incorrect name, password or email type!",
		}
	}

	// Checking if User trying to sign up already exists
	_, userExistsError := checkUserExists(req, false)

	if userExistsError != nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: userExistsError.Error(),
		}
	}

	// Encrypting Provided Password and Storing User info in DB.
	savedUser, persistUserError := persistUser(req);

	if persistUserError != nil || savedUser == nil {
		return &pb.LoginUserResponse {
			Status: false,
			ErrMsg: persistUserError.Error(),
		}
	}

	
}