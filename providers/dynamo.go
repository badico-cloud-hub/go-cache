package providers

import (
	"github.com/badico-cloud-hub/go-cache/infra"
	"github.com/badico-cloud-hub/go-cache/utils"
)

//CacheDynamo is struct the cache with dynamo
type CacheDynamo struct {
	client *infra.Dynamo
}

//NewCacheDynamo return new instance of cache dynamo,
//the table name required is `partition key is PK` and `sort key is SK`,
//in table required prop expiration is enabled with TTL.
func NewCacheDynamo(accessKeyId, secretAccessKey, defaultRegion, tableName string) *CacheDynamo {
	dynamo := infra.NewDynamo(accessKeyId, secretAccessKey, defaultRegion, tableName)
	return &CacheDynamo{
		client: dynamo,
	}
}

//Set is persist cache of payload with key in dynamo
func (p *CacheDynamo) Set(key, payload string, seconds int) (int, error) {
	expiration := utils.GetExpiration(seconds)
	err := p.client.Create(key, payload, expiration)
	if err != nil {
		return 0, err
	}
	return expiration, nil
}

//Get return value of dynamo with key
func (p *CacheDynamo) Get(key string) (string, int, error) {
	result, exp, err := p.client.Get(key)
	if err != nil {
		return "", 0, err
	}
	return result, exp, nil
}
