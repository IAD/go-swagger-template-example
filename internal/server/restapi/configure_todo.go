// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/IAD/go-swagger-template-example/internal/server/restapi/operations"
	"github.com/IAD/go-swagger-template-example/internal/server/restapi/operations/todos"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
)

//go:generate swagger generate server --target ../../server --name Todo --spec ../../../swagger.yml --template-dir /tmp/templates --exclude-main

func configureFlags(api *operations.TodoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Handler for POST /
	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params *todos.AddOneParams,
		addOneCreated todos.NewAddOneCreatedFunc,
		addOneInternalServerError todos.NewAddOneInternalServerErrorFunc,
	) middleware.Responder {
		return middleware.NotImplemented("operation todos.AddOne has not yet been implemented")
	})
	// Handler for DELETE /{id}
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params *todos.DestroyOneParams,
		destroyOneNoContent todos.NewDestroyOneNoContentFunc,
		destroyOneNotFound todos.NewDestroyOneNotFoundFunc,
		destroyOneInternalServerError todos.NewDestroyOneInternalServerErrorFunc,
	) middleware.Responder {
		return middleware.NotImplemented("operation todos.DestroyOne has not yet been implemented")
	})
	// Handler for GET /
	api.TodosFindHandler = todos.FindHandlerFunc(func(params *todos.FindParams,
		findOK todos.NewFindOKFunc,
		findNotFound todos.NewFindNotFoundFunc,
		findInternalServerError todos.NewFindInternalServerErrorFunc,
	) middleware.Responder {
		return middleware.NotImplemented("operation todos.Find has not yet been implemented")
	})
	// Handler for PUT /{id}
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params *todos.UpdateOneParams,
		updateOneOK todos.NewUpdateOneOKFunc,
		updateOneNotFound todos.NewUpdateOneNotFoundFunc,
		updateOneInternalServerError todos.NewUpdateOneInternalServerErrorFunc,
	) middleware.Responder {
		return middleware.NotImplemented("operation todos.UpdateOne has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/version" {
			serveVersion(w)
			return
		}

		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusOK)
			return
		}

		cors.AllowAll().Handler(handler).ServeHTTP(w, r)
	})
}

func serveVersion(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
