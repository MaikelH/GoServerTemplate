package httpserver

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/rs/cors"
	"goservertemplate/servicecontainer"
	"goservertemplate/services"
	"log/slog"
)

func StartHTTPServer(c *servicecontainer.Container) error {
	slog.Info("Initializing HTTP server")

	if c.Config.ListenAddress == "" {
		slog.Error("missing listen address")
		return services.NewServiceError(services.ErrServerError, nil, "missing listen address")
	}

	// TODO: Temporarily allow all CORS requests for development, this will be locked down later
	s := fuego.NewServer(fuego.WithAddr(c.Config.ListenAddress),
		fuego.WithSecurity(openapi3.SecuritySchemes{
			"bearerAuth": &openapi3.SecuritySchemeRef{
				Value: openapi3.NewSecurityScheme().
					WithType("http").
					WithScheme("bearer").
					WithBearerFormat("JWT").
					WithDescription("Enter your JWT token in the format: <token>"),
			},
		}),
		fuego.WithCorsMiddleware(cors.AllowAll().Handler))

	if c.Config.OpenAPIAddress != "" {
		s.OpenAPI.Description().Servers = []*openapi3.Server{{URL: c.Config.OpenAPIAddress}}
	} else {
		s.OpenAPI.Description().Servers = []*openapi3.Server{{URL: c.Config.ListenAddress}}
	}
	s.OpenAPI.Description().Info.Title = "GoServerTemplate"
	s.OpenAPI.Description().Info.Description = "This is the autogenerated OpenAPI documentation for GoServerTemplate."

	_ = fuego.Group(s, "/api/v1")

	err := s.Run()
	if err != nil {
		slog.Error("service_error starting HTTP server", "service_error", err)
		return services.NewServiceError(services.ErrServerError, err, "service_error starting HTTP server")
	}

	return nil
}
