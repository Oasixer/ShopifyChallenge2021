package resolvers

import (
	"context"
	"fmt"
	"log"
	"strconv"

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
	log.Print("FUCK_1")
	user := model.User{}
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, &getUserError{Code: "NotFound", Message: "User not found"}
	}

	uuid_ := uuid.MustParse(args.Uuid)

	// if !r.DB.Where("Uuid = ?", uuid_).First(&model.File{}).RecordNotFound() {
	// return nil, &fileCreationError{Code:"FileExists", Message: "File already exists within DB"}
	// }
	// log.Print("FUCK_2")

	userID64, _ := strconv.ParseUint(userID.(string), 10, 32)

	newFile := model.File{Name: args.Name, Uuid: uuid_, UserID: uint(userID64), Tags: args.Tags}
	r.DB.Model(&user).Association("Files").Append(&newFile);
	log.Print("FUCK_3")
	result := r.DB.Create(&newFile)
	if result.Error != nil {
		return nil, &fileCreationError{Code: "DBErr", Message: "Database had an error"}
	}
	result = r.DB.Save(&newFile)
	result = r.DB.Save(&user)

	r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user)

	// r.DB.S
	// r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user)
	// log.Printf("User files[0]:%s",user.Files[0].Name)
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
