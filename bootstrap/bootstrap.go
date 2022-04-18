package bootstrap

import (
	"github.com/Infintive/predictive-go/pkg/database"
	"github.com/Infintive/predictive-go/pkg/env"
	"github.com/Infintive/predictive-go/pkg/router"
	"github.com/gorilla/mux"
)

func NewApplication() *mux.Router {
	env.SetupEnvFile()
	//database.SetupDatabase()
	database.InitMemoryDB()
	appRouter := router.SetupRouter()
	return appRouter
}
