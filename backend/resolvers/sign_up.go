package resolvers

import (
	"fmt"
	"log"

	"github.com/Oasixer/ShopifyChallenge2021/model"
	"github.com/Oasixer/ShopifyChallenge2021/utils"
)

// SignUp mutation creates user
func (r *Resolvers) SignUp(args signUpMutationArgs) (*UserResponse, error) {

	// hash the password
	hashPassword, err := utils.HashPass(args.Password)
	if err != nil {
		log.Printf("ERROR, couldnt hash password: %#v", err)
		return nil, err
	}

	newUser := model.User{Email: args.Email, Password: hashPassword, Username: args.Username}

	// if user already exists return an error of UserExists
	if !r.DB.Where("email = ?", args.Email).First(&model.User{}).RecordNotFound() {
		return nil, &signUpError{Code: "UserExists", Message: "User already signed up"}
	}

	result := r.DB.Create(&newUser)
	if result.Error != nil {
		return nil, &signUpError{Code: "DBErr", Message: "Database had an error"}
	}

	return &UserResponse{u: &newUser}, nil
}

// arguements
type signUpMutationArgs struct {
	Email    string
	Password string
	Username string
}

// user exists error
type signUpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e signUpError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e signUpError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
