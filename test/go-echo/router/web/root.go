package web

import (
	"../../conf"
	"github.com/labstack/echo"
	"net/http"
)

func getRootHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title": conf.Conf.App.Name,
		"owner": conf.Conf.App.Owner,
	})
}
