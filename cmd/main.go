package main

import (
	"fmt"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/configs"
	internal "github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/database/redis"
	web "github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/web/handler"
	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/web/webserver"
	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/usecase"
	"github.com/redis/go-redis/v9"
)


func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	rdb :=  redis.NewClient(&redis.Options{
        Addr: configs.RedisUrlAddress,
        DB: 0,  
    })

	
	
	fmt.Println("Starting web server on port", configs.WebServerPort)

	clienteRedisRepository := internal.NewClienteRedisRepository(rdb)
	rateLimiterUseCase := usecase.NewRateLimiterUseCase(clienteRedisRepository);

	webserver := webserver.NewWebServer(configs.WebServerPort, rateLimiterUseCase)
	webTempHandler := web.NewWebHellopHandler();
	webserver.AddHandler("GET /", webTempHandler.Get)
	webserver.Start()
}
