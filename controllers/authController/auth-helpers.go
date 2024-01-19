package authcontrollers

import (
	"strings"
	"regexp"
	"errors"
	pb "github.com/aRKO872/first-go-app/proto"

	"github.com/aRKO872/first-go-app/controllers/dbcontroller"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

func validName (name string) bool {
	name = strings.TrimSpace(name);
	return len(name) > 3 
}

func validEmail (email string) bool {
	email = strings.TrimSpace(email);
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func validPassword (password string) bool {
	password = strings.TrimSpace(password);
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*(),.?":{}|<>]).{8,}$`
	re := regexp.MustCompile(passwordRegex)
	return re.MatchString(password)
}

func checkUserExists (req *pb.RegisterUserInput, fromLogin bool) (*pb.Users, error) {
	var user *pb.Users;
	result := dbcontroller.DB.Where("email = ?", req.GetEmail()).First (&user);

	if result.Error != nil {
		return nil, errors.New ("Error Occured While Checking Existence of User")
	}

	// For Registration
	if !fromLogin && user.GetId() != "" {
		return nil, errors.New ("User already exists! Please try logging in!")
	}

	// For Logging In
	if fromLogin && user.GetId() == "" {
		return nil, errors.New ("User Does Not Exist. Please Sign Up!")
	}

	if fromLogin {
		return user, nil
	}
	return nil, nil;
}

func persistUser (
	req *pb.RegisterUserInput,
) (*pb.Users, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil && len(hashedPasswordBytes) > 0 {
		return nil, errors.New("Server error, unable to create your account.")
	}

	hashedPassword := string (hashedPasswordBytes[:]);

	saveUser := &pb.Users {
		Id : uuid.NewString(),
		Name : req.GetName(),
		Email : req.GetEmail(),
		Password : hashedPassword,
	}

	result := dbcontroller.DB.Create(&saveUser)

	if result.RowsAffected == 0 {
		return nil, errors.New("Error Persisting User!")
	}

	return saveUser, nil;
}

func getUserToken (
	user_id string
) (
	string, string, error
) {
	
}