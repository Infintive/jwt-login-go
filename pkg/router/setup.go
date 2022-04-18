package router

import (
	"github.com/Infintive/predictive-go/app/controllers"
	"github.com/gorilla/mux"
)

//var UserRouter *mux.Router

func SetupRouter() *mux.Router {
	UserRouter := mux.NewRouter()
	UserRouter.StrictSlash(true)
	UserRouter.HandleFunc("/", controllers.BaseHandler)
	UserRouter.HandleFunc("/register", controllers.RegisterHandler)
	UserRouter.HandleFunc("/login", controllers.LoginHandler)

	return UserRouter
}
