package resolvers

import (
	"context"
	"fmt"
	"log"
	// "strconv"

	"github.com/Oasixer/ShopifyChallenge2021/handler"
	"github.com/Oasixer/ShopifyChallenge2021/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SaveFile mutation to create file
func (r *Resolvers) SaveFile(ctx context.Context, args saveFileArgs) (*FileResponse, error) {
	log.Print("saving file")
	userID := ctx.Value(handler.ContextKey("userID"))
	if userID == nil {
		return nil, &getUserError{Code: "NotAuth", Message: "Not Authorized"}
	}
	user := model.User{}
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, &getUserError{Code: "NotFound", Message: "User not found"}
	}

	uuid_ := uuid.MustParse(args.Uuid)
	newFile := model.File{Name: args.Name, Uuid: uuid_, Tags: args.Tags}
	r.DB.Model(&user).Association("Files").Append(&newFile);
	r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return &FileResponse{f: &newFile}, nil
}

type saveFileArgs struct {
	Uuid string
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
