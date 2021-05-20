package resolvers

import (
	"testing"
	"log"

	"github.com/Oasixer/ShopifyChallenge2021/db"
	"github.com/Oasixer/ShopifyChallenge2021/model"
)

func TestSignIn(t *testing.T) {
	log.Printf("test sign in start") // TODO remove temp
	db, err := db.ConnectDB()

	if err != nil {
		t.Errorf(err.Error())
	}
	user := model.User{}
	db.DB.Where("email = ?", "notexisting@test.com").First(&user)

	t.Log(user.ID)
}
