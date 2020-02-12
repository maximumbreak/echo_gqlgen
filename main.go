package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/beforesecond/gqlgen-todos/generated"
	"github.com/beforesecond/gqlgen-todos/resolver"
	"github.com/labstack/echo"
)

func graphqlHandler() echo.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil

	}

}

// Defining the Playground handler
func playgroundHandler() echo.HandlerFunc {

	h := handler.Playground("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

}

func main() {
	e := echo.New()
	e.GET("/", playgroundHandler())
	e.POST("/query", graphqlHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
