package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ptisma/go-rest-sample-app/cmd/api/handlers/createuser"
	"github.com/ptisma/go-rest-sample-app/cmd/api/handlers/getuser"
	"github.com/ptisma/go-rest-sample-app/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users/:id", getuser.Do(app))
	mux.POST("/users", createuser.Do(app))
	return mux
}
