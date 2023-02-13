package interfaces

type ICache interface {
	Get(key string) (string, int, error)
	Set(key, payload string, expiration int) (int, error)
}
