package resolvers

import (
	"github.com/Oasixer/ShopifyChallenge2021/db"
)

// Resolvers including query and mutation
type Resolvers struct {
	*db.DB
}
