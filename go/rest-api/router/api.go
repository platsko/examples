package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"lab/go-rest-api/app/http/controller"
	"lab/go-rest-api/app/http/middleware"
	"lab/go-rest-api/config"
)

var (
	// router keeps chi router instance.
	router = chi.NewRouter()
)

// Start starts the web server with request routing.
func Start() error {
	// Init the api router.
	initRouter()

	// Prepare the host name using the app config.
	conf := config.APP()
	host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	log.Printf("Listen on: %s and serve...\n", host)
	if err := http.ListenAndServe(host, router); err != nil {
		return err
	}

	return nil
}

// initRouter inits routes for the api service.
func initRouter() {
	// Use middleware section.
	router.Use(middleware.ApiRequest)

	// Create api path rotes.
	router.Route("/api", func(r chi.Router) {
		// api version selector.
		r.Route("/v1", func(r chi.Router) {
			// Mount routes for "/client" resource.
			r.Mount("/client", clientResource())
		})
	})
}

// clientResource defines routes for "/client" resource.
func clientResource() chi.Router {
	// Create a router to be mounted to root.
	r := chi.NewRouter()

	// Use middleware section.
	r.Use(middleware.AuthRequest)

	// List of routes of resource.
	r.Post("/", controller.FetchClient)

	return r
}
