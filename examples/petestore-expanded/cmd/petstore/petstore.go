// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml
//
// The code under api/petstore/ has been generated from that specification.
package main

import (
    "flag"
    "fmt"

    "github.com/deepmap/oapi-codegen/pkg/util"
    "os"

    "github.com/labstack/echo/v4"
    echomiddleware "github.com/labstack/echo/v4/middleware"

    "github.com/deepmap/oapi-codegen/examples/petestore-expanded/api"
    "github.com/deepmap/oapi-codegen/examples/petestore-expanded/internal"
    "github.com/deepmap/oapi-codegen/pkg/middleware"
)

func main() {
    var port = flag.Int("port", 8080, "Port for test HTTP server")
    var specPath = flag.String("spec", "../../api/petstore-expanded.yaml",
        "Path to OpenAPI specification for this server")
    flag.Parse()

    swagger, err := util.LoadSwagger(*specPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
        os.Exit(1)
    }

    // Create an instance of our handler which satisfies the generated interface
    petStore := internal.NewPetStore()

    // This is how you set up a basic Echo router
    e := echo.New()
    // Log all requests
    e.Use(echomiddleware.Logger())
    // Use our validation middleware to check all requests against the
    // OpenAPI schema.
    e.Use(middleware.OapiRequestValidator(swagger))

    // We now register our petStore above as the handler for the interface
    api.RegisterHandlers(e, petStore)

    // And we serve HTTP until the world ends.
    e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
