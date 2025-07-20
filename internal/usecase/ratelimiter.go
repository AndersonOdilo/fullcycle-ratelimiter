package usecase

import (
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/entity"
)

type AcessoLiberado bool

type RateLimiterUseCase struct {
	ClienteRepositoryInterface 	entity.ClienteRepositoryInterface 	
	mutex 						sync.Mutex	
}

func NewRateLimiterUseCase(clienteRepositoryInterface entity.ClienteRepositoryInterface) *RateLimiterUseCase {
	return &RateLimiterUseCase{
		ClienteRepositoryInterface: clienteRepositoryInterface,
	}
}

func (r *RateLimiterUseCase) Execute(ctx context.Context, ip string, token string, unixTimestampRequest int64) (AcessoLiberado, error)  {

	var chave string
	var nrMaximoRequestPorSegundo int64 
	var duracaoBloqueio time.Duration

	if (token != ""){

		chave = token
		nrMaximoRequestPorSegundo = getNrMaximoRequestPorSegundoToken(chave)
		duracaoBloqueio = getDuracaoBloqueioToken(chave)

	} else{

		chave = ip
		nrMaximoRequestPorSegundo = getNrMaximoRequestPorSegundoIP()
		duracaoBloqueio = getDuracaoBloqueioIP()
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	cliente, err := r.ClienteRepositoryInterface.Obtem(ctx, chave);

	if err != nil {
		return false, err
	}

	//primeiro acesso
	if (cliente.Chave == ""){

		cliente.Chave = chave
		cliente.InsereNovaRequest(unixTimestampRequest)
		
		err := r.ClienteRepositoryInterface.Grava(ctx, cliente)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	// verificar se o cliente tem um boqueio
	if (cliente.UnixBloqueio > 0){

		// verificar se o bloqueio ainda é valido
		if (cliente.UnixBloqueio >= time.Now().UnixMilli()){

			return false, nil

		} else {
			//inativa bloqueio
			cliente.UnixBloqueio = 0
			cliente.InsereNovaRequest(unixTimestampRequest)

			err := r.ClienteRepositoryInterface.Grava(ctx, cliente)

			if err != nil {
				return false, err
			}

			return true, nil
		}
	}

	//verifica se o número maximo de request do ultimo segundo não excedeu o limite
	if (cliente.CalculaNumeroRequestPorSegundo() < nrMaximoRequestPorSegundo){

		cliente.InsereNovaRequest(unixTimestampRequest)
		
		err := r.ClienteRepositoryInterface.Grava(ctx, cliente)

		if err != nil {
			return false, err
		}

		return true, nil
	
	} else {
		//excedeu o limite lança um bloqueio
		cliente.UnixBloqueio = unixTimestampRequest + duracaoBloqueio.Milliseconds()

		err := r.ClienteRepositoryInterface.Grava(ctx, cliente) 

		if err != nil {
			return false, err
		}

		return false, nil
	}
}

func getNrMaximoRequestPorSegundoToken(token string) int64 {
	nrMaximoRequestPorSegundo := os.Getenv("NR_MAXIMO_REQUEST_POR_SEGUNDO_TOKEN_"+token)

	if (nrMaximoRequestPorSegundo != ""){
		nrMaximo, err := strconv.ParseInt(nrMaximoRequestPorSegundo, 10, 64)
		if err != nil {
			return getNrMaximoRequestPorSegundoTokenPadrao()
		}

		return nrMaximo
	}

	return getNrMaximoRequestPorSegundoTokenPadrao()
}


func getNrMaximoRequestPorSegundoTokenPadrao() int64 {
	nrMaximoRequestPorSegundo := os.Getenv("NR_MAXIMO_REQUEST_POR_SEGUNDO_TOKEN")

	nrMaximo, err := strconv.ParseInt(nrMaximoRequestPorSegundo, 10, 64)
	if err != nil {
		return int64(10)
	}

	return nrMaximo
}

func getDuracaoBloqueioToken(token string) time.Duration {
	duracaoBloqueio := os.Getenv("DURACAO_BLOQUEIO_TOKEN_"+token)

	if (duracaoBloqueio != ""){
		duration, err := time.ParseDuration(duracaoBloqueio)
		if err != nil {
			return getDuracaoBloqueioTokenPadrao()
		}

		return duration
	}

	return getDuracaoBloqueioTokenPadrao()
}

func getDuracaoBloqueioTokenPadrao() time.Duration {
	duracaoBloqueio := os.Getenv("DURACAO_BLOQUEIO_TOKEN_PADRAO")

	duration, err := time.ParseDuration(duracaoBloqueio)
	if err != nil {
		return time.Second * 10
	}

	return duration
}


func getNrMaximoRequestPorSegundoIP() int64 {
	nrMaximoRequestPorSegundo := os.Getenv("NR_MAXIMO_REQUEST_POR_SEGUNDO_IP")

	nrMaximo, err := strconv.ParseInt(nrMaximoRequestPorSegundo, 10, 64)
	if err != nil {
		return int64(10)
	}

	return nrMaximo
}

func getDuracaoBloqueioIP() time.Duration {
	duracaoBloqueio := os.Getenv("DURACAO_BLOQUEIO_IP")

	duration, err := time.ParseDuration(duracaoBloqueio)
	if err != nil {
		return time.Second * 10
	}

	return duration
}