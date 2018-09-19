package main

import (
	"../go-echo/router"
	"../go-echo/conf"
	"fmt"
)

func main() {
	if err := conf.Init(""); err == nil {
		fmt.Println("config success")
	}
	router.RunSubDomains()
	//router := router.Init()
	//router.Logger.Fatal(router.Start(":1323"))
}
