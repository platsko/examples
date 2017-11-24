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
	// init the api router
	initRouter()

	// prepare the host name using the app config
	conf := config.App()
	host := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	log.Printf("Listen on: %s and serve...\n", host)
	if err := http.ListenAndServe(host, router); err != nil {
		return err
	}

	return nil
}

// initRouter inits routes for the api service.
func initRouter() {
	// use middleware section
	router.Use(middleware.ApiRequest)

	// create api path rotes
	router.Route("/api", func(r chi.Router) {
		// api version selector
		r.Route("/v1", func(r chi.Router) {
			// mount routes for "/client" resource
			r.Mount("/client", clientResource())
			// mount routes for "/document" resource
			r.Mount("/document", documentResource())
		})
	})
}

// clientResource defines routes for "/client" resource.
func clientResource() chi.Router {
	// create a router to be mounted to root
	r := chi.NewRouter()

	// use middleware section
	r.Use(middleware.AuthRequest)

	// list of routes of resource
	r.Post("/", controller.FetchClient)

	return r
}

// documentResource defines routes for "/document" resource.
func documentResource() chi.Router {
	// create a router to be mounted to root
	r := chi.NewRouter()

	// use middleware section
	r.Use(middleware.AuthRequest)

	// list of routes of resource
	r.Post("/", controller.Document)

	return r
}
