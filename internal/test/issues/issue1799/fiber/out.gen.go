// Package fiber provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package fiber

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /object)
	GetObject(c *fiber.Ctx) error

	// (GET /object-multibody)
	GetObjectMultibody(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// GetObject operation middleware
func (siw *ServerInterfaceWrapper) GetObject(c *fiber.Ctx) error {

	return siw.Handler.GetObject(c)
}

// GetObjectMultibody operation middleware
func (siw *ServerInterfaceWrapper) GetObjectMultibody(c *fiber.Ctx) error {

	return siw.Handler.GetObjectMultibody(c)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(fiber.Handler(m))
	}

	router.Get(options.BaseURL+"/object", wrapper.GetObject)

	router.Get(options.BaseURL+"/object-multibody", wrapper.GetObjectMultibody)

}

type GetObjectRequestObject struct {
}

type GetObjectResponseObject interface {
	VisitGetObjectResponse(ctx *fiber.Ctx) error
}

type GetObject200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse string

func (response GetObject200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse) VisitGetObjectResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetObjectMultibodyRequestObject struct {
}

type GetObjectMultibodyResponseObject interface {
	VisitGetObjectMultibodyResponse(ctx *fiber.Ctx) error
}

type GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse string

func (response GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3OrgnsactivitystreamsResponse) VisitGetObjectMultibodyResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3Orgnsactivitystreams2Response string

func (response GetObjectMultibody200ApplicationLdPlusJSONProfilehttpswwwW3Orgnsactivitystreams2Response) VisitGetObjectMultibodyResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams2\"")
	ctx.Status(200)

	return ctx.JSON(&response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /object)
	GetObject(ctx context.Context, request GetObjectRequestObject) (GetObjectResponseObject, error)

	// (GET /object-multibody)
	GetObjectMultibody(ctx context.Context, request GetObjectMultibodyRequestObject) (GetObjectMultibodyResponseObject, error)
}

type StrictHandlerFunc func(ctx *fiber.Ctx, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetObject operation middleware
func (sh *strictHandler) GetObject(ctx *fiber.Ctx) error {
	var request GetObjectRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetObject(ctx.UserContext(), request.(GetObjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetObject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetObjectResponseObject); ok {
		if err := validResponse.VisitGetObjectResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetObjectMultibody operation middleware
func (sh *strictHandler) GetObjectMultibody(ctx *fiber.Ctx) error {
	var request GetObjectMultibodyRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetObjectMultibody(ctx.UserContext(), request.(GetObjectMultibodyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetObjectMultibody")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetObjectMultibodyResponseObject); ok {
		if err := validResponse.VisitGetObjectMultibodyResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
