package pupilapi

import "github.com/kataras/iris"

//AuthAPI Handler for the /auth
type AuthAPI struct {
	*iris.Context
}

//Post Performs the auth
func (api AuthAPI) Post() {
	api.Write("FOO=BAR")
}
