package router

import (
	"../conf"
	"../router/api"
	//"../router/web"
	//"../router/web"
	"fmt"
	"github.com/labstack/echo"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	hosts := make(map[string]*Host)

	hosts[conf.Conf.Server.DomainApi] = &Host{api.Routers()}
	//hosts[conf.Conf.Server.DomainWeb] = &Host{web.Routers()}
	//hosts[conf.Conf.Server.DomainWebSocket] = &Host{api.Routers()}
	return hosts
}

func RunSubDomains() {
	e := echo.New()
	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		//u, _ := url.Parse(c.Scheme() + "://" + req.Host)
		if host := hosts[req.Host]; host == nil {
			//fmt.Println("req.Host : " + req.Host)
			//fmt.Println("req.RequestURI : " + req.RequestURI)
			/*error state*/
			err = echo.ErrNotFound
		} else {
			fmt.Println("req.Host : " + req.Host)
			fmt.Println("req.RequestURI : " + req.RequestURI)
			fmt.Println(host)
			host.Echo.ServeHTTP(res, req)
		}
		return
	})
	if err := e.Start(conf.Conf.Server.Addr); err != nil {
		/*error state pass*/
	}
	//go func(){
	//	if err := e.Start(conf.Conf.Server.Addr); err !=nil{
	//		/*error state pass*/
	//	}
	//}()
}
