package resolvers

import (
	"strconv"
	"log"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/Oasixer/ShopifyChallenge2021/model"
	// "github.com/Oasixer/ShopifyChallenge2021/db"
)

// UserResponse is the user response type
type UserResponse struct {
	u *model.User
}

// ID for UserResponse
func (r *UserResponse) ID() graphql.ID {
	id := strconv.Itoa(int(r.u.ID))
	return graphql.ID(id)
}

// Email for UserResponse
func (r *UserResponse) Email() string {
	return r.u.Email
}

// Password for UserResponse
func (r *UserResponse) Password() string {
	return r.u.Password
}

// Username for UserResponse
func (r *UserResponse) Username() string {
	return r.u.Username
}

// CreatedAt for UserResponse
func (r *UserResponse) CreatedAt() string {
	return r.u.CreatedAt.String()
}

// UpdatedAt for UserResponse
func (r *UserResponse) UpdatedAt() string {
	return r.u.UpdatedAt.String()
}

// Files for UserResponse
func (r *UserResponse) Files() []*FileResponse {
	log.Printf("%v",len(r.u.Files))
	var files []*FileResponse
	for _, file := range r.u.Files {
		files = append(files, &FileResponse{&file})
	}
	// var files []model.File
	// db.DB.Where(&files, "UserID = ?",r.u.ID).Find(&files)
	// files := db.DB.Find(&)
	return files
}
