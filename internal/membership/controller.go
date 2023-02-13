package membership

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Login(c echo.Context) error {
	var request LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "바인딩 에러")
	}

	login, err := controller.service.Login(request)
	if err != nil {

	}
	return c.JSON(http.StatusOK, login)
}

func (controller *Controller) Create(c echo.Context) error {

	// 바인딩 및 에러 처리
	var request CreateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "바인딩 에러")
	}

	// 로직 및 에러 처리
	response, err := controller.service.Create(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 응답 객체 리턴
	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) GetByID(c echo.Context) error {

	// path variable 바인딩
	request := c.Param("id")

	// 로직 및 에러 처리
	response, err := controller.service.GetByID(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 응답 객체 리턴
	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Update(c echo.Context) error {

	// 객체 바인딩
	var request UpdateRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "바인딩 에러")
	}

	// 로직 처리
	response, err := controller.service.Update(request)
	if err != nil { // 로직 에러
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 응답 객체 리턴
	return c.JSON(http.StatusOK, response)
}
