package user

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang_web_programming/internal/memberships"
)

type Middleware struct {
	membershipRepository memberships.Repository
}

func NewMiddleware(membershipRepository memberships.Repository) *Middleware {
	return &Middleware{membershipRepository: membershipRepository}
}

func (m Middleware) ValidateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		if !claims.IsAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func (m Middleware) ValidateMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*Claims)
		if claims.IsAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
