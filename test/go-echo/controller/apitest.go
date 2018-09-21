package controller

import (
	"../conf"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func ApiTestHandler(c echo.Context) error {
	cnt, _ := strconv.Atoi(c.Param("count"))
	start := time.Now().Nanosecond()
	for i := 0; i < cnt; i++ {
		http.Get(conf.Conf.TestServer.TestDomain + "api/json-test/single-json/")
		//resp, _ := http.Get(conf.Conf.TestServer.TestDomain + "api/json-test/single-json/")
		//fmt.Println(resp.Body)
	}
	end := time.Now().Nanosecond()

	//fmt.Sprintf가 더 낫다고 하네요... 무슨 C인줄 ㅡㅡ
	jsonData, _ := json.Marshal(map[string]interface{}{
		"count":           cnt,
		"sumResponseTime": strconv.Itoa((end-start)/1000) + "us",
		"avgResponseTime": strconv.Itoa((end-start)/cnt/1000) + "us",
	})
	return c.String(http.StatusOK, string(jsonData))
}

func ApiTestOriginHandler(c echo.Context) error {
	cnt, _ := strconv.Atoi(c.Param("count"))
	start := time.Now().UnixNano()
	isComplete := make(chan bool, cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			http.Get(conf.Conf.TestServer.TestDomain + "api/json-test/single-json/")
			isComplete <- true
		}()
	}
	for i := 0; i < cnt; i++ {
		<-isComplete
		//fmt.Printf("goroutin : %d번째 호출\n", i);
	}
	end := time.Now().UnixNano()
	fmt.Println(end)
	fmt.Println(start)

	//fmt.Sprintf가 더 낫다고 하네요... 무슨 C인줄 ㅡㅡ
	jsonData, _ := json.Marshal(map[string]interface{}{
		"count":           cnt,
		"sumResponseTime": strconv.Itoa(int((end-start)/1000)) + "us",
		"avgResponseTime": strconv.Itoa((int(end-start)/cnt/1000)) + "us",
	})
	return c.String(http.StatusOK, string(jsonData))
}
