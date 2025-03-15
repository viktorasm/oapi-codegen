// Package echo provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package echo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /object)
	GetObject(ctx echo.Context) error

	// (GET /object-multibody)
	GetObjectMultibody(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetObject converts echo context to params.
func (w *ServerInterfaceWrapper) GetObject(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetObject(ctx)
	return err
}

// GetObjectMultibody converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectMultibody(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetObjectMultibody(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/object", wrapper.GetObject)
	router.GET(baseURL+"/object-multibody", wrapper.GetObjectMultibody)

}

type GetObjectRequestObject struct {
}

type GetObjectResponseObject interface {
	VisitGetObjectResponse(w http.ResponseWriter) error
}

type GetObject200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse string

func (response GetObject200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse) VisitGetObjectResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetObjectMultibodyRequestObject struct {
}

type GetObjectMultibodyResponseObject interface {
	VisitGetObjectMultibodyResponse(w http.ResponseWriter) error
}

type GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse string

func (response GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse) VisitGetObjectMultibodyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3Orgnsactivitystreams2Response string

func (response GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3Orgnsactivitystreams2Response) VisitGetObjectMultibodyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams2\"")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /object)
	GetObject(ctx context.Context, request GetObjectRequestObject) (GetObjectResponseObject, error)

	// (GET /object-multibody)
	GetObjectMultibody(ctx context.Context, request GetObjectMultibodyRequestObject) (GetObjectMultibodyResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetObject operation middleware
func (sh *strictHandler) GetObject(ctx echo.Context) error {
	var request GetObjectRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetObject(ctx.Request().Context(), request.(GetObjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetObject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetObjectResponseObject); ok {
		return validResponse.VisitGetObjectResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetObjectMultibody operation middleware
func (sh *strictHandler) GetObjectMultibody(ctx echo.Context) error {
	var request GetObjectMultibodyRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetObjectMultibody(ctx.Request().Context(), request.(GetObjectMultibodyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetObjectMultibody")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetObjectMultibodyResponseObject); ok {
		return validResponse.VisitGetObjectMultibodyResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
