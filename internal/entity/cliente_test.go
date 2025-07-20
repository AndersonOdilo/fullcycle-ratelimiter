package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculaNumeroRequestPorSegundo_requestsSemIntervalo(t *testing.T) {

	cliente := Cliente{
		Chave: "10.10.10.1",
	}

	//insere 20 request
	for i := 1; i <= 20; i++ {	
		cliente.InsereNovaRequest(time.Now().UnixMilli())
	}

	//espera 500 millisegundo
	time.Sleep(time.Millisecond * 500)


	//realiza o calculo de requests realizadas nos ultimo segundo
	nrRequestsPorSegundo := cliente.CalculaNumeroRequestPorSegundo()

	assert.Equal(t, int64(20), nrRequestsPorSegundo)
}

func TestCalculaNumeroRequestPorSegundo_requestsComIntervalo(t *testing.T) {

	cliente := Cliente{
		Chave: "10.10.10.1",
	}

	//insere uma request
	cliente.InsereNovaRequest(time.Now().UnixMilli())

	//espera 1 segundo
	time.Sleep(time.Second)

	//insere 20 request
	for i := 1; i <= 10; i++ {	
		cliente.InsereNovaRequest(time.Now().UnixMilli())
	}
	
	//espera 500 millisegundo
	time.Sleep(time.Millisecond * 500)

	//realiza o calculo de requests realizadas nos ultimo segundo
	nrRequestsPorSegundo := cliente.CalculaNumeroRequestPorSegundo()

	assert.Equal(t, int64(10), nrRequestsPorSegundo)
}

