package main

import (
	"context"
	"fmt"
	appcontext "landlord-app/server/internal/appcontext"
	catalog "landlord-app/server/internal/catalog"
	"landlord-app/server/internal/handler"
	stdout "landlord-app/server/internal/logger/stdout"
	"landlord-app/server/pkg/system"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	ctx := appcontext.WithLogger(context.Background(), logger)

	// Initialize router and set API Routes
	// HTTP context
	responder := handler.NewResponder(logger)

	// Handlers
	user := handler.NewUser(ctx, responder)

	router := chi.NewRouter()

	// TODO: Separate this into a separate repo for the client and drop /views/ url namespace
	// Create a route along "/" root path that will serve contents from
	// the client/views/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "client/views"))
	FileServer(router, "/views", filesDir)

	// API Routes
	router.Get("/api/user/{userID}", user.Handle)
	router.Get("/api/user/{userID}/tenants", user.HandleTenants)
	router.Get("/api/user/{userID}/invoices", user.HandleInvoices)

	router.NotFound(responder.NotFoundHandler)

	fmt.Println("Starting server on the port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

// HandleRoot will load the root
// func HandleRoot(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	json.NewEncoder(w).Encode("this is the homepage")
// }

func errFatalHandle(err error, message string) {
	if err != nil {
		logrus.WithError(err).Fatalln(message)
	}
}
