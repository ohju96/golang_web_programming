package memberships

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

// Create godoc
// @Summary      멤버십 생성
// @Description  멤버십을 생성합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @Param        requestBody  body      CreateRequest  true  "user_name:사용자의 이름, membership_type:naver,toss,pacyco 중 하나"
// @Success      201          {object}  CreateResponse
// @Router       /v1/memberships [post]
func (controller *Controller) Create(c echo.Context) error {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	res, err := controller.service.Create(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

// GetByID godoc
// @Summary      멤버십 정보 단건 조회
// @Description  멤버십 정보를 조회합니다. (상세 설명)
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id             path      string  true  "Membership uuid"
// @Success      200            {object}  GetResponse
// @Failure      400            {object}  Fail400GetResponse
// @Router       /v1/memberships/{id} [get]
func (controller *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")
	res, err := controller.service.GetByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary      멤버십 정보 수정
// @Description  멤버십 정보를 수정합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id             path      string  true  "Membership uuid"
// @Param        UpdateRequest  body      UpdateRequest  true  "user_name:사용자의 이름, membership_type:naver,toss,pacyco 중 하나"
// @Success      200            {object}  UpdateResponse
// @Failure      400            {object}  Fail400GetResponse
// @Router       /v1/memberships/{id} [put]
func (controller *Controller) Update(c echo.Context) error {
	id := c.Param("id")
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	res, err := controller.service.Update(id, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary      멤버십 정보 삭제
// @Description  멤버십 정보를 삭제합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id             path      string  true  "Membership uuid"
// @Success      200            {object}  nil
// @Failure      400            {object}  Fail400GetResponse
// @Router       /v1/memberships/{id} [delete]
func (controller *Controller) Delete(c echo.Context) error {
	id := c.Param("id")
	err := controller.service.Delete(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, nil)
}
