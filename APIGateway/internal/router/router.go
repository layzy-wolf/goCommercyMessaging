package router

import (
	"github.com/julienschmidt/httprouter"
)

var prefix = "api"

func NewRouter() *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/test", testRoute())
	mux.POST(prefix+"/register", Register())

	return mux
}
