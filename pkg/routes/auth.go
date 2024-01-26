package routes

import (
	"book-crud/pkg/controllers"
	"github.com/labstack/echo/v4"
)

// AuthRoutes stores controller and echo instance for authentication.
type authRoutes struct {
	echo    *echo.Echo
	authCtr controllers.AuthController
}

// NewAuthRoutes returns a new instance of the AuthRoutes struct.
func AuthRoutes(echo *echo.Echo, authCtr controllers.AuthController) *authRoutes {
	return &authRoutes{
		echo:    echo,
		authCtr: authCtr,
	}
}

// InitAuthRoutes initializes the authentication routes.
func (routes *authRoutes) InitAuthRoutes() {
	e := routes.echo

	e.POST("/login", routes.authCtr.Login)
	e.POST("/signup", routes.authCtr.Signup)
}