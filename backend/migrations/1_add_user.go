package migrations

import (
	"github.com/Oasixer/ShopifyChallenge2021/db"
	"github.com/Oasixer/ShopifyChallenge2021/model"
)

func migrate_1_add_user(db *db.DB) {
	db.AutoMigrate(&model.User{})
}
