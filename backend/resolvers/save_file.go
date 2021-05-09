package resolvers

import (
	"context"
	"fmt"
	"github.com/Oasixer/ShopifyChallenge2021/handler"
	"github.com/Oasixer/ShopifyChallenge2021/model"
	"github.com/google/uuid"
)

// SaveFile mutation to create file
func (r *Resolvers) SaveFile(ctx context.Context, args saveFileArgs) (*FileResponse, error){
	userID := ctx.Value(handler.ContextKey("userID"))
	if userID == nil {
		return nil, &getUserError{Code: "NotAuth", Message: "Not Authorized"}
	}
	user := model.User{}
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, &getUserError{Code: "NotFound", Message: "User not found"}
	}

	if !r.DB.Where("UUID = ?", args.UUID).First(&model.File{}).RecordNotFound() {
		return nil, &fileCreationError{Code:"FileExists", Message: "File already exists within DB"}
	}
	
	userID_ := userID.(uint)

	newFile := model.File{Name: args.Name, Uuid: args.UUID, UserID: userID_, Tags: args.Tags}
	result := r.DB.Create(&newFile)
	if result.Error != nil {
		return nil, &fileCreationError{Code: "DBErr", Message: "Database had an error"}
	}
	return &FileResponse{f: &newFile}, nil
}

type saveFileArgs struct {
	UUID uuid.UUID
	Name string
	Tags string
}

// file exists error
type fileCreationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e fileCreationError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e fileCreationError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
} 
