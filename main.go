package main

import (
	"os"
	"pupil/pupilapi"

	"github.com/kataras/iris"
)

func setupMiddlewares() {
	cacheMW := pupilapi.CacheMiddleware{
		ConnectionString: os.Getenv("REDIS_URL"),
	}

	iris.Use(cacheMW)
}

func main() {
	setupMiddlewares()

	iris.API("/auth", pupilapi.AuthAPI{})
	iris.Listen(":8080")
}
