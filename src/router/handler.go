package router

import (
	_ "dialog-interview/docs/app"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Dialog-Interview Swagger API
// @version 1.0
// @description API Documentation

// @contact.name Luccas Machado
// @contact.url linkedin
// @contact.email luccaa.chagas23@gmail.com

// @host
// @BasePath /
func Handler() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(CorsConfig))
	e.POST("/auth", Auth)

	authors := e.Group("/authors")
	authors.Use(middleware.JWTWithConfig(Config))

	authors.GET("/:page", GetAuthors)
	authors.GET("", GetAuthorsByName)

	books := e.Group("/books")
	books.Use(middleware.JWTWithConfig(Config))

	books.POST("/create", CreateBooks)
	books.GET("/read", ReadBooks)
	books.PUT("/update", UpdateBooks)
	books.DELETE("/delete/:id", DeleteBooks)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
