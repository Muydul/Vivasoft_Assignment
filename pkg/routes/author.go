package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"
	// "fmt"
	// "net/http"

	"github.com/labstack/echo/v4"
)

type authorRoutes struct {
	echo      *echo.Echo
	authorCtr controllers.AuthorController
}

func AuthorRoutes(echo *echo.Echo, authorCtr controllers.AuthorController) *authorRoutes {
	return &authorRoutes{
		echo:      echo,
		authorCtr: authorCtr,
	}
}

func (ac *authorRoutes) InitAuthorRoute() {
	e := ac.echo
	ac.initAuthorRoutes(e)
}

func (ac *authorRoutes) initAuthorRoutes(e *echo.Echo) {
	//grouping route endpoints
	author := e.Group("/bookstore")

	// author.GET("/ping", Pong)
	author.GET("/author", ac.authorCtr.GetAuthor)

	//initializing http methods - routing endpoints and their handlers
	author.Use(middlewares.ValidateToken)
	author.POST("/author", ac.authorCtr.CreateAuthor)
	author.PUT("/author/:authorID", ac.authorCtr.UpdateAuthor)
	author.DELETE("/author/:authorID", ac.authorCtr.DeleteAuthor)
}

// func Pong(ctx echo.Context) error {
// 	fmt.Println("Pong")
// 	return ctx.JSON(http.StatusOK, "Pong")
// }
