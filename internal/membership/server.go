package membership

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

const _defaultPort = 8080

type Server struct {
	controller Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	service := NewService(*NewRepository(data))
	controller := NewController(*service)
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.controller)
}

func RouteMemberships(e *echo.Group, c Controller) {

	e.Use(
		// Http 로깅 미들웨어 설정
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogStatus: true,
			LogURI:    true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				fmt.Printf("httpMethod==%v, http-status-code==%v, uri==%v \n", c.Request().Method, v.Status, v.URI)
				return nil
			},
		}),
		// Request, Response 로깅 미들웨어 설정
		middleware.BodyDump(func(ctx echo.Context, requestBody []byte, responseBody []byte) {
			switch {
			case len(requestBody) > 0 && len(responseBody) > 0:
				fmt.Fprintf(os.Stdout, "requestBody = %s responseBody = %s\n", requestBody, responseBody)
			case len(responseBody) > 0:
				fmt.Fprintf(os.Stdout, "responseBody = %s", responseBody)
			case len(requestBody) > 0:
				fmt.Fprintf(os.Stdout, "requestBody = %s", requestBody)
			}
		}),
	)

	e.GET("/memberships/:id", c.GetByID) // 멤버십 조회
	e.POST("/memberships", c.Create)     // 멤버십 생성
	e.PUT("/memberships", c.Update)      // 멤버십 수정
	e.POST("/memberships/logins", c.Login)
}
