package interfaces

type ICache interface {
	Get(key string) (string, error)
	Set(key, payload string, expiration int) error
}
