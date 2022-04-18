package main

import (
	"net/http"

	"github.com/Infintive/predictive-go/bootstrap"
	"github.com/Infintive/predictive-go/pkg/env"
)

func main() {
	appRouter := bootstrap.NewApplication()
	port := env.GetEnv("APP_PORT", "4000")
	http.ListenAndServe(":"+port, appRouter)
}
