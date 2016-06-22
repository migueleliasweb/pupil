package cache

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

type redisCache struct {
	Conn redis.Conn
}

const (
	userPreffix string = "user_"
)

//Open Opens the conection with the redis server
func (_redisCache *redisCache) Open(config interface{}) error {
	var connErr error

	_redisCache.Conn, connErr = redis.DialURL(os.Getenv("REDIS_URL"))

	return connErr
}

//UserExists Returns true if a user with the given name exists
func (_redisCache *redisCache) UserExists(username string) bool {
	userExists, err := redis.Bool(_redisCache.Conn.Do("EXISTS", userPreffix+username))

	if err != nil {
		log.Println("Could not check user" + username + " for existance.")
	}

	return userExists
}
