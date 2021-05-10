package migrations

import (
	"github.com/Oasixer/ShopifyChallenge2021/db"
	"github.com/Oasixer/ShopifyChallenge2021/model"
)

// delete the tables
func Nuke(db *db.DB) {
	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.File{})
}
