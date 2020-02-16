package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/handler"
	"github.com/beforesecond/gqlgen-todos/generated"
	"github.com/beforesecond/gqlgen-todos/resolver"
	"github.com/beforesecond/gqlgen-todos/service"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
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

	switch os := strings.ToUpper(os.Getenv("ENV")); os {
	case "PROD":
		viper.SetConfigFile("./configs/env.production.yaml")
	case "DEV":
		viper.SetConfigFile("./configs/env.development.yaml")
	default:
		viper.SetConfigFile("./configs/env.local.yaml")
		fmt.Printf("%s.\n", os)
	}

	err := viper.ReadInConfig()

	// Handle errors reading the config file
	if err != nil {
		log.Fatal("Fatal error config file", err)
	}

	port := viper.GetString("app.port")
	e := echo.New()
	// e.Use(
	// 	middleware.Recover(),
	// 	middleware.Secure(),
	// 	middleware.Logger(),
	// 	middleware.Gzip(),
	// 	middleware.BodyLimit("2M"),
	// 	// middleware.CORSWithConfig(middleware.CORSConfig{
	// 	// 	// AllowOrigins: []string{
	// 	// 	// 	"http://localhost:8080",
	// 	// 	// },
	// 	// 	AllowHeaders: []string{
	// 	// 		echo.HeaderOrigin,
	// 	// 		echo.HeaderContentLength,
	// 	// 		echo.HeaderAcceptEncoding,
	// 	// 		echo.HeaderContentType,
	// 	// 		echo.HeaderAuthorization,
	// 	// 	},
	// 	// 	AllowMethods: []string{
	// 	// 		echo.GET,
	// 	// 		echo.POST,
	// 	// 	},
	// 	// 	MaxAge: 3600,
	// 	// }),
	// )

	// Health check
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/", playgroundHandler())
	e.POST("/query", graphqlHandler())

	// Register services
	service.Auth(e.Group("/auth"))

	e.Logger.Fatal(e.Start(":" + port))
}
