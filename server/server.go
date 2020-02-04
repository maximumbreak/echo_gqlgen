package main

import (
	"github.com/99designs/gqlgen/handler"
	test_gqlgen "github.com/beforesecond/gqlgen-todos"
	"github.com/labstack/echo"
)

func graphqlHandler() echo.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(test_gqlgen.NewExecutableSchema(test_gqlgen.Config{Resolvers: &test_gqlgen.Resolver{}}))

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
