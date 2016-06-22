package handler

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
)

//CacheMiddleware Setups cache features
type CacheMiddleware struct {
	ConnectionString string
}

//Serve Handles the middleware implementation
func (mw CacheMiddleware) Serve(ctx *iris.Context) {
	redisConn, errConn := redis.DialURL(mw.ConnectionString)

	if errConn != nil {
		log.Println("Could not connect to redis server error: " + errConn.Error())
		return
	}

	ctx.Set("cache", redisConn)
	ctx.Next()
}
