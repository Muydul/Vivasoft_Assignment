package containers

import (
	"book-crud/pkg/config"
	"book-crud/pkg/connection"
	"book-crud/pkg/controllers"
	"book-crud/pkg/repositories"
	"book-crud/pkg/routes"
	"book-crud/pkg/services"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	//config initialization
	config.SetConfig()

	//database initializations
	db := connection.GetDB()

	// repository initialization
	bookRepo := repositories.BookDBInstance(db)
	authorRepo := repositories.AuthorDBInstance(db)
	userRepo := repositories.UserDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo)
	authorService := services.AuthorServiceInstance(authorRepo)
	authService := services.AuthServiceInstance(userRepo)

	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService)
	authCtr := controllers.NewAuthController(authService)

	//route initialization
	b := routes.BookRoutes(e, bookCtr)
	c := routes.AuthorRoutes(e, authorCtr)
	a := routes.AuthRoutes(e, authCtr)

	b.InitBookRoute()
	c.InitAuthorRoute()
	a.InitAuthRoutes()
	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
