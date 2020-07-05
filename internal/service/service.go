package service

import (
	"github.com/IAD/go-swagger-template-example/internal/server/restapi"
	"github.com/IAD/go-swagger-template-example/internal/server/restapi/operations"
	"github.com/IAD/go-swagger-template-example/internal/server/restapi/operations/todos"
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"
)

func PrepareServer(port int64, logger *logrus.Entry) (*restapi.Server, error) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	h := NewHandlers()

	api := operations.NewTodoAPI(swaggerSpec)

	// Handler for POST /
	api.TodosAddOneHandler = todos.AddOneHandlerFunc(h.AddOneHandler)
	// Handler for DELETE /{id}
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(h.DestroyOneHandler)
	// Handler for GET /
	api.TodosFindHandler = todos.FindHandlerFunc(h.FindHandler)
	// Handler for PUT /{id}
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(h.UpdateOneHandler)

	server := restapi.NewServerWithMiddleware(api, "todo", logger)
	server.Port = int(port)

	return server, nil
}
