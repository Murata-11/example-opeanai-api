package router

import (
	"app/internal/app/handler"
	"app/validate"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3001", "http://localhost:3000", "http://frontend:3001", "http://frontend:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.Validator = validate.NewValidator()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/login", handler.LoginHandler)
	e.GET("/authz", handler.AuthZHandler)
	e.GET("/refresh", handler.RefreshTokenHandler)
	e.GET("/refresh-server", handler.RefreshTokenForServerSideHandler)

	e.POST("/brainstorm", handler.BrainstormHandler)
}
