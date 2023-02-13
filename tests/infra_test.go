package tests

import (
	"log"
	"os"
	"testing"

	"github.com/badico-cloud-hub/go-cache/infra"
	"github.com/badico-cloud-hub/go-cache/utils"
	"github.com/joho/godotenv"
)

func TestNewDynamo(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("TestNewDynamo: expect(nio) - got(%s)\n", err.Error())
	}
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	defaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	tableName := os.Getenv("TABLE_NAME")
	dynamo := infra.NewDynamo(accessKeyId, secretAccessKey, defaultRegion, tableName)
	if dynamo == nil {
		t.Errorf("TestNewDynamo: expect(!nil) - got(nil)\n")
	}
	log.Printf("dynamo: %+v\n", dynamo)
}

func TestDynamoCreateItem(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("TestDynamoCreateItem: expect(nil) - got(%s)\n", err.Error())
	}
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	defaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	tableName := os.Getenv("TABLE_NAME")
	dynamo := infra.NewDynamo(accessKeyId, secretAccessKey, defaultRegion, tableName)
	if dynamo == nil {
		t.Errorf("TestDynamoCreateItem: expect(!nil) - got(nil)\n")
	}
	expiration := utils.GetExpiration(10)
	if err := dynamo.Create("test", "testando", expiration); err != nil {
		t.Errorf("TestDynamoCreateItem: expect(nil) - got(%s)\n", err.Error())
	}
}

func TestDynamoGetItem(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("TestDynamoGetItem: expect(nil) - got(%s)\n", err.Error())
	}
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	defaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	tableName := os.Getenv("TABLE_NAME")
	dynamo := infra.NewDynamo(accessKeyId, secretAccessKey, defaultRegion, tableName)
	if dynamo == nil {
		t.Errorf("TestDynamoGetItem: expect(!nil) - got(nil)\n")
	}
	payload, _, err := dynamo.Get("test")
	if err != nil {
		t.Errorf("TestDynamoGetItem: expect(nil) - got(%s)\n", err.Error())
	}
	log.Printf("payload: %s\n", payload)
}
