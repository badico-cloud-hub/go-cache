package infra

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/badico-cloud-hub/go-cache/entity"
)

type Dynamo struct {
	client *dynamodbClient
}

//DynamodbClient is struct for dynamodb client
type dynamodbClient struct {
	client    *dynamodb.DynamoDB
	tableName string
}

//NewDynamo return new instance of dynamo
func NewDynamo(accessKey, secretKey, region, tableName string) *Dynamo {
	dy := dynamodbClient{}
	err := dy.setup(accessKey, secretKey, region, tableName)
	if err != nil {
		log.Fatal(err)
	}
	return &Dynamo{client: &dy}
}

//Create execute creation of item in dynamo
func (dy *Dynamo) Create(key, payload string, expiration int) error {
	err := dy.client.create(key, payload, expiration)
	if err != nil {
		return err
	}
	return nil
}

//Get return item of dynamo
func (dy *Dynamo) Get(key string) (string, int, error) {
	result, exp, err := dy.client.get(key)
	if err != nil {
		return "", 0, err
	}
	return result, exp, nil

}

func (d *dynamodbClient) setup(accessKey, secretKey, region, tableName string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		return err
	}
	cvc := dynamodb.New(sess)
	d.client = cvc
	d.tableName = tableName
	return nil
}

func (d *dynamodbClient) get(key string) (string, int, error) {
	keyCond := expression.Key("PK").Equal(expression.Value(fmt.Sprintf("CACHE_KEY#%s", key)))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		return "", 0, err
	}

	input := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(d.tableName),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	output, err := d.client.QueryWithContext(ctx, input)
	if err != nil {
		return "", 0, err
	}
	caches := []entity.Cache{}
	if err := dynamodbattribute.UnmarshalListOfMaps(output.Items, &caches); err != nil {
		return "", 0, err
	}
	if *output.Count == 0 {
		return "", 0, ErrCacheNotFound
	}

	return caches[0].Value, caches[0].Expiration, nil
}

func (d *dynamodbClient) create(key, payload string, expiration int) error {
	cache := entity.Cache{
		Pk:         fmt.Sprintf("CACHE_KEY#%s", key),
		Sk:         fmt.Sprintf("CACHE_SORT_KEY#%s", key),
		Expiration: expiration,
		Value:      payload,
	}
	inputItem, err := dynamodbattribute.MarshalMap(cache)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item:      inputItem,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = d.client.PutItemWithContext(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
