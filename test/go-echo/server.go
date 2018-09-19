package main

import (
	"../go-echo/conf"
	"../go-echo/route"
	"fmt"
)

func main() {
	if err := conf.Init(""); err == nil {
		fmt.Println("config success")
	}
	router := route.Init()
	router.Logger.Fatal(router.Start(":1323"))
}
