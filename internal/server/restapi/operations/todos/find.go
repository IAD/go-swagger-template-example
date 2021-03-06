// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime/debug"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// FindHandlerFunc turns a function with the right signature into a find handler
type FindHandlerFunc func(*FindParams, NewFindOKFunc, NewFindNotFoundFunc, NewFindInternalServerErrorFunc) middleware.Responder

// Handle executing the request and returning a response
func (fn FindHandlerFunc) Handle(params *FindParams, findOK NewFindOKFunc, findNotFound NewFindNotFoundFunc, findInternalServerError NewFindInternalServerErrorFunc) middleware.Responder {
	return fn(params, findOK, findNotFound, findInternalServerError)
}

// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(*FindParams, NewFindOKFunc, NewFindNotFoundFunc, NewFindInternalServerErrorFunc) middleware.Responder
}

// NewFind creates a new http.Handler for the find operation
func NewFind(ctx *middleware.Context, handler FindHandler) *Find {
	return &Find{Context: ctx, Handler: handler}
}

/*Find swagger:route GET / todos find

Find find API

*/
type Find struct {
	Context *middleware.Context
	Handler FindHandler
}

func (o *Find) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindParams()

	if err := o.Context.BindValidRequest(r, route, Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	// pass predefined values from middleware
	requestCtx := r.Context()
	logger := requestCtx.Value("logger").(*logrus.Entry)
	requestID := requestCtx.Value("requestID").(string)

	// pass body
	Params.RequestBody = requestCtx.Value("body").([]byte)

	Params.Scope = struct {
		Ctx       context.Context
		RequestID string
		Logger    *logrus.Entry
	}{
		Ctx:       requestCtx,
		RequestID: requestID,
		Logger:    logger,
	}

	defer func() {
		if rec := recover(); rec != nil {
			Params.Scope.Logger.Errorf("%s: %s", rec, debug.Stack())

			requestID, ok := requestCtx.Value("requestID").(string)
			if ok {
				rw.Header().Add("requestID", requestID)
			}

			rw.WriteHeader(http.StatusInternalServerError)

			o.Context.Respond(rw, r, route.Produces, route, json.RawMessage([]byte(`{"code":"panic","message":""}`)))
		}
	}()

	res := o.Handler.Handle(Params, NewFindOK, NewFindNotFound, NewFindInternalServerError) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)
}
