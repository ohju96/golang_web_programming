package internal

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_web_programming/internal/logo"
	"golang_web_programming/internal/memberships"
	"golang_web_programming/internal/user"
	"log"
	"net/http"
)

const _defaultPort = 8080

type Server struct {
	logoController       logo.Controller
	membershipController memberships.Controller
	userController       user.Controller
	userMiddleware       user.Middleware
}

func NewDefaultServer() *Server {
	data := map[string]memberships.Membership{}
	membershipRepository := memberships.NewRepository(data)
	membershipService := memberships.NewService(*membershipRepository)
	membershipController := memberships.NewController(*membershipService)
	return &Server{
		membershipController: *membershipController,
		logoController:       *logo.NewController(),
		userController:       *user.NewController(*user.NewService(user.DefaultSecret)),
		userMiddleware:       *user.NewMiddleware(*membershipRepository),
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, context echo.Context) {
		if errors.Is(err, user.ErrInvalidPassword) {
			context.JSON(http.StatusBadRequest, map[string]string{"message": "invalid password"})
			return
		}
		e.DefaultHTTPErrorHandler(err, context)
	}

	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.membershipController, s.userMiddleware)
	RouteLogo(g, s.logoController)
	RouteUser(g, s.userController)
}

func RouteMemberships(e *echo.Group, c memberships.Controller, userMiddleware user.Middleware) {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})
	e.GET("/memberships/:id", c.GetByID, jwtMiddleware, userMiddleware.ValidateAdmin)
	e.POST("/memberships", c.Create)
}

func RouteLogo(e *echo.Group, c logo.Controller) {
	e.GET("/logo", c.GetEtag)
}

func RouteUser(e *echo.Group, c user.Controller) {
	e.POST("/login", c.Login)
}
