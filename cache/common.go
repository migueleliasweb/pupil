package cache

//Cache Interface that must be satisfied to be used as cache driver
type Cache interface {
	// returns cache value by key
	Get(key string) ([]byte, error)

	// Sets key to value, config is driver specific
	Set(key string, value []byte, config interface{}) error

	// Deletes key in cache
	Delete(key string) error

	// Increments key, 1 is default
	Incr(key string, num int) (int, error)

	// Decrements key, 1 is default
	Decr(key string, num int) (int, error)

	// Closes the connection
	Close() error
}
