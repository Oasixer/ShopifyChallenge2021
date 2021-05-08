package resolvers

import (
	"context"
	"fmt"

	"github.com/Oasixer/ShopifyChallenge2021/handler"
	"github.com/Oasixer/ShopifyChallenge2021/model"
)

// GetUser resolver
func (r *Resolvers) GetUser(ctx context.Context) (*UserResponse, error) {
	userID := ctx.Value(handler.ContextKey("userID"))

	if userID == nil {
		return nil, &getUserError{Code: "NotAuth", Message: "Not Authorized"}
	}

	user := model.User{}
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, &getUserError{Code: "NotFound", Message: "User not found"}
	}
	return &UserResponse{u: &user}, nil
}

// user exists error
type getUserError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e getUserError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e getUserError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
