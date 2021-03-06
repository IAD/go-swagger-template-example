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

// UpdateOneHandlerFunc turns a function with the right signature into a update one handler
type UpdateOneHandlerFunc func(*UpdateOneParams, NewUpdateOneOKFunc, NewUpdateOneNotFoundFunc, NewUpdateOneInternalServerErrorFunc) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateOneHandlerFunc) Handle(params *UpdateOneParams, updateOneOK NewUpdateOneOKFunc, updateOneNotFound NewUpdateOneNotFoundFunc, updateOneInternalServerError NewUpdateOneInternalServerErrorFunc) middleware.Responder {
	return fn(params, updateOneOK, updateOneNotFound, updateOneInternalServerError)
}

// UpdateOneHandler interface for that can handle valid update one params
type UpdateOneHandler interface {
	Handle(*UpdateOneParams, NewUpdateOneOKFunc, NewUpdateOneNotFoundFunc, NewUpdateOneInternalServerErrorFunc) middleware.Responder
}

// NewUpdateOne creates a new http.Handler for the update one operation
func NewUpdateOne(ctx *middleware.Context, handler UpdateOneHandler) *UpdateOne {
	return &UpdateOne{Context: ctx, Handler: handler}
}

/*UpdateOne swagger:route PUT /{id} todos updateOne

UpdateOne update one API

*/
type UpdateOne struct {
	Context *middleware.Context
	Handler UpdateOneHandler
}

func (o *UpdateOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateOneParams()

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

	res := o.Handler.Handle(Params, NewUpdateOneOK, NewUpdateOneNotFound, NewUpdateOneInternalServerError) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)
}
