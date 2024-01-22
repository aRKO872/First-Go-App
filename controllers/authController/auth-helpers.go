package authcontrollers

import (
	"errors"
	"regexp"
	"strings"
	"time"

	pb "github.com/aRKO872/first-go-app/proto"
	"gorm.io/gorm"

	"github.com/aRKO872/first-go-app/controllers/dbcontroller"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserId string 
	jwt.RegisteredClaims
}

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

func validPassword(password string) bool {
	password = strings.TrimSpace(password)

	// Check minimum length
	if len(password) < 8 {
		return false
	}

	// Check for at least one lowercase letter
	hasLower := false
	for _, char := range password {
		if 'a' <= char && char <= 'z' {
			hasLower = true
			break
		}
	}

	if !hasLower {
		return false
	}

	// Check for at least one uppercase letter
	hasUpper := false
	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasUpper = true
			break
		}
	}

	if !hasUpper {
		return false
	}

	// Check for at least one digit
	hasDigit := false
	for _, char := range password {
		if '0' <= char && char <= '9' {
			hasDigit = true
			break
		}
	}

	if !hasDigit {
		return false
	}

	// Check for at least one special character
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !specialCharRegex.MatchString(password) {
		return false
	}

	return true
}

func checkUserExists (req *pb.RegisterUserInput, fromLogin bool) (*pb.Users, error) {
	var user *pb.Users;
	result := dbcontroller.DB.Where("email = ?", req.GetEmail()).First (&user);

	if result.Error != nil {
		// User not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New ("error occured while checking existence of user")
	}

	// For Registration
	if !fromLogin && user.GetId() != "" {
		return nil, errors.New ("user already exists! please try logging in")
	}

	// For Logging In
	if fromLogin && user.GetId() == "" {
		return nil, errors.New ("user does not xxist. please sign up")
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
		return nil, errors.New("server error, unable to create your account")
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
		return nil, errors.New("error persisting user")
	}

	return saveUser, nil;
}

func generateToken(key string, expiration time.Duration, user_id string) (string, error) {
	claims := &Claims{
		UserId: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key[:]))
	if err != nil {
		return "", errors.New("error occured while generating token")
	}

	return signedToken, nil
}

func ParseToken(tokenString string, key string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, errors.New("error occurred while parsing token")
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return claims, nil
}