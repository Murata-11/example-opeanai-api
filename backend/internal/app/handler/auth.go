package handler

import (
	"app/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func LoginHandler(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	if err := c.Validate(req); err != nil {
		return echo.ErrBadRequest
	}

	if req.Email != "test@example.com" && req.Password != "password" {
		return echo.ErrUnauthorized
	}

	c.SetCookie(pkg.SetAccessToken("dummy_token"))
	c.SetCookie(pkg.SetRefreshToken("dummy_token"))

	return c.JSON(http.StatusOK, nil)
}

func AuthZHandler(c echo.Context) error {
	accessToken, err := c.Cookie("AccessToken")
	if err != nil {
		return echo.ErrForbidden
	}

	if accessToken.Value != "dummy_token" {
		return echo.ErrForbidden
	}

	return c.JSON(http.StatusOK, nil)
}

func RefreshTokenHandler(c echo.Context) error {
	refreshToken, err := c.Cookie("RefreshToken")
	if err != nil {
		return echo.ErrUnauthorized
	}

	if refreshToken.Value != "dummy_token" {
		return echo.ErrUnauthorized
	}

	c.SetCookie(pkg.SetAccessToken("dummy_token"))

	return c.JSON(http.StatusOK, nil)
}

type RefreshTokenForServerSideResponse struct {
	AccessToken string `json:"accessToken"`
}

func RefreshTokenForServerSideHandler(c echo.Context) error {
	refreshToken, err := c.Cookie("RefreshToken")
	if err != nil {
		return echo.ErrUnauthorized
	}

	if refreshToken.Value != "dummy_token" {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, RefreshTokenForServerSideResponse{"dummy_token"})
}
