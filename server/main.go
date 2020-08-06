package main

import (
	"encoding/json"
	"fmt"
	catalog "landlord-app/server/internal/catalog"
	"landlord-app/server/internal/handler"
	stdout "landlord-app/server/internal/logger/stdout"
	system "landlord-app/server/pkg/system/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configurations
	config, err := system.LoadEnvironmentConfig()
	errFatalHandle(err, catalog.ErrorLoadingConfig)

	// Configure logger
	logger, err := stdout.NewLogger(config.LogLevel)
	errFatalHandle(err, catalog.ErrorInitLogger)

	// Initialize router and set API Routes
	// HTTP context
	responder := handler.NewResponder(logger)

	router := chi.NewRouter()

	router.Get("/", HandleRoot)
	router.NotFound(responder.NotFoundHandler)

	fmt.Println("Starting server on the port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// HandleRoot will load the root
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode("this is the homepage")
}

func errFatalHandle(err error, message string) {
	if err != nil {
		logrus.WithError(err).Fatalln(message)
	}
}
