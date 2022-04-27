package main

import (
	"log"
	"meli-challenge-compliance/pkg/config"
	custom_error "meli-challenge-compliance/pkg/errors"
	googleauth "meli-challenge-compliance/pkg/google-auth"
	httpserver "meli-challenge-compliance/pkg/http"
	drivemanagement "meli-challenge-compliance/pkg/service"
	"net/http"
)

func main() {
	// llama a LoadConfig para parcear las credenciales
	cfg, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal(custom_error.ErrBadConfiguration)
	}

	// se crea el googleClient para realizar la autenticacion
	googleClient, err := googleauth.GetGoogleClient(cfg.CredentialFileLocation, cfg.TokenFileLocation)
	if err != nil {
		log.Fatal(custom_error.ErrNotDefined)
	}

	driveManagementComponent := drivemanagement.NewGoogleDriveComponent(googleClient)
	handler := drivemanagement.NewDriveHandler(driveManagementComponent)
	routes := httpserver.NewRoutes(handler)

	// lanza el servidor HTTP
	log.Fatal(http.ListenAndServe(cfg.HTTPPort, routes))
}
