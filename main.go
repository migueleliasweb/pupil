package main

import (
	"os"
	pupilHandler "pupil/handler"
	"pupil/util"

	"github.com/kataras/iris"
)

func setupMiddlewares() {
	cacheMW := pupilHandler.CacheMiddleware{
		ConnectionString: os.Getenv("REDIS_URL"),
	}

	iris.Use(cacheMW)
}

func setupJSONError() {
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		//ctx.JSON(iris.StatusOK, map[string]string{"hello": "json"})
		ctx.Write("{\"error\":\"Not Found\"}")
	})
}

func main() {
	util.Log("FOO", util.LogLevelDebug)
	setupMiddlewares()
	setupJSONError()

	iris.API("/auth", pupilHandler.AuthAPI{})

	iris.Listen(":8080")
}
