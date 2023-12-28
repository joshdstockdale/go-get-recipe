package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

func WithUser(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "meme@yep.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}