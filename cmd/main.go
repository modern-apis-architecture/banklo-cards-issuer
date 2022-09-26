package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/modern-apis-architecture/banklo-cards-issuer/internal/api"
	"net/http"
	"os"
)

func main() {
	var port = flag.Int("port", 7777, "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	cardIssuer, err := buildAppContainer()
	e := echo.New()
	e.Use(echomiddleware.Logger())
	api.RegisterHandlers(e, cardIssuer)
	e.GET("/health", health)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "UP")
}
