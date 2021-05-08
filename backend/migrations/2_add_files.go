package migrations

import (
	"github.com/Oasixer/ShopifyChallenge2021/db"
	"github.com/Oasixer/ShopifyChallenge2021/model"
)

func migrate_2_add_files(db *db.DB) {
	db.AutoMigrate(&model.File{})
}
