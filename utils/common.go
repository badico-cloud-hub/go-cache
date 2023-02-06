package utils

import (
	"time"
)

//GetExpiration return expiration for dynamodb
func GetExpiration(seconds int) int {
	exp := time.Now().Add(time.Second * time.Duration(seconds)).Unix()
	return int(exp)
}
