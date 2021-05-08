package resolvers

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/Oasixer/ShopifyChallenge2021/model"
)

// file response type
type FileResponse struct {
	f *model.File
}

func (r *FileResponse) ID() graphql.ID {
	id := strconv.Itoa(int(r.f.ID))
	return graphql.ID(id)
}

func (r *FileResponse) Url() string {
	return r.f.Url
}

func (r *FileResponse) FileType() string {
	return r.f.FileType
}

func (r *FileResponse) DocumentType() string {
	return r.f.DocumentType
}

func (r *FileResponse) Name() string {
	return r.f.Name
}

func (r *FileResponse) CreatedAt() string {
	return r.f.CreatedAt.String()
}

func (r *FileResponse) UpdatedAt() string {
	return r.f.UpdatedAt.String()
}
