package web

import (
	"../../conf"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	IsOwner bool   `json:"is-owner"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Routers() *echo.Echo {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("template/layout/web/*.html")),
	}
	e.Renderer = t

	// set static
	switch conf.Conf.Static.Type {
	case conf.BINDATA:
		/*pass*/
	default:
		e.Static("/assets", "./assets")
	}

	v1 := e.Group("")
	{
		v1.GET("/", getRootHandler)
	}
	return e
}
