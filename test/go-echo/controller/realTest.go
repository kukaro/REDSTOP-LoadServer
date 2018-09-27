package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ApiRealTestData struct {
	Method string `json:"method"`
	Rtype  string `json:"rtype"`
	Rurl   string `json:"rurl"`
	Body   string `json:"body"`
}

func RealTest(c echo.Context) error {
	method := strings.ToUpper(c.Param("method"))
	rtype := c.Param("rtype")
	rurl, _ := url.QueryUnescape(c.Param("rurl"))
	//fmt.Println("huh")
	//fmt.Println(test)
	//fmt.Println(rurl)
	//isComplete := make(chan bool)
	//isComplete := make(chan bool)
	body := make(chan string)
	go func() {
		/*
			test : http://localhost:1323/api/v1/api/real-test/get/http/localhost:3000%2fapi%2fjson-test%2fsingle-json%2f
		 */
		resp, _ := http.Get(rtype + "://" + rurl)
		data, _ := ioutil.ReadAll(resp.Body)
		body <- string(data)
	}()

	apiRealTestData := ApiRealTestData{rurl, rtype, method, <-body}
	jsonData, _ := json.Marshal(apiRealTestData)

	fmt.Println(string(jsonData))
	return c.String(http.StatusOK, string(jsonData))
}
