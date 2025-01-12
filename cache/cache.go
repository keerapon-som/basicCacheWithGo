package cache

type Cache interface {
	Set(key string, value interface{}, duration int) error
	Get(key string, receiver interface{}) error
}
