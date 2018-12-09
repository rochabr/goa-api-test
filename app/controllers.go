// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "test": Application Controllers
//
// Command:
// $ goagen
// --design=goa-api-test/design
// --out=$(GOPATH)/src/goa-api-test
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Health(*HealthHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	service.Mux.Handle("GET", "/_ah/health", ctrl.MuxHandler("health", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Health", "route", "GET /_ah/health")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger-ui/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/index.html", "route", "GET /swagger-ui/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// TestController is the controller interface for the Test actions.
type TestController interface {
	goa.Muxer
	Create(*CreateTestContext) error
	List(*ListTestContext) error
	Show(*ShowTestContext) error
	Update(*UpdateTestContext) error
}

// MountTestController "mounts" a Test resource controller on the given service.
func MountTestController(service *goa.Service, ctrl TestController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTestContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateTestPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/tests", ctrl.MuxHandler("create", h, unmarshalCreateTestPayload))
	service.LogInfo("mount", "ctrl", "Test", "action", "Create", "route", "POST /tests")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListTestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/tests", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Test", "action", "List", "route", "GET /tests")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/tests/:testId", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Test", "action", "Show", "route", "GET /tests/:testId")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTestContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TestPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/tests/:testId", ctrl.MuxHandler("update", h, unmarshalUpdateTestPayload))
	service.LogInfo("mount", "ctrl", "Test", "action", "Update", "route", "PATCH /tests/:testId")
}

// unmarshalCreateTestPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateTestPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createTestPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTestPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTestPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &testPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}