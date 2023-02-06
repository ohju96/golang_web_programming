package internal

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

func (controller *Controller) Create(c echo.Context) error {

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "바인딩 에러")
	}
	res, err := controller.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "서비스 로직 에러")
	}

	//response := CreateResponse{
	//	ID:             createResponse.ID,
	//	MembershipType: createResponse.MembershipType,
	//}

	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	res, err := controller.service.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "서비스 로직 에러")
	}

	//response := GetResponse{
	//	ID:             user.ID,
	//	UserName:       user.UserName,
	//	MembershipType: user.MembershipType,
	//}

	return c.JSON(http.StatusOK, res)
}
