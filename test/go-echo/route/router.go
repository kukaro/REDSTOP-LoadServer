package route

import (
	"../conf"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	IsOwner bool   `json:"is-owner"`
}

func Init() *echo.Echo {
	e := echo.New()

	switch conf.Conf.Static.Type {
	case conf.BINDATA:
		/*pass*/
	default:
		e.Static("/assets","./assets")
	}
	
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", getRoot)
		v1.POST("/user", postUser)
	}
	return e
}

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func postUser(c echo.Context) error {
	id := c.QueryParam("id")
	var user User
	if strings.Compare(id, "jiharu") == 0 {
		user = User{"jiharu", 26, "male", true}
	} else {
		user = User{id, 22, "female", false}
	}
	return c.JSON(http.StatusOK, user)
}
