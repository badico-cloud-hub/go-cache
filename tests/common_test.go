package tests

import (
	"log"
	"testing"

	"github.com/badico-cloud-hub/go-cache/utils"
	"github.com/joho/godotenv"
)

func TestGetExpiration(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("TestGetExpiration: expect(nil) - got(%s)\n", err.Error())
	}
	expiration := utils.GetExpiration(30)
	if expiration == 0 {
		t.Errorf("TestGetExpiration: expect(!nil) - got(nil)\n")
	}
	log.Printf("expiration: %+v\n", expiration)
}
