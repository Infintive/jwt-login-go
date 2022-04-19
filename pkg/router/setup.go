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
	UserRouter.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	UserRouter.HandleFunc("/login", controllers.LoginHandler).Methods("POST")

	return UserRouter
}
