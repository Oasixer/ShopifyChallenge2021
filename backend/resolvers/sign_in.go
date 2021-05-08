package resolvers

import (
	"fmt"
	"strconv"

	"github.com/Oasixer/ShopifyChallenge2021/model"
	"github.com/Oasixer/ShopifyChallenge2021/utils"
)

// SignIn mutation creates user
func (r *Resolvers) SignIn(args signInMutationArgs) (*string, error) {
	user := model.User{}

	r.DB.Where("email = ?", args.Email).First(&user)

	if user.ID == 0 {
		return nil, &signInError{Code: "NotSignedUp", Message: "User is not signed up"}
	}

	if !utils.CompareHashPass(args.Password, user.Password) {
		return nil, &signInError{Code: "PassBad", Message: "Wrong Password"}
	}

	userIDString := strconv.Itoa(int(user.ID))
	tokenString, err := utils.SignJWT(&userIDString)
	if err != nil {
		return nil, &signInError{Code: "ErrJWT", Message: "Error generating JWT"}
	}

	return tokenString, nil
}

// sign in error
type signInError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e signInError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e signInError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

type signInMutationArgs struct {
	Email    string
	Password string
}
