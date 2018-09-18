package route

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func Init() *echo.Echo {
	e := echo.New()
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
	user := User{"jiharu", 26, "male"}
	return c.JSON(http.StatusOK, user)
}
