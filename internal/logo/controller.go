package logo

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller Controller) Get(c echo.Context) error {
	url := "./assets/membership.png"

	file, err := os.Stat(url)
	if err != nil {
		return echo.ErrInternalServerError
	}

	modifiedTime := file.ModTime()
	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedTime.String())))
	c.Response().Header().Set("ETag", etag)
	return c.File(url)
}

func (controller Controller) GetEtag(c echo.Context) error {
	url := "./assets/membership.png"

	file, err := os.Stat(url)
	if err != nil {
		return echo.ErrInternalServerError
	}

	etag := fmt.Sprintf("%x", md5.Sum([]byte(file.ModTime().String()+file.Mode().String()+strconv.FormatInt(file.Size(), 10))))

	if match := c.Request().Header.Get("If-None-Match"); match != "" {
		if match == etag {
			return c.NoContent(http.StatusNotModified)
		}
	}

	c.Response().Header().Set("ETag", etag)
	c.Response().Header().Set("Cache-Control", "max-age=120")
	return c.File(url)
}
