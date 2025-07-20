package main

import (
	"fmt"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/configs"
	web "github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/web/handler"
	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/web/webserver"
)


func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Starting web server on port", configs.WebServerPort)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webTempHandler := web.NewWebHellopHandler();
	webserver.AddHandler("GET /", webTempHandler.Get)
	webserver.Start()
}
