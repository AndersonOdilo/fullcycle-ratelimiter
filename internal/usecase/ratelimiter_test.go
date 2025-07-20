package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/AndersonOdilo/fullcycle-ratelimiter/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



type ClienteRepositoryMock struct {
	mock.Mock
}

func (r *ClienteRepositoryMock) Obtem(ctx context.Context, chave string) (entity.Cliente, error){
	args := r.Called(ctx, chave)

    return args.Get(0).(entity.Cliente), args.Error(1)
}

func (r *ClienteRepositoryMock) Grava(ctx context.Context, cliente entity.Cliente) error {
	args := r.Called(ctx, cliente)

	return args.Error(0)
}

func TestRateLimiter_PrimeiraRequestIp(t *testing.T) {

	ctx := context.Background()
	unixTimestampRequest := time.Now().UnixMilli()
	ip := "10.10.10.1"
	token := ""

	clienteRepositoryMock := new(ClienteRepositoryMock)

	clienteRepositoryMock.On("Obtem", ctx, ip).Return(entity.Cliente{}, nil)

	cliente := entity.Cliente{
		Chave: ip,
		UnixRequest: []int64 {unixTimestampRequest},
	}

	clienteRepositoryMock.On("Grava", ctx, cliente).Return(nil)

	rateLimiterUseCase := NewRateLimiterUseCase(clienteRepositoryMock)

	acessoLiberado, err := rateLimiterUseCase.Execute(ctx, ip, token, unixTimestampRequest)
	
	assert.Nil(t, err)
	assert.Equal(t, AcessoLiberado(true), acessoLiberado)	
}

func TestRateLimiter_RequestComBloqueioIp(t *testing.T) {

	ctx := context.Background()
	unixTimestampRequest := time.Now().UnixMilli()
	ip := "10.10.10.1"
	token := ""

	clienteRepositoryMock := new(ClienteRepositoryMock)

	clienteRepositoryMock.On("Obtem", ctx, ip).Return(entity.Cliente{
		Chave: ip,
		UnixBloqueio: time.Now().Add(time.Second * 10).UnixMilli(),
	}, nil)

	rateLimiterUseCase := NewRateLimiterUseCase(clienteRepositoryMock)

	acessoLiberado, err := rateLimiterUseCase.Execute(ctx, ip, token, unixTimestampRequest)
	
	assert.Nil(t, err)
	assert.Equal(t, AcessoLiberado(false), acessoLiberado)
}

func TestRateLimiter_ExcedeuNumeroDeRequestIp(t *testing.T) {

	ctx := context.Background()
	unixTimestampRequest := time.Now().UnixMilli()
	ip := "10.10.10.1"
	token := ""

	clienteRepositoryMock := new(ClienteRepositoryMock)

	cliente := entity.Cliente{
		Chave: ip,
	}

	for i := 0; i < 20; i++ {
		cliente.InsereNovaRequest(unixTimestampRequest)
	}

	clienteRepositoryMock.On("Obtem", ctx, ip).Return(cliente, nil)

	cliente.UnixBloqueio = unixTimestampRequest + time.Duration(time.Second * 10).Milliseconds()

	clienteRepositoryMock.On("Grava", ctx, cliente).Return(nil)


	rateLimiterUseCase := NewRateLimiterUseCase(clienteRepositoryMock)

	acessoLiberado, err := rateLimiterUseCase.Execute(ctx, ip, token, unixTimestampRequest)
	
	assert.Nil(t, err)
	assert.Equal(t, AcessoLiberado(false), acessoLiberado)	
}

func TestRateLimiter_PrimeiraRequestToken(t *testing.T) {

	ctx := context.Background()
	unixTimestampRequest := time.Now().UnixMilli()
	ip := "10.10.10.1"
	token := "ABC1234"

	clienteRepositoryMock := new(ClienteRepositoryMock)

	clienteRepositoryMock.On("Obtem", ctx, token).Return(entity.Cliente{}, nil)

	cliente := entity.Cliente{
		Chave: token,
		UnixRequest: []int64 {unixTimestampRequest},
	}

	clienteRepositoryMock.On("Grava", ctx, cliente).Return(nil)

	rateLimiterUseCase := NewRateLimiterUseCase(clienteRepositoryMock)

	acessoLiberado, err := rateLimiterUseCase.Execute(ctx, ip, token, unixTimestampRequest)
	
	assert.Nil(t, err)
	assert.Equal(t, AcessoLiberado(true), acessoLiberado)	
}

func TestRateLimiter_RequestComBloqueioToken(t *testing.T) {

	ctx := context.Background()
	unixTimestampRequest := time.Now().UnixMilli()
	ip := "10.10.10.1"
	token := "ABC1234"

	clienteRepositoryMock := new(ClienteRepositoryMock)

	clienteRepositoryMock.On("Obtem", ctx, token).Return(entity.Cliente{
		Chave: token,
		UnixBloqueio: time.Now().Add(time.Second * 10).UnixMilli(),
	}, nil)

	rateLimiterUseCase := NewRateLimiterUseCase(clienteRepositoryMock)

	acessoLiberado, err := rateLimiterUseCase.Execute(ctx, ip, token, unixTimestampRequest)
	
	assert.Nil(t, err)
	assert.Equal(t, AcessoLiberado(false), acessoLiberado)
}

func TestRateLimiter_ExcedeuNumeroDeRequestToekn(t *testing.T) {

	ctx := context.Background()
	unixTimestampRequest := time.Now().UnixMilli()
	ip := "10.10.10.1"
	token := "ABC1234"

	clienteRepositoryMock := new(ClienteRepositoryMock)

	cliente := entity.Cliente{
		Chave: token,
	}

	for i := 0; i < 20; i++ {
		cliente.InsereNovaRequest(unixTimestampRequest)
	}

	clienteRepositoryMock.On("Obtem", ctx, token).Return(cliente, nil)

	cliente.UnixBloqueio = unixTimestampRequest + time.Duration(time.Second * 10).Milliseconds()

	clienteRepositoryMock.On("Grava", ctx, cliente).Return(nil)


	rateLimiterUseCase := NewRateLimiterUseCase(clienteRepositoryMock)

	acessoLiberado, err := rateLimiterUseCase.Execute(ctx, ip, token, unixTimestampRequest)
	
	assert.Nil(t, err)
	assert.Equal(t, AcessoLiberado(false), acessoLiberado)	
}



