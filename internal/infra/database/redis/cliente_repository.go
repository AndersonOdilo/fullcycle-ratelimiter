package redis

import (
	"context"
	"encoding/json"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/entity"
	"github.com/redis/go-redis/v9"
)

type ClienteRedisRepository struct {
	RedisClient 	*redis.Client 	
}

func NewClienteRedisRepository(redisClient *redis.Client) *ClienteRedisRepository {
	return &ClienteRedisRepository{
		RedisClient: redisClient,
	}
}


func (r ClienteRedisRepository)	Obtem(ctx context.Context, chave string) (entity.Cliente, error){
	
	val, err := r.RedisClient.Get(ctx, chave).Result()

	if err != nil {

		if (err == redis.Nil) {
			return entity.Cliente{}, nil
		}

		return entity.Cliente{}, err
	}

	cliente := entity.Cliente{}
	err = json.Unmarshal([]byte(val), &cliente)

	if err != nil {
		return entity.Cliente{}, err
	}

	return cliente, nil
}

func (r ClienteRedisRepository)	Grava(ctx context.Context, cliente entity.Cliente) error{

	jsonData, err := json.Marshal(cliente)

	if (err != nil){
		return err
	}

	return r.RedisClient.Set(ctx, cliente.Chave, jsonData, 0).Err()

}