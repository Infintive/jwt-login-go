package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Infintive/predictive-go/bootstrap"
	"github.com/Infintive/predictive-go/pkg/env"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func main() {
	appRouter := bootstrap.NewApplication()
	port := env.GetEnv("APP_PORT", "4000")

	keyFunc := func(ctx context.Context) (interface{}, error) {
		// Our token must be signed using this data.
		return []byte("secret"), nil
	}

	// Set up the validator.
	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
		"https://<issuer-url>/",
		[]string{"<audience>"},
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	// Set up the middleware.
	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)

	http.ListenAndServe(":"+port, appRouter)
}
