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
	Method     string      `json:"method"`
	Rtype      string      `json:"rtype"`
	Rurl       string      `json:"rurl"`
	Body       string      `json:"body"`
	Header     interface{} `json:"header"`
	StatusCode int         `json:"status-code"`
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
	header := make(chan interface{})
	statusCode := make(chan int)
	go func() {
		/*
			test : http://localhost:1323/api/v1/api/real-test/get/http/localhost:3000%2fapi%2fjson-test%2fsingle-json%2f
		 */
		resp, _ := http.Get(rtype + "://" + rurl)
		data, _ := ioutil.ReadAll(resp.Body)
		body <- string(data)
		header <- resp.Header
		statusCode <- resp.StatusCode
	}()

	apiRealTestData := ApiRealTestData{method, rtype, rurl, <-body, <-header, <-statusCode}
	jsonData, _ := json.Marshal(apiRealTestData)

	fmt.Println(string(jsonData))
	return c.String(http.StatusOK, string(jsonData))
}
