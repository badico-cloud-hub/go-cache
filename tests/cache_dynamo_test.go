package tests

import (
	"log"
	"os"
	"testing"

	"github.com/badico-cloud-hub/go-cache/providers"
	"github.com/joho/godotenv"
)

func TestNewCacheDynamo(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("TestNewCacheDynamo: expect(nio) - got(%s)\n", err.Error())
	}
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	defaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	tableName := os.Getenv("TABLE_NAME")
	dynamo := providers.NewCacheDynamo(accessKeyId, secretAccessKey, defaultRegion, tableName)
	if dynamo == nil {
		t.Errorf("TestNewProviderDynamo: expect(!nil) - got(nil)\n")
	}
	log.Printf("dynamo: %+v\n", dynamo)
}
