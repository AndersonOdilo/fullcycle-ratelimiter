package database

import (
	"log"
	"os"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/entity"
	internal "github.com/AndersonOdilo/fullcycle-ratelimiter/internal/infra/database/redis"
	"github.com/redis/go-redis/v9"
)

type TipoBancoDeDados int

const (
	REDIS  TipoBancoDeDados = iota + 1
)

func FabricaClienteRepository(tipoBancoDeDados TipoBancoDeDados) entity.ClienteRepositoryInterface {

	if (tipoBancoDeDados == REDIS){
		
		rdb :=  redis.NewClient(&redis.Options{
			Addr: getRedisUrlAddress(),
			DB: 0,  
		})

			return internal.NewClienteRedisRepository(rdb)
	}

	return nil

}


func getRedisUrlAddress() string {
	redisUrlAddress := os.Getenv("REDIS_URL_ADDRESS")

	if redisUrlAddress == "" {
		log.Fatal("Não foi localizada a variavel de ambiente REDIS_URL_ADDRESS")
	}

	return redisUrlAddress
}
