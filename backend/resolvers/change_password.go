package resolvers

import (
	"context"
	"fmt"
	"log"

	"github.com/Oasixer/ShopifyChallenge2021/handler"
	"github.com/Oasixer/ShopifyChallenge2021/model"
	"github.com/Oasixer/ShopifyChallenge2021/utils"
)

// ChangePassword mutation change password
func (r *Resolvers) ChangePassword(ctx context.Context, args changePasswordMutationArgs) (*UserResponse, error) {
	userID := ctx.Value(handler.ContextKey("userID"))

	if userID == nil {
		return nil, &changePassError{Code: "NotAuth", Message: "Not authorized"}
	}
	user := model.User{}

	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, &changePassError{Code: "NotExist", Message: "User Doesnt Exist"}
	}

	hashPassword, err := utils.HashPass(args.Password)
	if err != nil {
		log.Printf("ERROR, couldnt hash password: %#v", err)
		return nil, err
	}

	user.Password = hashPassword

	r.DB.Save(&user)
	return &UserResponse{u: &user}, nil
}

type changePasswordMutationArgs struct {
	Password string
}

type changePassError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e changePassError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e changePassError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
