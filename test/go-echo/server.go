package main

import (
	"../go-echo/route"
)

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":1323"))
}

