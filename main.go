package main

import (
	"os"
	pupilHandler "pupil/handler"

	"github.com/kataras/iris"
)

func setupMiddlewares() {
	cacheMW := pupilHandler.CacheMiddleware{
		ConnectionString: os.Getenv("REDIS_URL"),
	}

	iris.Use(cacheMW)
}

func main() {
	setupMiddlewares()

	iris.API("/auth", pupilHandler.AuthAPI{})

	iris.Listen(":8080")
}
