package migrations

import (
	"github.com/Oasixer/ShopifyChallenge2021/db"
)

// migrate the db
func Migrate(db *db.DB) {
	migrate_1_add_user(db)
	migrate_2_add_files(db)
}
