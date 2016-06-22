package handler

iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
    ctx.Write("{\"error\":\"Not Found\"}")
})
