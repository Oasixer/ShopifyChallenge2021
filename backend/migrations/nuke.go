package migrations

import (
	"github.com/Oasixer/ShopifyChallenge2021/db"
	"github.com/Oasixer/ShopifyChallenge2021/model"
)

// delete the tables
func Nuke(db *db.DB) {
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.File{})
}
