package memberships

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {

	var request CreateRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "바인딩 에러 입니다.")
	}

	response, err := controller.service.Create(request)
	if err != nil {
		// todo 알맞는 에러 넣기
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}

func (controller *Controller) Update(c echo.Context) error {
	id := c.Get("user")
	fmt.Println("id : ", id)
	return c.JSON(http.StatusOK, id)
}

/*func (controller *Controller) Delete(c echo.Context) error {
	return
}*/
