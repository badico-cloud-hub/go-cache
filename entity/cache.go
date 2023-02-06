package entity

type Cache struct {
	Pk         string `dynamodbav:"PK,omitempty"`
	Sk         string `dynamodbav:"SK,omitempty"`
	Value      string `dynamodbav:"value,omitempty"`
	Expiration int    `dynamodbav:"expiration,omitempty"`
}
