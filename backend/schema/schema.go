package schema

import (
	gql "github.com/mattdamon108/gqlmerge/lib"
	"github.com/Oasixer/ShopifyChallenge2021/utils"
	"log"
	"fmt"
)

func NewSchema() *string {
	root := utils.RootDir()
	log.Printf(root)
	dir := fmt.Sprintf("%s/schema", root)
	log.Printf(dir)
	schema := gql.Merge("  ", dir)

	return schema
}
