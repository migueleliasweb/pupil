package cache

//Cache Interface that must be satisfied to be used as cache driver
type Cache interface {
	//Open the connection, config is driver specific
	Open(config interface{}) error

	//Get Returns cache value by key
	Get(key string) ([]byte, error)

	//Set Sets key to value, config is driver specific
	Set(key string, value []byte, config ...interface{}) error

	//Delete Deletes key in cache
	Delete(key string) error

	//Incr Increments key, 1 is default
	Incr(key string, num int) (int, error)

	//Decr Decrements key, 1 is default
	Decr(key string, num int) (int, error)

	//Close Closes the connection
	Close() error
}
