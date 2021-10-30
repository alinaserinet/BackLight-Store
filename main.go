package main

import (
	"BackLight/server"
	"BackLight/services"
)

func main()  {
	services.Bootstrap()
	server.Init()
}

